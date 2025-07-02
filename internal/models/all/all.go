package all

import "github.com/kangyueyue/short-url/internal/models"

// Tables 所有表
func Tables() []interface{} {
	return []interface{}{
		&models.PShortUrlData{},
	}
}
