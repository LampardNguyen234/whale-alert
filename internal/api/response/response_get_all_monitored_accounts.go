package response

import "github.com/LampardNguyen234/whale-alert/internal/store"

// APIAllMonitoredAccountResponse holds all monitored account.
type APIAllMonitoredAccountResponse struct {
	Result map[string]*store.AccountDetail `json:"result"`
}
