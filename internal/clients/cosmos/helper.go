package cosmos

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethermintTypes "github.com/evmos/ethermint/types"
	"strings"
)

func initSdkConfig(cfg CosmosClientConfig) {
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetPurpose(44)
	sdkConfig.SetCoinType(ethermintTypes.Bip44CoinType)

	bech32PrefixAccAddr := fmt.Sprintf("%v", cfg.Prefix)
	bech32PrefixAccPub := fmt.Sprintf("%vpub", cfg.Prefix)
	bech32PrefixValAddr := fmt.Sprintf("%vvaloper", cfg.Prefix)
	bech32PrefixValPub := fmt.Sprintf("%vvaloperpub", cfg.Prefix)
	bech32PrefixConsAddr := fmt.Sprintf("%vvalcons", cfg.Prefix)
	bech32PrefixConsPub := fmt.Sprintf("%vvalconspub", cfg.Prefix)

	sdkConfig.SetBech32PrefixForAccount(bech32PrefixAccAddr, bech32PrefixAccPub)
	sdkConfig.SetBech32PrefixForValidator(bech32PrefixValAddr, bech32PrefixValPub)
	sdkConfig.SetBech32PrefixForConsensusNode(bech32PrefixConsAddr, bech32PrefixConsPub)

	return
}

func makeGrpcURL(url string) string {
	url = strings.Replace(url, "https://", "", -1)
	url = strings.Replace(url, "http://", "", -1)
	return url
}
