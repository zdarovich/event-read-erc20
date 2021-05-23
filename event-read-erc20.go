package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	token "./pancakeswap/factory"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

type PairCreated struct {
	token0 common.Address `json:"token0"`
	token1 common.Address `json:"token1"`
	pair   common.Address `json:"pair"`
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6383820),
		ToBlock:   big.NewInt(6383840),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.PancakeswapABI)))
	if err != nil {
		log.Fatal(err)
	}

	pairCreatedSig := []byte("PairCreated(address,address,address,uint256)")
	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	pairCreatedSigHash := crypto.Keccak256Hash(pairCreatedSig)
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			_, err := contractAbi.Unpack("Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			_, err := contractAbi.Unpack("Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		case pairCreatedSigHash.Hex():
			fmt.Printf("Log Name: PairCreated\n")

			out, err := contractAbi.Unpack("PairCreated", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			pairCreated := PairCreated{}

			pairCreated.pair = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

			pairCreated.token0 = common.HexToAddress(vLog.Topics[1].Hex())
			pairCreated.token1 = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("token0: %s\n", pairCreated.token0.Hex())
			fmt.Printf("token1: %s\n", pairCreated.token1.Hex())
			fmt.Printf("pair: %s\n", pairCreated.pair.Hex())
		}

		fmt.Printf("\n\n")
	}

}
