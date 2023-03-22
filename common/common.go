package common

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// FormatMDLink creates a ref-text message following the MarkDown standard.
func FormatMDLink(msg string, link string) string {
	return fmt.Sprintf("[%v](%v)", msg, link)
}

func AccountAddressToHex(addr string) (string, error) {
	if strings.HasPrefix(addr, sdk.GetConfig().GetBech32AccountAddrPrefix()) {
		// Check to see if address is Cosmos bech32 formatted
		toAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return "", fmt.Errorf("%v is not a valid Bech32 address", addr)
		}
		ethAddr := common.BytesToAddress(toAddr.Bytes())
		return ethAddr.Hex(), nil
	}

	if !strings.HasPrefix(addr, "0x") {
		addr = "0x" + addr
	}

	valid := common.IsHexAddress(addr)
	if !valid {
		return "", fmt.Errorf("%s is not a valid Ethereum or Cosmos address", addr)
	}

	ethAddr := common.HexToAddress(addr)

	return ethAddr.Hex(), nil
}

// MustAccountAddressToHex creates
func MustAccountAddressToHex(addr string) string {
	addr, err := AccountAddressToHex(addr)
	if err != nil {
		return addr
	}

	return addr
}
