package task_1

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"MetaNode-Task/Store"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ActionContract() {
	//合约地址
	contractAddr := "0x5536116A208F710a01ACFd56c85a0BDdB6c189f9"

	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := Store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_JXD_key_JKK"))
	copy(value[:], []byte("demo_JXD_value_JKK2222"))
	//11155111 是sepolia 的测试网id
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" valueInContract:", valueInContract)
	fmt.Println(" value:", value)
	fmt.Println("is value saving in contract equals to origin :", valueInContract == value)
}
