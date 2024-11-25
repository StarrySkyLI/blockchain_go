package cli

import (
	block2 "blockchain_go/core/block"
	"blockchain_go/utils"
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !utils.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := block2.CreateBlockchain(address, nodeID)
	defer bc.db.Close()

	UTXOSet := utils.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
