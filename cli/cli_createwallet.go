package cli

import (
	"blockchain_go/utils"
	"fmt"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := utils.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
