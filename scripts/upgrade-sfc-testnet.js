#!/usr/bin/env node
/**
 * SFC Contract Upgrade Script for VinuChain Testnet
 *
 * This script:
 * 1. Deploys the new SFC implementation as a regular contract
 * 2. Calls upgradeCode() on NodeDriverAuth to copy the bytecode to the SFC proxy address
 *
 * Prerequisites:
 * - npm install ethers@5
 * - The admin account (NodeDriverAuth owner) private key
 * - VinuChain testnet RPC endpoint
 *
 * Usage:
 *   ADMIN_PRIVATE_KEY=0x... node scripts/upgrade-sfc-testnet.js
 *
 * Resume after partial failure (skip deploy, go straight to upgradeCode):
 *   ADMIN_PRIVATE_KEY=0x... RESUME_IMPL=0x<deployed-impl-address> node scripts/upgrade-sfc-testnet.js
 */

const { ethers } = require("ethers");
const fs = require("fs");
const path = require("path");

// --- Configuration ---
const SFC_ADDRESS = "0xfc00face00000000000000000000000000000000";
const NODE_DRIVER_AUTH_ADDRESS = "0xd100ae0000000000000000000000000000000000";
const VINUCHAIN_TESTNET_CHAIN_ID = 206; // 0xce
const MIN_SFC_CODE_BYTES = 30000; // SFC runtime bytecode should be at least this large

// NodeDriverAuth ABI — upgradeCode verifies both addresses have code before copying
const NODE_DRIVER_AUTH_ABI = [
  "function upgradeCode(address acc, address from) external",
  "function isOwner() view returns (bool)",
  "function owner() view returns (address)"
];

// Minimal SFC ABI for post-upgrade smoke test
const SFC_ABI = [
  "function currentEpoch() view returns (uint256)",
  "function totalStake() view returns (uint256)"
];

async function main() {
  const rpcUrl = process.env.RPC_URL || "https://testnet-rpc.vinuchain.org";
  const adminKey = process.env.ADMIN_PRIVATE_KEY;
  const resumeImpl = process.env.RESUME_IMPL;

  if (!adminKey) {
    console.error("Required: ADMIN_PRIVATE_KEY=0x<admin-private-key>");
    console.error("Optional: RPC_URL (defaults to https://testnet-rpc.vinuchain.org)");
    console.error("Optional: RESUME_IMPL=0x<impl-address> (skip deploy, resume from upgradeCode)");
    process.exit(1);
  }

  // Validate private key format
  if (!/^0x[0-9a-fA-F]{64}$/.test(adminKey)) {
    console.error("ERROR: ADMIN_PRIVATE_KEY must be a 0x-prefixed 64-character hex string.");
    process.exit(1);
  }

  const provider = new ethers.providers.JsonRpcProvider(rpcUrl);
  const network = await provider.getNetwork();

  if (network.chainId !== VINUCHAIN_TESTNET_CHAIN_ID) {
    console.error("ERROR: Connected to chain ID", network.chainId, "but expected VinuChain Testnet (206).");
    console.error("Aborting to prevent upgrading the wrong network.");
    process.exit(1);
  }

  const wallet = new ethers.Wallet(adminKey, provider);
  console.log("Network: VinuChain Testnet (chain ID " + network.chainId + ")");
  console.log("RPC:", rpcUrl);
  console.log("Admin address:", wallet.address);

  // Pre-flight balance check
  const balance = await provider.getBalance(wallet.address);
  console.log("Admin balance:", ethers.utils.formatEther(balance), "VC");
  if (balance.isZero()) {
    console.error("ERROR: Admin wallet has zero balance. Gas fees cannot be paid.");
    process.exit(1);
  }

  // Verify admin is the NodeDriverAuth owner
  const authDriver = new ethers.Contract(NODE_DRIVER_AUTH_ADDRESS, NODE_DRIVER_AUTH_ABI, wallet);
  const owner = await authDriver.owner();
  console.log("NodeDriverAuth owner:", owner);

  if (owner.toLowerCase() !== wallet.address.toLowerCase()) {
    console.error("ERROR: Admin wallet is not the NodeDriverAuth owner!");
    console.error("  Expected:", owner);
    console.error("  Got:     ", wallet.address);
    process.exit(1);
  }
  console.log("Admin ownership verified.\n");

  let newImplAddress;
  let deployedCode;

  if (resumeImpl) {
    // --- Resume mode: skip deploy ---
    console.log("--- RESUME MODE: Skipping deploy, using existing impl ---");
    newImplAddress = ethers.utils.getAddress(resumeImpl);
    deployedCode = await provider.getCode(newImplAddress);
    if (deployedCode === "0x") {
      console.error("ERROR: No code at resume address", newImplAddress);
      process.exit(1);
    }
    console.log("Impl address:", newImplAddress);
    console.log("Impl code size:", (deployedCode.length - 2) / 2, "bytes\n");
  } else {
    // --- Step 1: Deploy new SFC implementation ---
    const deployBytecode = "0x" + fs.readFileSync(
      path.join(__dirname, "..", "build", "sfc_deploy_bytecode.txt"),
      "utf8"
    ).trim();

    console.log("--- Step 1: Deploy new SFC implementation ---");
    console.log("Deploy bytecode size:", (deployBytecode.length - 2) / 2, "bytes");

    // Estimate gas
    const gasEstimate = await provider.estimateGas({ from: wallet.address, data: deployBytecode });
    const gasLimit = gasEstimate.mul(120).div(100); // 20% buffer
    console.log("Estimated gas:", gasEstimate.toString(), "(using", gasLimit.toString(), "with 20% buffer)");

    const nonce = await provider.getTransactionCount(wallet.address, "pending");
    console.log("Nonce:", nonce);

    const deployTx = await wallet.sendTransaction({
      data: deployBytecode,
      gasLimit: gasLimit,
      nonce: nonce,
    });
    console.log("Deploy tx hash:", deployTx.hash);
    console.log("Waiting for confirmation...");

    const receipt = await deployTx.wait();
    if (receipt.status !== 1) {
      console.error("ERROR: Deployment transaction failed!");
      console.error("Block:", receipt.blockNumber, "Gas used:", receipt.gasUsed.toString());
      process.exit(1);
    }

    newImplAddress = receipt.contractAddress;
    console.log("New SFC implementation deployed at:", newImplAddress);
    console.log("Block:", receipt.blockNumber, "Gas used:", receipt.gasUsed.toString());

    // Verify bytecode was deployed
    deployedCode = await provider.getCode(newImplAddress);
    const deployedBytes = (deployedCode.length - 2) / 2;
    if (deployedCode === "0x" || deployedBytes < MIN_SFC_CODE_BYTES) {
      console.error("ERROR: Deployed code too small (" + deployedBytes + " bytes). Expected at least", MIN_SFC_CODE_BYTES, "bytes.");
      process.exit(1);
    }
    console.log("Deployed code size:", deployedBytes, "bytes");
    console.log("\nIf the next step fails, resume with:");
    console.log("  RESUME_IMPL=" + newImplAddress + "\n");
  }

  // --- Step 2: upgradeCode to SFC proxy address ---
  console.log("--- Step 2: Upgrade code at SFC proxy address ---");
  console.log("Calling upgradeCode(" + SFC_ADDRESS + ", " + newImplAddress + ")");

  const upgradeGasEstimate = await authDriver.estimateGas.upgradeCode(SFC_ADDRESS, newImplAddress);
  const upgradeGasLimit = upgradeGasEstimate.mul(120).div(100);
  console.log("Estimated gas:", upgradeGasEstimate.toString(), "(using", upgradeGasLimit.toString(), "with 20% buffer)");

  const upgradeNonce = await provider.getTransactionCount(wallet.address, "pending");
  const upgradeTx = await authDriver.upgradeCode(SFC_ADDRESS, newImplAddress, {
    gasLimit: upgradeGasLimit,
    nonce: upgradeNonce,
  });
  console.log("upgradeCode tx hash:", upgradeTx.hash);
  console.log("Waiting for confirmation...");

  const upgradeReceipt = await upgradeTx.wait();
  if (upgradeReceipt.status !== 1) {
    console.error("ERROR: upgradeCode transaction failed!");
    console.error("Block:", upgradeReceipt.blockNumber, "Gas used:", upgradeReceipt.gasUsed.toString());
    process.exit(1);
  }
  console.log("Block:", upgradeReceipt.blockNumber, "Gas used:", upgradeReceipt.gasUsed.toString());

  // --- Step 3: Verify ---
  console.log("\n--- Step 3: Verification ---");
  const sfcCode = await provider.getCode(SFC_ADDRESS);
  if (sfcCode === deployedCode) {
    console.log("Bytecode verification: PASSED");
    console.log("SFC code size:", (sfcCode.length - 2) / 2, "bytes");
  } else {
    console.error("Bytecode verification: FAILED");
    console.error("SFC code size:", (sfcCode.length - 2) / 2, "bytes");
    console.error("Impl code size:", (deployedCode.length - 2) / 2, "bytes");
    process.exit(1);
  }

  // Post-upgrade smoke test
  console.log("\n--- Post-upgrade smoke test ---");
  const sfcContract = new ethers.Contract(SFC_ADDRESS, SFC_ABI, provider);
  let smokeFailed = false;
  try {
    const epoch = await sfcContract.currentEpoch();
    console.log("currentEpoch():", epoch.toString(), "- OK");
  } catch (err) {
    console.error("FAIL: currentEpoch() call failed:", err.message);
    smokeFailed = true;
  }
  try {
    const totalStake = await sfcContract.totalStake();
    console.log("totalStake():", ethers.utils.formatEther(totalStake), "VC - OK");
  } catch (err) {
    console.error("FAIL: totalStake() call failed:", err.message);
    smokeFailed = true;
  }

  console.log("\n--- Summary ---");
  console.log("New impl address:", newImplAddress);
  console.log("Deploy block:", resumeImpl ? "(resumed)" : "see above");
  console.log("upgradeCode block:", upgradeReceipt.blockNumber);
  console.log("upgradeCode tx:", upgradeTx.hash);

  if (smokeFailed) {
    console.error("\nSFC upgrade applied but SMOKE TESTS FAILED. The contract may be broken.");
    process.exit(2);
  }
  console.log("\nSFC upgrade completed successfully.");
}

main().catch((err) => {
  // Sanitize error output to avoid leaking private key
  const errMsg = err.message || String(err);
  console.error("Fatal error:", errMsg);
  process.exit(1);
});
