package common

import (
	"fmt"
	internalCommon "github.com/LampardNguyen234/whale-alert/internal/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// FormatMDLink creates a ref-text message following the MarkDown standard.
func FormatMDLink(msg string, link string) string {
	return fmt.Sprintf("[%v](%v)", msg, link)
}

func AccountAddressToHex(addr string) (string, error) {
	ethAddr, err := AccountAddressToEthAddr(addr)
	if err != nil {
		return "", err
	}

	return ethAddr.Hex(), nil
}

// MustAccountAddressToHex is the same as AccountAddressToHex except that it will panic upon errors.
func MustAccountAddressToHex(addr string) string {
	addr, err := AccountAddressToHex(addr)
	if err != nil {
		panic(err)
	}

	return addr
}

// AccountAddressToEthAddr parses the given address to an ETH address.
func AccountAddressToEthAddr(addr string) (common.Address, error) {
	zeroAddr := common.HexToAddress(addr)
	if addr == internalCommon.ZeroAddress {
		return zeroAddr, nil
	}
	if strings.HasPrefix(addr, sdk.GetConfig().GetBech32AccountAddrPrefix()) {
		// Check to see if address is Cosmos bech32 formatted
		toAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return zeroAddr, fmt.Errorf("%v is not a valid Bech32 address", addr)
		}
		ethAddr := common.BytesToAddress(toAddr.Bytes())
		return ethAddr, nil
	}

	if !strings.HasPrefix(addr, "0x") {
		addr = "0x" + addr
	}

	valid := common.IsHexAddress(addr)
	if !valid {
		return zeroAddr, fmt.Errorf("%s is not a valid Ethereum or Cosmos address", addr)
	}

	ethAddr := common.HexToAddress(addr)

	return ethAddr, nil
}
