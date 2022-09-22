package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		fmt.Println(err)
		return
	}

	address := common.HexToAddress("0xE0b1EF69BC4AB4173989C1190f0d77A813f3B726")
	token, err := NewToken(address, client)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalSupply, err := token.TotalSupply(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
		return
	}

	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
		return
	}

	s := totalSupply.String()
	fmt.Println(s[:len(s)-int(decimals)] + "." + s[len(s)-int(decimals):])

}
