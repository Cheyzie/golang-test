package cache

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemcacheConfig struct {
	Host string
	Port string
}

func NewMemcacheClient(cfg MemcacheConfig) *memcache.Client {
	return memcache.New(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
}
