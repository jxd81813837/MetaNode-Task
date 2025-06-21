package task_1

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func TransferETH() {
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress:", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())

	// 创建交易并签名交易
	// 估算 gas 价格
	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	head, err := client.HeaderByNumber(context.Background(), nil)

	gasFeeCap := new(big.Int).Add(head.BaseFee, new(big.Int).Mul(gasTipCap, big.NewInt(100)))

	to := common.HexToAddress("")
	var data []byte
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: new(big.Int).Mul(gasTipCap, big.NewInt(10)),
		Gas:       uint64(60000),
		To:        &to,
		Value:     big.NewInt(0),
		Data:      data,
	})

	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
