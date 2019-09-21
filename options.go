package client

import "time"

type Config struct {
	expiration      time.Duration
	cleanupInterval time.Duration
}

type Option func(*Config)

func Expiration(expiration time.Duration) Option {
	return func(c *Config) {
		c.expiration = expiration
	}
}

func CleanupInterval(cleanupInterval time.Duration) Option {
	return func(c *Config) {
		c.cleanupInterval = cleanupInterval
	}
}
