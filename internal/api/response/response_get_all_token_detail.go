package response

import "github.com/LampardNguyen234/whale-alert/internal/store"

// APIAllTokenDetailResponse holds all stored account details.
type APIAllTokenDetailResponse struct {
	Result map[string]store.TokenDetail `json:"result"`
}
