package redirect

import (
	"github.com/kangyueyue/short-url/internal/infrastructure/store"
)

// RedirectSvr 重定向服务
type RedirectSvr struct {
	store *store.Store
}

// NewRedirectSvr 创建重定向服务
func NewRedirectSvr(store *store.Store) *RedirectSvr {
	return &RedirectSvr{
		store: store,
	}
}
