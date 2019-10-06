package cacheclient

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

const (
	NoExpiration           = cache.NoExpiration
	DefaultExpiration      = cache.DefaultExpiration
	defaultExpiration      = 24 * time.Hour
	defaultCleanupInterval = 25 * time.Hour
)

type CacheClient struct {
	config *Config
	client *cache.Cache
	mu     sync.Mutex
}

func NewCacheClient(opts ...Option) *CacheClient {
	//default
	config := &Config{
		expiration:      defaultExpiration,
		cleanupInterval: defaultCleanupInterval,
	}
	for _, opt := range opts {
		opt(config)
	}
	return &CacheClient{
		config: config,
		client: cache.New(config.expiration, config.cleanupInterval),
	}
}

func (cc CacheClient) getClient() *cache.Cache {
	if cc.client == nil {
		return cache.New(cc.config.expiration, cc.config.cleanupInterval)
	}
	return cc.client
}

func (cc CacheClient) Get(key string) (interface{}, bool) {
	return cc.getClient().Get(key)
}

func (cc CacheClient) Delete(key string) {
	cc.getClient().Delete(key)
}

func (cc CacheClient) Set(key string, value interface{}) {
	cc.getClient().Set(key, value, NoExpiration)
}

func (cc CacheClient) IncrementInt(k string) (int, error) {
	return cc.IncrementIntWithExpiration(k, NoExpiration)
}

func (cc CacheClient) IncrementInt8(k string) (int8, error) {
	return cc.IncrementInt8WithExpiration(k, NoExpiration)
}

func (cc CacheClient) IncrementInt16(k string) (int16, error) {
	return cc.IncrementInt16WithExpiration(k, NoExpiration)
}

func (cc CacheClient) IncrementInt32(k string) (int32, error) {
	return cc.IncrementInt32WithExpiration(k, NoExpiration)
}

func (cc CacheClient) IncrementInt64(k string) (int64, error) {
	return cc.IncrementInt64WithExpiration(k, NoExpiration)
}

func (cc CacheClient) SetWithExpiration(key string, value interface{}, d time.Duration) {
	cc.getClient().Set(key, value, d)
}

func (cc CacheClient) IncrementIntWithExpiration(k string, d time.Duration) (int, error) {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	_, b := cc.Get(k)
	if !b {
		cc.SetWithExpiration(k, int(0), d)
	}
	return cc.getClient().IncrementInt(k, 1)
}

func (cc CacheClient) IncrementInt8WithExpiration(k string, d time.Duration) (int8, error) {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	_, b := cc.Get(k)
	if !b {
		cc.SetWithExpiration(k, int8(0), d)
	}
	return cc.getClient().IncrementInt8(k, 1)
}

func (cc CacheClient) IncrementInt16WithExpiration(k string, d time.Duration) (int16, error) {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	_, b := cc.Get(k)
	if !b {
		cc.SetWithExpiration(k, int16(0), d)
	}
	return cc.getClient().IncrementInt16(k, 1)
}

func (cc CacheClient) IncrementInt32WithExpiration(k string, d time.Duration) (int32, error) {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	_, b := cc.Get(k)
	if !b {
		cc.SetWithExpiration(k, int32(0), d)
	}
	return cc.getClient().IncrementInt32(k, 1)
}

func (cc CacheClient) IncrementInt64WithExpiration(k string, d time.Duration) (int64, error) {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	_, b := cc.Get(k)
	if !b {
		cc.SetWithExpiration(k, int64(0), d)
	}
	return cc.getClient().IncrementInt64(k, 1)
}
