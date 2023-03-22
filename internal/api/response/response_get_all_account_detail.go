package response

import "github.com/LampardNguyen234/whale-alert/internal/store"

// APIAllAccountDetailResponse holds all stored account details.
type APIAllAccountDetailResponse struct {
	Result []store.AccountDetail `json:"result"`
}
