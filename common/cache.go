package common

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	defaultExpiration      = 10 * time.Minute
	defaultCleanUpInterval = 60 * time.Minute
)

// Cache is a simple cache.
type Cache interface {
	SetDefault(string, interface{})
	Set(string, interface{}, time.Duration)
	Get(string) (interface{}, bool)
	GetTo(string, interface{}) error
	Remove(string)
	RemoveExpired()
}

type SimpleCache struct {
	*cache.Cache
}

// NewSimpleCache creates a new SimpleCache with default expiration and cleanUpInterval.
func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		Cache: cache.New(cache.NoExpiration, defaultCleanUpInterval),
	}
}

func (c *SimpleCache) Set(k string, v interface{}, expired time.Duration) {
	c.Cache.Set(k, v, expired)
}

func (c *SimpleCache) SetDefault(k string, v interface{}) {
	c.Cache.SetDefault(k, v)
}

func (c *SimpleCache) Get(k string) (interface{}, bool) {
	return c.Cache.Get(k)
}

func (c *SimpleCache) GetTo(k string, ret interface{}) error {
	tmpRet, exist := c.Get(k)
	if !exist {
		return fmt.Errorf("item not existed")
	}

	jsb, err := json.Marshal(tmpRet)
	if err != nil {
		return fmt.Errorf("failed to marshal result")
	}
	err = json.Unmarshal(jsb, &ret)
	if err != nil {
		return fmt.Errorf("failed to read result")
	}

	return nil
}

func (c *SimpleCache) Remove(k string) {
	c.Cache.Delete(k)
}

func (c *SimpleCache) RemoveExpired() {
	c.Cache.DeleteExpired()
}
