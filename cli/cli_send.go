package cli

import (
	block2 "blockchain_go/core/block"
	"blockchain_go/core/net"
	"blockchain_go/core/transaction"
	"blockchain_go/utils"
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !utils.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !utils.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := block2.NewBlockchain(nodeID)
	UTXOSet := utils.UTXOSet{bc}
	defer bc.db.Close()

	wallets, err := utils.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := transaction.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := transaction.NewCoinbaseTX(from, "")
		txs := []*transaction.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		net.sendTx(net.knownNodes[0], tx)
	}

	fmt.Println("Success!")
}
