package store

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	"testing"
)

func TestStore_GetTokenDetail(t *testing.T) {
	d := s.GetTokenDetail(common.ZeroAddress)
	fmt.Println(d)
}
