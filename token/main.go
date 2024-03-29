//go:generate cfxabigen --sol ./contract/hello.sol --pkg hello --out ./contract/hello.go
package main

import (
	"log"
	"math/big"
	"os"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

func main() {
	deploy()
	writeContract()
	accessContract()
	filterEvent()

	go watchEvent()
	go func() {
		for i := 0; i < 10; i++ {
			writeContract()
			time.Sleep(3 * time.Second)
		}
	}()
	select {}
}

func deploy() {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../keystore",
		Logger:       os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = client.AccountManager.UnlockDefault("hello")
	if err != nil {
		log.Fatal(err)
	}

	oneCfx := new(big.Int).Mul(big.NewInt(1e9), big.NewInt(1e9))
	tx, hash, _, err := DeployMyERC20Token(nil, client, new(big.Int).Mul(big.NewInt(1000000), oneCfx), "ABC", 18, "ABC")
	if err != nil {
		panic(err)
	}

	receipt, err := client.WaitForTransationReceipt(*hash, time.Second)
	if err != nil {
		panic(err)
	}

	logrus.WithFields(logrus.Fields{
		"tx":               tx,
		"hash":             hash,
		"contract address": receipt.ContractCreated,
	}).Info("deploy token done")
}

func writeContract() {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../keystore",
	})
	if err != nil {
		panic(err)
	}

	err = client.AccountManager.UnlockDefault("hello")
	if err != nil {
		panic(err)
	}

	contractAddr := cfxaddress.MustNew("cfxtest:acd7apn6pnfhna7w1pa8evzhwhv3085vjjp1b8bav5")

	instance, err := NewMyERC20Token(contractAddr, client)
	if err != nil {
		panic(err)
	}

	to := cfxaddress.MustNew("cfxtest:aasfup1wgjyxkzy3575cbnn87xj5tam2zud125twew")
	tx, hash, err := instance.Transfer(nil, to.MustGetCommonAddress(), big.NewInt(1))
	if err != nil {
		panic(err)
	}

	logrus.WithField("tx", tx).WithField("hash", hash).Info("transfer")
	receipt, err := client.WaitForTransationReceipt(*hash, time.Second)
	if err != nil {
		panic(err)
	}

	logrus.WithField("transfer receipt", receipt).Info()
}

func accessContract() {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../keystore",
	})
	if err != nil {
		panic(err)
	}

	err = client.AccountManager.UnlockDefault("hello")
	if err != nil {
		panic(err)
	}

	contractAddr := cfxaddress.MustNew("cfxtest:acd7apn6pnfhna7w1pa8evzhwhv3085vjjp1b8bav5")
	instance, err := NewMyERC20Token(contractAddr, client)
	if err != nil {
		panic(err)
	}

	user := cfxaddress.MustNew("cfxtest:aasfup1wgjyxkzy3575cbnn87xj5tam2zud125twew")
	result, err := instance.BalanceOf(nil, user.MustGetCommonAddress())
	if err != nil {
		panic(err)
	}

	logrus.WithField("balance", result).WithField("user", user).Info("access contract") // "bar"
}

func filterEvent() {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../keystore",
		Logger:       os.Stdout,
	})
	if err != nil {
		panic(err)
	}

	err = client.AccountManager.UnlockDefault("hello")
	if err != nil {
		panic(err)
	}

	contractAddr := cfxaddress.MustNew("cfxtest:acd7apn6pnfhna7w1pa8evzhwhv3085vjjp1b8bav5")
	instance, err := NewMyERC20Token(contractAddr, client)
	if err != nil {
		panic(err)
	}

	start := big.NewInt(35779622)
	end := big.NewInt(35779722)

	it, err := instance.FilterTransfer(&bind.FilterOpts{
		Start: types.NewEpochNumber(types.NewBigIntByRaw(start)),
		End:   types.NewEpochNumber(types.NewBigIntByRaw(end)),
	}, []common.Address{common.HexToAddress("0x1502ADd5a4a14c85C525e30a850c58fA15325f8C")}, nil,
	)

	if err != nil {
		panic(err)
	}

	for {
		if it.Next() {
			logrus.WithField("Transfer", it.Event).Info("Transfer log")
		} else {
			if err := it.Error(); err != nil {
				panic(err)
			}
			return
		}
	}
}

func watchEvent() {
	client, err := sdk.NewClient("ws://test.confluxrpc.com/ws", sdk.ClientOption{
		KeystorePath: "../keystore",
	})
	if err != nil {
		panic(err)
	}

	err = client.AccountManager.UnlockDefault("hello")
	if err != nil {
		panic(err)
	}

	contractAddr := cfxaddress.MustNew("cfxtest:acd7apn6pnfhna7w1pa8evzhwhv3085vjjp1b8bav5")
	instance, err := NewMyERC20Token(contractAddr, client)
	if err != nil {
		panic(err)
	}

	eventCh := make(chan *MyERC20TokenTransferOrChainReorg, 100)
	// reorgCh := make(chan types.ChainReorg, 100)
	sub, err := instance.WatchTransfer(nil, eventCh, nil, nil)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case l, ok := <-eventCh:
			if l.ChainReorg == nil {
				logrus.WithFields(logrus.Fields{
					"log": l.Event,
					"ok":  ok,
				}).Info("receive setted log")
			} else {
				logrus.WithFields(logrus.Fields{
					"reorg": l.ChainReorg,
					"ok":    ok,
				}).Info("receive setted log")
			}

		case err := <-sub.Err():
			panic(err)
		}
	}
}
