package common

import (
	"fmt"
	"strings"
)

const (
	EvmClientName    = "EVM"
	CosmosClientName = "COSMOS"

	EvmClientID    = 0
	CosmosClientID = 1
)

var Explorer ExplorerDetail

func InitExplorer(chainID string) {
	if strings.Contains(chainID, "11110") {
		Explorer = ExplorerDetail{
			ExplorerUrl: "https://explorer.astranaut.io",
			PingPubURL:  "https://ping.astranaut.io/astra",
		}
	}

	Explorer = ExplorerDetail{
		ExplorerUrl: "https://explorer.astranaut.dev",
		PingPubURL:  "https://ping.astranaut.dev/astra",
	}
}

type ExplorerDetail struct {
	ExplorerUrl string
	PingPubURL  string
}

func (d ExplorerDetail) FormatTxURL(txHash string) string {
	return fmt.Sprintf("%v/tx/%v", d.ExplorerUrl, txHash)
}

func (d ExplorerDetail) FormatAccountURL(accountAddr string) string {
	return fmt.Sprintf("%v/address/%v", d.ExplorerUrl, accountAddr)
}

func (d ExplorerDetail) FormatValidatorURL(valAddr string) string {
	if d.PingPubURL == "" {
		return fmt.Sprintf("%v/address/%v", d.ExplorerUrl, valAddr)
	}

	return fmt.Sprintf("%v/staking/%v", d.PingPubURL, valAddr)
}
