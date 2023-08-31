//// Code generated - DO NOT EDIT.
//// This file is a generated binding and any manual changes will be lost.
//
package driverauth100

//
//import (
//	"math/big"
//	"strings"
//
//	ethereum "github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/accounts/abi"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/event"
//)
//
//// Reference imports to suppress errors if they are not otherwise used.
//var (
//	_ = big.NewInt
//	_ = strings.NewReader
//	_ = ethereum.NotFound
//	_ = bind.Bind
//	_ = common.Big1
//	_ = types.BloomLookup
//	_ = event.NewSubscription
//)
//
//// ContractABI is the input ABI used to generate the binding from.
//const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sfc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_driver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDriverAuth\",\"type\":\"address\"}],\"name\":\"migrateTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"incBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"copyCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"diff\",\"type\":\"bytes\"}],\"name\":\"updateNetworkRules\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateNetworkVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"advanceEpochs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateValidatorWeight\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"updateValidatorPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
//
//// ContractBin is the compiled bytecode used for deploying new contracts.
//var ContractBin = "0x608060405234801561001057600080fd5b50611705806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80638da5cb5b116100ad578063d6a0c7af11610071578063d6a0c7af14610252578063e08d7e6614610265578063ebdf104c14610278578063f2fde38b1461028b578063fd1b6ec11461029e5761012c565b80638da5cb5b146101e65780638f32d59b14610204578063a4066fbe14610219578063b9cc6b1c1461022c578063c0c53b8b1461023f5761012c565b80634ddaf8f2116100f45780634ddaf8f2146101925780634feb92f3146101a557806366e7ea0f146101b8578063715018a6146101cb57806379bead38146101d35761012c565b80630aeeca001461013157806318f628d4146101465780631e702f8314610159578063242a6e3f1461016c578063267ab4461461017f575b600080fd5b61014461013f3660046110cb565b6102b1565b005b610144610154366004610ecd565b610342565b61014461016736600461113f565b6103e9565b61014461017a3660046110e9565b61047b565b61014461018d3660046110cb565b610510565b6101446101a0366004610d24565b610564565b6101446101b3366004610e01565b6105b8565b6101446101c6366004610dd1565b610622565b6101446106c0565b6101446101e1366004610dd1565b61072e565b6101ee610784565b6040516101fb91906113db565b60405180910390f35b61020c610793565b6040516101fb9190611581565b61014461022736600461113f565b6107a4565b61014461023a366004611095565b610800565b61014461024d366004610d84565b610856565b610144610260366004610d4a565b610915565b610144610273366004610f81565b61096b565b610144610286366004610fc3565b6109c7565b610144610299366004610d24565b610a6b565b6101446102ac366004610d4a565b610a9b565b6102b9610793565b6102de5760405162461bcd60e51b81526004016102d5906115e1565b60405180910390fd5b6067546040516205776560e91b81526001600160a01b0390911690630aeeca009061030d908490600401611621565b600060405180830381600087803b15801561032757600080fd5b505af115801561033b573d6000803e3d6000fd5b5050505050565b6067546001600160a01b0316331461036c5760405162461bcd60e51b81526004016102d590611611565b60665460405163063d8a3560e21b81526001600160a01b03909116906318f628d4906103ac908c908c908c908c908c908c908c908c908c906004016114a6565b600060405180830381600087803b1580156103c657600080fd5b505af11580156103da573d6000803e3d6000fd5b50505050505050505050505050565b6067546001600160a01b031633146104135760405162461bcd60e51b81526004016102d590611611565b606654604051631e702f8360e01b81526001600160a01b0390911690631e702f83906104459085908590600401611659565b600060405180830381600087803b15801561045f57600080fd5b505af1158015610473573d6000803e3d6000fd5b505050505050565b6066546001600160a01b031633146104a55760405162461bcd60e51b81526004016102d5906115d1565b60675460405163242a6e3f60e01b81526001600160a01b039091169063242a6e3f906104d99086908690869060040161162f565b600060405180830381600087803b1580156104f357600080fd5b505af1158015610507573d6000803e3d6000fd5b50505050505050565b610518610793565b6105345760405162461bcd60e51b81526004016102d5906115e1565b60675460405163133d5a2360e11b81526001600160a01b039091169063267ab4469061030d908490600401611621565b61056c610793565b6105885760405162461bcd60e51b81526004016102d5906115e1565b60675460405163da7fc24f60e01b81526001600160a01b039091169063da7fc24f9061030d9084906004016113db565b6067546001600160a01b031633146105e25760405162461bcd60e51b81526004016102d590611611565b606654604051634feb92f360e01b81526001600160a01b0390911690634feb92f3906103ac908c908c908c908c908c908c908c908c908c90600401611426565b6066546001600160a01b0316331461064c5760405162461bcd60e51b81526004016102d5906115d1565b6066546001600160a01b038381169116146106795760405162461bcd60e51b81526004016102d590611601565b6067546001600160a01b039081169063e30443bc9084906106a3908216318563ffffffff610af416565b6040518363ffffffff1660e01b815260040161044592919061140b565b6106c8610793565b6106e45760405162461bcd60e51b81526004016102d5906115e1565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b610736610793565b6107525760405162461bcd60e51b81526004016102d5906115e1565b606754604051630f37d5a760e31b81526001600160a01b03909116906379bead3890610445908590859060040161140b565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b6066546001600160a01b031633146107ce5760405162461bcd60e51b81526004016102d5906115d1565b60675460405163520337df60e11b81526001600160a01b039091169063a4066fbe906104459085908590600401611659565b610808610793565b6108245760405162461bcd60e51b81526004016102d5906115e1565b606754604051632e731ac760e21b81526001600160a01b039091169063b9cc6b1c90610445908590859060040161158f565b600054610100900460ff168061086f575061086f610b22565b8061087d575060005460ff16155b6108995760405162461bcd60e51b81526004016102d5906115f1565b600054610100900460ff161580156108c4576000805460ff1961ff0019909116610100171660011790555b6108cd82610b28565b606780546001600160a01b038086166001600160a01b0319928316179092556066805492871692909116919091179055801561090f576000805461ff00191690555b50505050565b61091d610793565b6109395760405162461bcd60e51b81526004016102d5906115e1565b60675460405163d6a0c7af60e01b81526001600160a01b039091169063d6a0c7af9061044590859085906004016113e9565b6067546001600160a01b031633146109955760405162461bcd60e51b81526004016102d590611611565b606654604051637046bf3360e11b81526001600160a01b039091169063e08d7e6690610445908590859060040161151e565b6067546001600160a01b031633146109f15760405162461bcd60e51b81526004016102d590611611565b606654604051633af7c41360e21b81526001600160a01b039091169063ebdf104c90610a2f908b908b908b908b908b908b908b908b90600401611530565b600060405180830381600087803b158015610a4957600080fd5b505af1158015610a5d573d6000803e3d6000fd5b505050505050505050505050565b610a73610793565b610a8f5760405162461bcd60e51b81526004016102d5906115e1565b610a9881610bfb565b50565b610aa3610793565b610abf5760405162461bcd60e51b81526004016102d5906115e1565b610ac882610c7d565b8015610ad85750610ad881610c7d565b6109395760405162461bcd60e51b81526004016102d5906115b1565b600082820183811015610b195760405162461bcd60e51b81526004016102d5906115c1565b90505b92915050565b303b1590565b600054610100900460ff1680610b415750610b41610b22565b80610b4f575060005460ff16155b610b6b5760405162461bcd60e51b81526004016102d5906115f1565b600054610100900460ff16158015610b96576000805460ff1961ff0019909116610100171660011790555b603380546001600160a01b0319166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015610bf7576000805461ff00191690555b5050565b6001600160a01b038116610c215760405162461bcd60e51b81526004016102d5906115a1565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b3b151590565b8035610b1c816116a5565b60008083601f840112610ca057600080fd5b50813567ffffffffffffffff811115610cb857600080fd5b602083019150836020820283011115610cd057600080fd5b9250929050565b60008083601f840112610ce957600080fd5b50813567ffffffffffffffff811115610d0157600080fd5b602083019150836001820283011115610cd057600080fd5b8035610b1c816116b9565b600060208284031215610d3657600080fd5b6000610d428484610c83565b949350505050565b60008060408385031215610d5d57600080fd5b6000610d698585610c83565b9250506020610d7a85828601610c83565b9150509250929050565b600080600060608486031215610d9957600080fd5b6000610da58686610c83565b9350506020610db686828701610c83565b9250506040610dc786828701610c83565b9150509250925092565b60008060408385031215610de457600080fd5b6000610df08585610c83565b9250506020610d7a85828601610d19565b60008060008060008060008060006101008a8c031215610e2057600080fd5b6000610e2c8c8c610c83565b9950506020610e3d8c828d01610d19565b98505060408a013567ffffffffffffffff811115610e5a57600080fd5b610e668c828d01610cd7565b97509750506060610e798c828d01610d19565b9550506080610e8a8c828d01610d19565b94505060a0610e9b8c828d01610d19565b93505060c0610eac8c828d01610d19565b92505060e0610ebd8c828d01610d19565b9150509295985092959850929598565b60008060008060008060008060006101208a8c031215610eec57600080fd5b6000610ef88c8c610c83565b9950506020610f098c828d01610d19565b9850506040610f1a8c828d01610d19565b9750506060610f2b8c828d01610d19565b9650506080610f3c8c828d01610d19565b95505060a0610f4d8c828d01610d19565b94505060c0610f5e8c828d01610d19565b93505060e0610f6f8c828d01610d19565b925050610100610ebd8c828d01610d19565b60008060208385031215610f9457600080fd5b823567ffffffffffffffff811115610fab57600080fd5b610fb785828601610c8e565b92509250509250929050565b6000806000806000806000806080898b031215610fdf57600080fd5b883567ffffffffffffffff811115610ff657600080fd5b6110028b828c01610c8e565b9850985050602089013567ffffffffffffffff81111561102157600080fd5b61102d8b828c01610c8e565b9650965050604089013567ffffffffffffffff81111561104c57600080fd5b6110588b828c01610c8e565b9450945050606089013567ffffffffffffffff81111561107757600080fd5b6110838b828c01610c8e565b92509250509295985092959890939650565b600080602083850312156110a857600080fd5b823567ffffffffffffffff8111156110bf57600080fd5b610fb785828601610cd7565b6000602082840312156110dd57600080fd5b6000610d428484610d19565b6000806000604084860312156110fe57600080fd5b600061110a8686610d19565b935050602084013567ffffffffffffffff81111561112757600080fd5b61113386828701610cd7565b92509250509250925092565b6000806040838503121561115257600080fd5b6000610df08585610d19565b61116781611670565b82525050565b60006111798385611667565b93506001600160fb1b0383111561118f57600080fd5b6020830292506111a083858461168f565b50500190565b6111678161167b565b60006111bb8385611667565b93506111c883858461168f565b6111d18361169b565b9093019392505050565b60006111e8602683611667565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000611230600e83611667565b6d1b9bdd08184818dbdb9d1c9858dd60921b815260200192915050565b600061125a601b83611667565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b6000611293601e83611667565b7f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000815260200192915050565b60006112cc602083611667565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000611305602e83611667565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626581526d195b881a5b9a5d1a585b1a5e995960921b602082015260400192915050565b6000611355602183611667565b7f726563697069656e74206973206e6f74207468652053464320636f6e747261638152601d60fa1b602082015260400192915050565b6000611398602583611667565b7f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e8152641d1c9858dd60da1b602082015260400192915050565b6111678161168c565b60208101610b1c828461115e565b604081016113f7828561115e565b611404602083018461115e565b9392505050565b60408101611419828561115e565b61140460208301846113d2565b6101008101611435828c61115e565b611442602083018b6113d2565b818103604083015261145581898b6111af565b905061146460608301886113d2565b61147160808301876113d2565b61147e60a08301866113d2565b61148b60c08301856113d2565b61149860e08301846113d2565b9a9950505050505050505050565b61012081016114b5828c61115e565b6114c2602083018b6113d2565b6114cf604083018a6113d2565b6114dc60608301896113d2565b6114e960808301886113d2565b6114f660a08301876113d2565b61150360c08301866113d2565b61151060e08301856113d2565b6114986101008301846113d2565b60208082528101610d4281848661116d565b60808082528101611542818a8c61116d565b9050818103602083015261155781888a61116d565b9050818103604083015261156c81868861116d565b9050818103606083015261149881848661116d565b60208101610b1c82846111a6565b60208082528101610d428184866111af565b60208082528101610b1c816111db565b60208082528101610b1c81611223565b60208082528101610b1c8161124d565b60208082528101610b1c81611286565b60208082528101610b1c816112bf565b60208082528101610b1c816112f8565b60208082528101610b1c81611348565b60208082528101610b1c8161138b565b60208101610b1c82846113d2565b6040810161163d82866113d2565b81810360208301526116508184866111af565b95945050505050565b6040810161141982856113d2565b90815260200190565b6000610b1c82611680565b151590565b6001600160a01b031690565b90565b82818337506000910152565b601f01601f191690565b6116ae81611670565b8114610a9857600080fd5b6116ae8161168c56fea365627a7a72315820f0f87bdd681162e5bf587e7b3f90de86c6b8812003df0c950a57724bff5f2a186c6578706572696d656e74616cf564736f6c63430005110040"
//
//// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
//func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
//	parsed, err := abi.JSON(strings.NewReader(ContractABI))
//	if err != nil {
//		return common.Address{}, nil, nil, err
//	}
//
//	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
//	if err != nil {
//		return common.Address{}, nil, nil, err
//	}
//	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
//}
//
//// Contract is an auto generated Go binding around an Ethereum contract.
//type Contract struct {
//	ContractCaller     // Read-only binding to the contract
//	ContractTransactor // Write-only binding to the contract
//	ContractFilterer   // Log filterer for contract events
//}
//
//// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
//type ContractCaller struct {
//	contract *bind.BoundContract // Generic contract wrapper for the low level calls
//}
//
//// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
//type ContractTransactor struct {
//	contract *bind.BoundContract // Generic contract wrapper for the low level calls
//}
//
//// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
//type ContractFilterer struct {
//	contract *bind.BoundContract // Generic contract wrapper for the low level calls
//}
//
//// ContractSession is an auto generated Go binding around an Ethereum contract,
//// with pre-set call and transact options.
//type ContractSession struct {
//	Contract     *Contract         // Generic contract binding to set the session for
//	CallOpts     bind.CallOpts     // Call options to use throughout this session
//	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
//}
//
//// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
//// with pre-set call options.
//type ContractCallerSession struct {
//	Contract *ContractCaller // Generic contract caller binding to set the session for
//	CallOpts bind.CallOpts   // Call options to use throughout this session
//}
//
//// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
//// with pre-set transact options.
//type ContractTransactorSession struct {
//	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
//	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
//}
//
//// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
//type ContractRaw struct {
//	Contract *Contract // Generic contract binding to access the raw methods on
//}
//
//// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
//type ContractCallerRaw struct {
//	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
//}
//
//// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
//type ContractTransactorRaw struct {
//	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
//}
//
//// NewContract creates a new instance of Contract, bound to a specific deployed contract.
//func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
//	contract, err := bindContract(address, backend, backend, backend)
//	if err != nil {
//		return nil, err
//	}
//	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
//}
//
//// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
//func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
//	contract, err := bindContract(address, caller, nil, nil)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractCaller{contract: contract}, nil
//}
//
//// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
//func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
//	contract, err := bindContract(address, nil, transactor, nil)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractTransactor{contract: contract}, nil
//}
//
//// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
//func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
//	contract, err := bindContract(address, nil, nil, filterer)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractFilterer{contract: contract}, nil
//}
//
//// bindContract binds a generic wrapper to an already deployed contract.
//func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
//	parsed, err := abi.JSON(strings.NewReader(ContractABI))
//	if err != nil {
//		return nil, err
//	}
//	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
//}
//
//// Call invokes the (constant) contract method with params as input values and
//// sets the output to result. The result type might be a single field for simple
//// returns, a slice of interfaces for anonymous returns and a struct for named
//// returns.
//func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
//	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
//}
//
//// Transfer initiates a plain transaction to move funds to the contract, calling
//// its default method if one is available.
//func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
//	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
//}
//
//// Transact invokes the (paid) contract method with params as input values.
//func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
//	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
//}
//
//// Call invokes the (constant) contract method with params as input values and
//// sets the output to result. The result type might be a single field for simple
//// returns, a slice of interfaces for anonymous returns and a struct for named
//// returns.
//func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
//	return _Contract.Contract.contract.Call(opts, result, method, params...)
//}
//
//// Transfer initiates a plain transaction to move funds to the contract, calling
//// its default method if one is available.
//func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
//	return _Contract.Contract.contract.Transfer(opts)
//}
//
//// Transact invokes the (paid) contract method with params as input values.
//func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
//	return _Contract.Contract.contract.Transact(opts, method, params...)
//}
//
//// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
////
//// Solidity: function isOwner() view returns(bool)
//func (_Contract *ContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "isOwner")
//
//	if err != nil {
//		return *new(bool), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
//
//	return out0, err
//
//}
//
//// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
////
//// Solidity: function isOwner() view returns(bool)
//func (_Contract *ContractSession) IsOwner() (bool, error) {
//	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
//}
//
//// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
////
//// Solidity: function isOwner() view returns(bool)
//func (_Contract *ContractCallerSession) IsOwner() (bool, error) {
//	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
//}
//
//// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
////
//// Solidity: function owner() view returns(address)
//func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "owner")
//
//	if err != nil {
//		return *new(common.Address), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
//
//	return out0, err
//
//}
//
//// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
////
//// Solidity: function owner() view returns(address)
//func (_Contract *ContractSession) Owner() (common.Address, error) {
//	return _Contract.Contract.Owner(&_Contract.CallOpts)
//}
//
//// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
////
//// Solidity: function owner() view returns(address)
//func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
//	return _Contract.Contract.Owner(&_Contract.CallOpts)
//}
//
//// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
////
//// Solidity: function advanceEpochs(uint256 num) returns()
//func (_Contract *ContractTransactor) AdvanceEpochs(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "advanceEpochs", num)
//}
//
//// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
////
//// Solidity: function advanceEpochs(uint256 num) returns()
//func (_Contract *ContractSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
//}
//
//// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
////
//// Solidity: function advanceEpochs(uint256 num) returns()
//func (_Contract *ContractTransactorSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
//}
//
//// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
////
//// Solidity: function copyCode(address acc, address from) returns()
//func (_Contract *ContractTransactor) CopyCode(opts *bind.TransactOpts, acc common.Address, from common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "copyCode", acc, from)
//}
//
//// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
////
//// Solidity: function copyCode(address acc, address from) returns()
//func (_Contract *ContractSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
//}
//
//// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
////
//// Solidity: function copyCode(address acc, address from) returns()
//func (_Contract *ContractTransactorSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
//}
//
//// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
////
//// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
//func (_Contract *ContractTransactor) DeactivateValidator(opts *bind.TransactOpts, validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "deactivateValidator", validatorID, status)
//}
//
//// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
////
//// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
//func (_Contract *ContractSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
//}
//
//// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
////
//// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
//func (_Contract *ContractTransactorSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
//}
//
//// IncBalance is a paid mutator transaction binding the contract method 0x66e7ea0f.
////
//// Solidity: function incBalance(address acc, uint256 diff) returns()
//func (_Contract *ContractTransactor) IncBalance(opts *bind.TransactOpts, acc common.Address, diff *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "incBalance", acc, diff)
//}
//
//// IncBalance is a paid mutator transaction binding the contract method 0x66e7ea0f.
////
//// Solidity: function incBalance(address acc, uint256 diff) returns()
//func (_Contract *ContractSession) IncBalance(acc common.Address, diff *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.IncBalance(&_Contract.TransactOpts, acc, diff)
//}
//
//// IncBalance is a paid mutator transaction binding the contract method 0x66e7ea0f.
////
//// Solidity: function incBalance(address acc, uint256 diff) returns()
//func (_Contract *ContractTransactorSession) IncBalance(acc common.Address, diff *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.IncBalance(&_Contract.TransactOpts, acc, diff)
//}
//
//// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
////
//// Solidity: function initialize(address _sfc, address _driver, address _owner) returns()
//func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _sfc common.Address, _driver common.Address, _owner common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "initialize", _sfc, _driver, _owner)
//}
//
//// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
////
//// Solidity: function initialize(address _sfc, address _driver, address _owner) returns()
//func (_Contract *ContractSession) Initialize(_sfc common.Address, _driver common.Address, _owner common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _sfc, _driver, _owner)
//}
//
//// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
////
//// Solidity: function initialize(address _sfc, address _driver, address _owner) returns()
//func (_Contract *ContractTransactorSession) Initialize(_sfc common.Address, _driver common.Address, _owner common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _sfc, _driver, _owner)
//}
//
//// MigrateTo is a paid mutator transaction binding the contract method 0x4ddaf8f2.
////
//// Solidity: function migrateTo(address newDriverAuth) returns()
//func (_Contract *ContractTransactor) MigrateTo(opts *bind.TransactOpts, newDriverAuth common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "migrateTo", newDriverAuth)
//}
//
//// MigrateTo is a paid mutator transaction binding the contract method 0x4ddaf8f2.
////
//// Solidity: function migrateTo(address newDriverAuth) returns()
//func (_Contract *ContractSession) MigrateTo(newDriverAuth common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.MigrateTo(&_Contract.TransactOpts, newDriverAuth)
//}
//
//// MigrateTo is a paid mutator transaction binding the contract method 0x4ddaf8f2.
////
//// Solidity: function migrateTo(address newDriverAuth) returns()
//func (_Contract *ContractTransactorSession) MigrateTo(newDriverAuth common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.MigrateTo(&_Contract.TransactOpts, newDriverAuth)
//}
//
//// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
////
//// Solidity: function renounceOwnership() returns()
//func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "renounceOwnership")
//}
//
//// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
////
//// Solidity: function renounceOwnership() returns()
//func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
//	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
//}
//
//// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
////
//// Solidity: function renounceOwnership() returns()
//func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
//	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
//}
//
//// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
////
//// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
//func (_Contract *ContractTransactor) SealEpoch(opts *bind.TransactOpts, offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "sealEpoch", offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
//}
//
//// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
////
//// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
//func (_Contract *ContractSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
//}
//
//// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
////
//// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
//func (_Contract *ContractTransactorSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
//}
//
//// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
////
//// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
//func (_Contract *ContractTransactor) SealEpochValidators(opts *bind.TransactOpts, nextValidatorIDs []*big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "sealEpochValidators", nextValidatorIDs)
//}
//
//// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
////
//// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
//func (_Contract *ContractSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
//}
//
//// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
////
//// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
//func (_Contract *ContractTransactorSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
//}
//
//// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
////
//// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
//func (_Contract *ContractTransactor) SetGenesisDelegation(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "setGenesisDelegation", delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
//}
//
//// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
////
//// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
//func (_Contract *ContractSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
//}
//
//// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
////
//// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
//func (_Contract *ContractTransactorSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
//}
//
//// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
////
//// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
//func (_Contract *ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, _auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "setGenesisValidator", _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
//}
//
//// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
////
//// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
//func (_Contract *ContractSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
//}
//
//// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
////
//// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
//func (_Contract *ContractTransactorSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
//}
//
//// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
////
//// Solidity: function transferOwnership(address newOwner) returns()
//func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
//}
//
//// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
////
//// Solidity: function transferOwnership(address newOwner) returns()
//func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
//}
//
//// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
////
//// Solidity: function transferOwnership(address newOwner) returns()
//func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
//}
//
//// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
////
//// Solidity: function updateNetworkRules(bytes diff) returns()
//func (_Contract *ContractTransactor) UpdateNetworkRules(opts *bind.TransactOpts, diff []byte) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateNetworkRules", diff)
//}
//
//// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
////
//// Solidity: function updateNetworkRules(bytes diff) returns()
//func (_Contract *ContractSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
//}
//
//// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
////
//// Solidity: function updateNetworkRules(bytes diff) returns()
//func (_Contract *ContractTransactorSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
//}
//
//// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
////
//// Solidity: function updateNetworkVersion(uint256 version) returns()
//func (_Contract *ContractTransactor) UpdateNetworkVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateNetworkVersion", version)
//}
//
//// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
////
//// Solidity: function updateNetworkVersion(uint256 version) returns()
//func (_Contract *ContractSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
//}
//
//// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
////
//// Solidity: function updateNetworkVersion(uint256 version) returns()
//func (_Contract *ContractTransactorSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
//}
//
//// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
////
//// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
//func (_Contract *ContractTransactor) UpdateValidatorPubkey(opts *bind.TransactOpts, validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateValidatorPubkey", validatorID, pubkey)
//}
//
//// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
////
//// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
//func (_Contract *ContractSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
//}
//
//// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
////
//// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
//func (_Contract *ContractTransactorSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
//}
//
//// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
////
//// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
//func (_Contract *ContractTransactor) UpdateValidatorWeight(opts *bind.TransactOpts, validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateValidatorWeight", validatorID, value)
//}
//
//// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
////
//// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
//func (_Contract *ContractSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
//}
//
//// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
////
//// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
//func (_Contract *ContractTransactorSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
//}
//
//// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
//type ContractOwnershipTransferredIterator struct {
//	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractOwnershipTransferredIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractOwnershipTransferred)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractOwnershipTransferred)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractOwnershipTransferredIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractOwnershipTransferredIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
//type ContractOwnershipTransferred struct {
//	PreviousOwner common.Address
//	NewOwner      common.Address
//	Raw           types.Log // Blockchain specific contextual infos
//}
//
//// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
////
//// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
//func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {
//
//	var previousOwnerRule []interface{}
//	for _, previousOwnerItem := range previousOwner {
//		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
//	}
//	var newOwnerRule []interface{}
//	for _, newOwnerItem := range newOwner {
//		newOwnerRule = append(newOwnerRule, newOwnerItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
//}
//
//// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
////
//// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
//func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
//
//	var previousOwnerRule []interface{}
//	for _, previousOwnerItem := range previousOwner {
//		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
//	}
//	var newOwnerRule []interface{}
//	for _, newOwnerItem := range newOwner {
//		newOwnerRule = append(newOwnerRule, newOwnerItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractOwnershipTransferred)
//				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
////
//// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
//func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
//	event := new(ContractOwnershipTransferred)
//	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//var ContractBinRuntime = "0x608060405234801561001057600080fd5b506004361061012c5760003560e01c80638da5cb5b116100ad578063d6a0c7af11610071578063d6a0c7af14610252578063e08d7e6614610265578063ebdf104c14610278578063f2fde38b1461028b578063fd1b6ec11461029e5761012c565b80638da5cb5b146101e65780638f32d59b14610204578063a4066fbe14610219578063b9cc6b1c1461022c578063c0c53b8b1461023f5761012c565b80634ddaf8f2116100f45780634ddaf8f2146101925780634feb92f3146101a557806366e7ea0f146101b8578063715018a6146101cb57806379bead38146101d35761012c565b80630aeeca001461013157806318f628d4146101465780631e702f8314610159578063242a6e3f1461016c578063267ab4461461017f575b600080fd5b61014461013f3660046110cb565b6102b1565b005b610144610154366004610ecd565b610342565b61014461016736600461113f565b6103e9565b61014461017a3660046110e9565b61047b565b61014461018d3660046110cb565b610510565b6101446101a0366004610d24565b610564565b6101446101b3366004610e01565b6105b8565b6101446101c6366004610dd1565b610622565b6101446106c0565b6101446101e1366004610dd1565b61072e565b6101ee610784565b6040516101fb91906113db565b60405180910390f35b61020c610793565b6040516101fb9190611581565b61014461022736600461113f565b6107a4565b61014461023a366004611095565b610800565b61014461024d366004610d84565b610856565b610144610260366004610d4a565b610915565b610144610273366004610f81565b61096b565b610144610286366004610fc3565b6109c7565b610144610299366004610d24565b610a6b565b6101446102ac366004610d4a565b610a9b565b6102b9610793565b6102de5760405162461bcd60e51b81526004016102d5906115e1565b60405180910390fd5b6067546040516205776560e91b81526001600160a01b0390911690630aeeca009061030d908490600401611621565b600060405180830381600087803b15801561032757600080fd5b505af115801561033b573d6000803e3d6000fd5b5050505050565b6067546001600160a01b0316331461036c5760405162461bcd60e51b81526004016102d590611611565b60665460405163063d8a3560e21b81526001600160a01b03909116906318f628d4906103ac908c908c908c908c908c908c908c908c908c906004016114a6565b600060405180830381600087803b1580156103c657600080fd5b505af11580156103da573d6000803e3d6000fd5b50505050505050505050505050565b6067546001600160a01b031633146104135760405162461bcd60e51b81526004016102d590611611565b606654604051631e702f8360e01b81526001600160a01b0390911690631e702f83906104459085908590600401611659565b600060405180830381600087803b15801561045f57600080fd5b505af1158015610473573d6000803e3d6000fd5b505050505050565b6066546001600160a01b031633146104a55760405162461bcd60e51b81526004016102d5906115d1565b60675460405163242a6e3f60e01b81526001600160a01b039091169063242a6e3f906104d99086908690869060040161162f565b600060405180830381600087803b1580156104f357600080fd5b505af1158015610507573d6000803e3d6000fd5b50505050505050565b610518610793565b6105345760405162461bcd60e51b81526004016102d5906115e1565b60675460405163133d5a2360e11b81526001600160a01b039091169063267ab4469061030d908490600401611621565b61056c610793565b6105885760405162461bcd60e51b81526004016102d5906115e1565b60675460405163da7fc24f60e01b81526001600160a01b039091169063da7fc24f9061030d9084906004016113db565b6067546001600160a01b031633146105e25760405162461bcd60e51b81526004016102d590611611565b606654604051634feb92f360e01b81526001600160a01b0390911690634feb92f3906103ac908c908c908c908c908c908c908c908c908c90600401611426565b6066546001600160a01b0316331461064c5760405162461bcd60e51b81526004016102d5906115d1565b6066546001600160a01b038381169116146106795760405162461bcd60e51b81526004016102d590611601565b6067546001600160a01b039081169063e30443bc9084906106a3908216318563ffffffff610af416565b6040518363ffffffff1660e01b815260040161044592919061140b565b6106c8610793565b6106e45760405162461bcd60e51b81526004016102d5906115e1565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b610736610793565b6107525760405162461bcd60e51b81526004016102d5906115e1565b606754604051630f37d5a760e31b81526001600160a01b03909116906379bead3890610445908590859060040161140b565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b6066546001600160a01b031633146107ce5760405162461bcd60e51b81526004016102d5906115d1565b60675460405163520337df60e11b81526001600160a01b039091169063a4066fbe906104459085908590600401611659565b610808610793565b6108245760405162461bcd60e51b81526004016102d5906115e1565b606754604051632e731ac760e21b81526001600160a01b039091169063b9cc6b1c90610445908590859060040161158f565b600054610100900460ff168061086f575061086f610b22565b8061087d575060005460ff16155b6108995760405162461bcd60e51b81526004016102d5906115f1565b600054610100900460ff161580156108c4576000805460ff1961ff0019909116610100171660011790555b6108cd82610b28565b606780546001600160a01b038086166001600160a01b0319928316179092556066805492871692909116919091179055801561090f576000805461ff00191690555b50505050565b61091d610793565b6109395760405162461bcd60e51b81526004016102d5906115e1565b60675460405163d6a0c7af60e01b81526001600160a01b039091169063d6a0c7af9061044590859085906004016113e9565b6067546001600160a01b031633146109955760405162461bcd60e51b81526004016102d590611611565b606654604051637046bf3360e11b81526001600160a01b039091169063e08d7e6690610445908590859060040161151e565b6067546001600160a01b031633146109f15760405162461bcd60e51b81526004016102d590611611565b606654604051633af7c41360e21b81526001600160a01b039091169063ebdf104c90610a2f908b908b908b908b908b908b908b908b90600401611530565b600060405180830381600087803b158015610a4957600080fd5b505af1158015610a5d573d6000803e3d6000fd5b505050505050505050505050565b610a73610793565b610a8f5760405162461bcd60e51b81526004016102d5906115e1565b610a9881610bfb565b50565b610aa3610793565b610abf5760405162461bcd60e51b81526004016102d5906115e1565b610ac882610c7d565b8015610ad85750610ad881610c7d565b6109395760405162461bcd60e51b81526004016102d5906115b1565b600082820183811015610b195760405162461bcd60e51b81526004016102d5906115c1565b90505b92915050565b303b1590565b600054610100900460ff1680610b415750610b41610b22565b80610b4f575060005460ff16155b610b6b5760405162461bcd60e51b81526004016102d5906115f1565b600054610100900460ff16158015610b96576000805460ff1961ff0019909116610100171660011790555b603380546001600160a01b0319166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015610bf7576000805461ff00191690555b5050565b6001600160a01b038116610c215760405162461bcd60e51b81526004016102d5906115a1565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b3b151590565b8035610b1c816116a5565b60008083601f840112610ca057600080fd5b50813567ffffffffffffffff811115610cb857600080fd5b602083019150836020820283011115610cd057600080fd5b9250929050565b60008083601f840112610ce957600080fd5b50813567ffffffffffffffff811115610d0157600080fd5b602083019150836001820283011115610cd057600080fd5b8035610b1c816116b9565b600060208284031215610d3657600080fd5b6000610d428484610c83565b949350505050565b60008060408385031215610d5d57600080fd5b6000610d698585610c83565b9250506020610d7a85828601610c83565b9150509250929050565b600080600060608486031215610d9957600080fd5b6000610da58686610c83565b9350506020610db686828701610c83565b9250506040610dc786828701610c83565b9150509250925092565b60008060408385031215610de457600080fd5b6000610df08585610c83565b9250506020610d7a85828601610d19565b60008060008060008060008060006101008a8c031215610e2057600080fd5b6000610e2c8c8c610c83565b9950506020610e3d8c828d01610d19565b98505060408a013567ffffffffffffffff811115610e5a57600080fd5b610e668c828d01610cd7565b97509750506060610e798c828d01610d19565b9550506080610e8a8c828d01610d19565b94505060a0610e9b8c828d01610d19565b93505060c0610eac8c828d01610d19565b92505060e0610ebd8c828d01610d19565b9150509295985092959850929598565b60008060008060008060008060006101208a8c031215610eec57600080fd5b6000610ef88c8c610c83565b9950506020610f098c828d01610d19565b9850506040610f1a8c828d01610d19565b9750506060610f2b8c828d01610d19565b9650506080610f3c8c828d01610d19565b95505060a0610f4d8c828d01610d19565b94505060c0610f5e8c828d01610d19565b93505060e0610f6f8c828d01610d19565b925050610100610ebd8c828d01610d19565b60008060208385031215610f9457600080fd5b823567ffffffffffffffff811115610fab57600080fd5b610fb785828601610c8e565b92509250509250929050565b6000806000806000806000806080898b031215610fdf57600080fd5b883567ffffffffffffffff811115610ff657600080fd5b6110028b828c01610c8e565b9850985050602089013567ffffffffffffffff81111561102157600080fd5b61102d8b828c01610c8e565b9650965050604089013567ffffffffffffffff81111561104c57600080fd5b6110588b828c01610c8e565b9450945050606089013567ffffffffffffffff81111561107757600080fd5b6110838b828c01610c8e565b92509250509295985092959890939650565b600080602083850312156110a857600080fd5b823567ffffffffffffffff8111156110bf57600080fd5b610fb785828601610cd7565b6000602082840312156110dd57600080fd5b6000610d428484610d19565b6000806000604084860312156110fe57600080fd5b600061110a8686610d19565b935050602084013567ffffffffffffffff81111561112757600080fd5b61113386828701610cd7565b92509250509250925092565b6000806040838503121561115257600080fd5b6000610df08585610d19565b61116781611670565b82525050565b60006111798385611667565b93506001600160fb1b0383111561118f57600080fd5b6020830292506111a083858461168f565b50500190565b6111678161167b565b60006111bb8385611667565b93506111c883858461168f565b6111d18361169b565b9093019392505050565b60006111e8602683611667565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000611230600e83611667565b6d1b9bdd08184818dbdb9d1c9858dd60921b815260200192915050565b600061125a601b83611667565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b6000611293601e83611667565b7f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000815260200192915050565b60006112cc602083611667565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000611305602e83611667565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626581526d195b881a5b9a5d1a585b1a5e995960921b602082015260400192915050565b6000611355602183611667565b7f726563697069656e74206973206e6f74207468652053464320636f6e747261638152601d60fa1b602082015260400192915050565b6000611398602583611667565b7f63616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e8152641d1c9858dd60da1b602082015260400192915050565b6111678161168c565b60208101610b1c828461115e565b604081016113f7828561115e565b611404602083018461115e565b9392505050565b60408101611419828561115e565b61140460208301846113d2565b6101008101611435828c61115e565b611442602083018b6113d2565b818103604083015261145581898b6111af565b905061146460608301886113d2565b61147160808301876113d2565b61147e60a08301866113d2565b61148b60c08301856113d2565b61149860e08301846113d2565b9a9950505050505050505050565b61012081016114b5828c61115e565b6114c2602083018b6113d2565b6114cf604083018a6113d2565b6114dc60608301896113d2565b6114e960808301886113d2565b6114f660a08301876113d2565b61150360c08301866113d2565b61151060e08301856113d2565b6114986101008301846113d2565b60208082528101610d4281848661116d565b60808082528101611542818a8c61116d565b9050818103602083015261155781888a61116d565b9050818103604083015261156c81868861116d565b9050818103606083015261149881848661116d565b60208101610b1c82846111a6565b60208082528101610d428184866111af565b60208082528101610b1c816111db565b60208082528101610b1c81611223565b60208082528101610b1c8161124d565b60208082528101610b1c81611286565b60208082528101610b1c816112bf565b60208082528101610b1c816112f8565b60208082528101610b1c81611348565b60208082528101610b1c8161138b565b60208101610b1c82846113d2565b6040810161163d82866113d2565b81810360208301526116508184866111af565b95945050505050565b6040810161141982856113d2565b90815260200190565b6000610b1c82611680565b151590565b6001600160a01b031690565b90565b82818337506000910152565b601f01601f191690565b6116ae81611670565b8114610a9857600080fd5b6116ae8161168c56fea365627a7a72315820f0f87bdd681162e5bf587e7b3f90de86c6b8812003df0c950a57724bff5f2a186c6578706572696d656e74616cf564736f6c63430005110040"
