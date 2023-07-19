package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var defaultCache = cache.New(5*time.Minute, 10*time.Minute)

func Get[T any](key string) (T, bool, func(T)) {
	var zeroT T

	v, ok := defaultCache.Get(key)
	if !ok {
		return zeroT, false, nil
	}

	vt, ok := v.(T)
	if !ok {
		return zeroT, false, nil
	}

	return vt, true, func(v T) {
		defaultCache.SetDefault(key, v)
	}
}
