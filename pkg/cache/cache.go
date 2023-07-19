package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var defaultCache = cache.New(5*time.Minute, 10*time.Minute)

func Default() *cache.Cache {
	return defaultCache
}
