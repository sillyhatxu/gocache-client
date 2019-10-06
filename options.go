package cacheclient

import (
	"github.com/allegro/bigcache"
	"time"
)

type Config struct {
	shards             int
	lifeWindow         time.Duration
	cleanWindow        time.Duration
	maxEntriesInWindow int
	maxEntrySize       int
	verbose            bool
	hardMaxCacheSize   int
	logger             bigcache.Logger
}

//type Config struct {
//	shards             int
//	lifeWindow         time.Duration
//	cleanWindow        time.Duration
//	maxEntriesInWindow int
//	maxEntrySize       int
//	verbose            bool
//	hasher             bigcache.Hasher
//	hardMaxCacheSize   int
//	onRemove           func(key string, entry []byte)
//	onRemoveWithReason func(key string, entry []byte, reason bigcache.RemoveReason)
//	onRemoveFilter     int
//	logger             bigcache.Logger
//}

type Option func(*Config)

func Shards(shards int) Option {
	return func(c *Config) {
		c.shards = shards
	}
}
