// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {
// 	// Ethereumノードへの接続
// 	client, err := ethclient.Dial("https://mainnet.infura.io")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// アカウントアドレスの設定と現在の残高取得
// 	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
// 	balance, err := client.BalanceAt(context.Background(), account, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(balance) // 25893180161173005034

// 	// 特定のブロック時点での残高取得
// 	blockNumber := big.NewInt(5532993)
// 	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(balanceAt) // 25729324269165216042

// 	// WeiからETHへの変換
// 	fbalance := new(big.Float)
// 	fbalance.SetString(balanceAt.String())
// 	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
// 	fmt.Println(ethValue) // 25.729324269165216041

// 	// ペンディング状態の残高取得
// 	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
// 	fmt.Println(pendingBalance) // 25729324269165216042
// }
