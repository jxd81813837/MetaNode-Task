package task_1

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func QueryBlock(blockId int64) {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/alcht_Xm60QgvFAezQmXFFbyyLYM3nryPeL6")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(blockId)

	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	fmt.Println(header.Number.Uint64())     // 5671744
	fmt.Println(header.Time)                // 1712798400
	fmt.Println(header.Difficulty.Uint64()) // 0
	fmt.Println(header.Hash().Hex())        // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	if err != nil {
		log.Fatal(err)
	}
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())                   // 区块id 5671744
	fmt.Println(block.Time())                              // 时间戳1712798400
	fmt.Println("Difficulty", block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())                        //区块哈希  0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println(len(block.Transactions()))                 //交易数量 70
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 70
}
