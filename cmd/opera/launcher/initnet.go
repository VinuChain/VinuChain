package launcher

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/Fantom-foundation/go-opera/gossip/emitter"
	"github.com/Fantom-foundation/go-opera/integration"
	"github.com/Fantom-foundation/go-opera/integration/makefakegenesis"
	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis/gpos"
	futils "github.com/Fantom-foundation/go-opera/utils"
	"github.com/Fantom-foundation/go-opera/valkeystore"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/crypto"
	cli "gopkg.in/urfave/cli.v1"
	// "github.com/Fantom-foundation/go-opera/integration/makefakegenesis"
)

// FakeNetFlag enables special testnet, where validators are automatically created
var InitNetCommand = cli.Command{
	Name:     "network",
	Usage:    "network command [command options] [arguments...]",
	Category: "network COMMANDS",

	Subcommands: []cli.Command{
		{
			Name:      "new",
			Usage:     "Create a new network",
			Action:    utils.MigrateFlags(newVinuNetwork),
			ArgsUsage: "network new <val_num> [flags]",
			Flags: []cli.Flag{
				DataDirFlag,
				utils.KeyStoreDirFlag,
				utils.PasswordFileFlag,
				ValidatorsFileFlag,
			},
			Description: `
opera network new
`,
		},
	},
}

func getValidatorsNum(ctx *cli.Context) (num idx.Validator, err error) {
	if ctx.NArg() != 1 {
		err = fmt.Errorf("use <val_num> format")
		return
	}
	s := ctx.Args().First()
	var u32 uint64
	u32, err = strconv.ParseUint(s, 10, 32)
	if err != nil {
		return
	}
	num = idx.Validator(u32)
	if num < 0 {
		err = fmt.Errorf("key-num should be in range from 1 to validators (<key-num>/<validators>), or should be zero for non-validator node")
		return
	}
	return
}

// validatorKeyCreate creates a new validator key into the keystore defined by the CLI flags.
func ValidatorCreate(ctx *cli.Context, valId int) (*gpos.Validator, error) {
	cfg := makeAllConfigs(ctx)
	utils.SetNodeConfig(ctx, &cfg.Node)

	password := getPassPhrase("Your new validator key is locked with a password. Please give a password. Do not forget this password.", true, 0, utils.MakePasswordList(ctx))

	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		utils.Fatalf("Failed to create account: %v", err)
	}
	privateKey := crypto.FromECDSA(privateKeyECDSA)
	publicKey := validatorpk.PubKey{
		Raw:  crypto.FromECDSAPub(&privateKeyECDSA.PublicKey),
		Type: validatorpk.Types.Secp256k1,
	}

	addr := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)

	fmt.Println("Account address: ", addr.String())

	//
	stack := makeConfigNode(ctx, &cfg.Node)
	coinbase := integration.SetAccountKey(stack.AccountManager(), privateKeyECDSA, password)
	fmt.Println("Unlocked fake validator account", "address", coinbase.Address.Hex())

	//

	valKeystore := valkeystore.NewDefaultFileRawKeystore(path.Join(getValKeystoreDir(cfg.Node), "validator"))
	err = valKeystore.Add(publicKey, privateKey, password)
	if err != nil {
		utils.Fatalf("Failed to create account: %v", err)
	}

	// Sanity check
	_, err = valKeystore.Get(publicKey, password)
	if err != nil {
		utils.Fatalf("Failed to decrypt the account: %v", err)
	}

	fmt.Printf("\nYour new key was generated\n\n")
	fmt.Printf("Validator ID:                %d\n", valId)
	fmt.Printf("Public key:                  %s\n", publicKey.String())
	fmt.Printf("Path of the secret key file: %s\n\n", valKeystore.PathOf(publicKey))

	return &gpos.Validator{
		ID:               idx.ValidatorID(valId),
		Address:          addr,
		PubKey:           publicKey,
		CreationTime:     makefakegenesis.FakeGenesisTime,
		CreationEpoch:    0,
		DeactivatedTime:  0,
		DeactivatedEpoch: 0,
		Status:           0,
	}, nil

}

func newVinuNetwork(ctx *cli.Context) error {

	num, err := getValidatorsNum(ctx)
	if err != nil {
		return err
	}
	validators := make([]gpos.Validator, 0, num)

	origDatadir := ctx.GlobalString(DataDirFlag.Name)
	for i := 1; i <= int(num); i++ {
		ctx.GlobalSet(DataDirFlag.Name, fmt.Sprintf("%s%d", origDatadir, i))
		fmt.Println("tmpCtx.DataDir: ", ctx.GlobalString(DataDirFlag.Name))
		val, err := ValidatorCreate(ctx, i)
		if err != nil {
			return err
		}
		validators = append(validators, *val)
		ctx.GlobalSet(DataDirFlag.Name, origDatadir)
	}

	epoch := idx.Epoch(2)
	block := idx.Block(1)
	// Create genesisStore
	genesisStore := makefakegenesis.VinuTestGenesisStoreWithRulesAndStart(futils.ToFtm(1000000000), futils.ToFtm(5000000), opera.VitainuTestNetRules(), epoch, block, validators)

	// Save validators to file for future use
	if ctx.GlobalIsSet(ValidatorsFileFlag.Name) {
		err := saveValidators(ctx, validators)
		if err != nil {
			return err
		}
	}

	for _, val := range validators {

		ctx.GlobalSet(DataDirFlag.Name, fmt.Sprintf("%s%d", origDatadir, val.ID))
		tmpCfg := makeAllConfigs(ctx)
		//tmpCfg.Node.DataDir = fmt.Sprintf("%s%d", cfg.Node.DataDir, val.ID)

		tmpCfg.Emitter.Validator = emitter.ValidatorConfig{
			ID:     val.ID,
			PubKey: val.PubKey,
		}
		emitCfg := emitter.DefaultConfig()
		emitCfg.EmitIntervals.Max = 10 * time.Second // don't wait long in vinunet
		emitCfg.EmitIntervals.DoublesignProtection = emitCfg.EmitIntervals.Max / 2

		fmt.Println("make node with config: ", tmpCfg)
		time.Sleep(5 * time.Second)
		node, _, nodeCloser := makeNode(ctx, tmpCfg, genesisStore)

		defer nodeCloser()
		fmt.Printf("Node %s created (validator %d)\n", node.Config().P2P.ListenAddr, val.ID)
		node.Close()
		node.Wait()
		time.Sleep(5 * time.Second)
	}

	return nil
}

func saveValidators(ctx *cli.Context, validators []gpos.Validator) error {
	// Save validators to file for future use
	if ctx.GlobalIsSet(ValidatorsFileFlag.Name) {
		f, err := os.Create(ctx.GlobalString(ValidatorsFileFlag.Name))
		if err != nil {
			return err
		}
		defer f.Close()
		for i, val := range validators {
			line := val.PubKey.String() + "\n"
			if i == len(validators)-1 {
				line = val.PubKey.String()
			}

			_, err := f.WriteString(line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
