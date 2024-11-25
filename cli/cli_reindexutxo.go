package cli

import (
	block2 "blockchain_go/core/block"
	"blockchain_go/utils"
	"fmt"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := block2.NewBlockchain(nodeID)
	UTXOSet := utils.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
