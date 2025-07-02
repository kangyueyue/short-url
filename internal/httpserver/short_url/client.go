package short_url

import "github.com/kangyueyue/short-url/internal/infrastructure/store"

// ShortUrlSvr 短网址服务
type ShortUrlSvr struct {
	store *store.Store
}

// NewShortUrlSvr 创建短网址服务
func NewShortUrlSvr(store *store.Store) *ShortUrlSvr {
	return &ShortUrlSvr{
		store: store,
	}
}
