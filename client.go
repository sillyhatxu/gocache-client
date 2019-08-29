package client

import (
	gocache "github.com/patrickmn/go-cache"
	"time"
)

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

type CacheConfig struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
	client            *gocache.Cache
}

func NewCacheConfig(defaultExpiration, cleanupInterval time.Duration) *CacheConfig {
	return &CacheConfig{
		DefaultExpiration: defaultExpiration,
		CleanupInterval:   cleanupInterval,
		client:            gocache.New(defaultExpiration, cleanupInterval),
	}
}

func (cc CacheConfig) getClient() *gocache.Cache {
	if cc.client == nil {
		return gocache.New(cc.DefaultExpiration, cc.CleanupInterval)
	}
	return cc.client
}

func (cc CacheConfig) Set(key string, value interface{}, d time.Duration) {
	cc.getClient().Set(key, value, d)
}

func (cc CacheConfig) Get(key string) (interface{}, bool) {
	value, found := cc.getClient().Get(key)
	return value, found
}

func (cc CacheConfig) Delete(key string) {
	cc.getClient().Delete(key)
}
