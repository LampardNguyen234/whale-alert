package tiki

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *TikiClient) parseResponse(resp *http.Response, ret interface{}) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with code: %v", resp.StatusCode)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			c.log.Errorf("Close body error: %v", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return err
	}

	return nil
}

func (c *TikiClient) parseUrl(shortPath string) string {
	return fmt.Sprintf("%v/%v", c.cfg.Endpoint, shortPath)
}
