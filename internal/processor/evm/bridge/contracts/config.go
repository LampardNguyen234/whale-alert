package contracts

import "fmt"

type BridgeNetworkConfig struct {
	// Name is the name of the network.
	Name string `json:"Name"`

	// URL is the RPC endpoint.
	URL string `json:"URL"`

	// FromHeight specifies the block height from which to start listening for events.
	FromHeight uint64 `json:"FromHeight"`

	// BlockInterval specifies the number of blocks for each log retrieval attempt.
	BlockInterval uint64 `json:"BlockInterval"`

	// Bridge is the address of the bridge contract.
	Bridge string `json:"Bridge"`

	// Explorer is the URL of the explorer.
	Explorer string `json:"Explorer"`
}

// IsValid checks if the current BridgeNetworkConfig is sanity-valid.
func (cfg *BridgeNetworkConfig) IsValid() (bool, error) {
	if cfg.Name == "" {
		return false, fmt.Errorf("empty cfg.Name")
	}

	if cfg.URL == "" {
		return false, fmt.Errorf("empty URL")
	}

	if cfg.BlockInterval == 0 {
		return false, fmt.Errorf("invalid cfg.BlockInterval: must not be 0")
	}

	if cfg.Bridge == "" {
		return false, fmt.Errorf("empty cfg.bridge")
	}

	return true, nil
}
