package cli

import (
	block2 "blockchain_go/core/block"
	"blockchain_go/utils"
	"fmt"
	"log"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !utils.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := block2.NewBlockchain(nodeID)
	UTXOSet := utils.UTXOSet{bc}
	defer bc.db.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
