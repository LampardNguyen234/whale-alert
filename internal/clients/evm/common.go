package evm

import "github.com/ethereum/go-ethereum/core/types"

func (c *EvmClient) MustGetEvmTxSender(tx *types.Transaction) string {
	signer, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		return ""
	}

	return signer.String()
}
