package cosmos

import (
	"fmt"
	"testing"
)

func TestTxsClient_TxByHash(t *testing.T) {
	resp, err := c.TxByHash("716E7391B71DFFD6059D645131647D2F2526B26887D7C7CD37B704A3A2136208")
	if err != nil {
		panic(err)
	}

	c.log.Infof(fmt.Sprintf("%v", resp))
}
