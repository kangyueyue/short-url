package models

import "time"

// ShortUrlData 内部数据定义
type ShortUrlData struct {
	// 长链-使用utf8mb4_bin字符集，区分大小写
	LongUrl string `gorm:"type:varchar(8182) COLLATE utf8mb4_bin;not null;comment:长链接"`
	// 短链-使用utf8mb4_bin字符集，区分大小写
	ShortUrl string `gorm:"type:varchar(14) COLLATE utf8mb4_bin;not null;comment:短链接"`
	// 用户id
	ClientId string `gorm:"type:varchar(255);not null;comment:用户id"`
	// 过期时间
	ExpireAt time.Time `gorm:"type:datetime;not null;index;comment:过期时间"`
}

// PShortUrlData db
type PShortUrlData struct {
	Model
	*ShortUrlData `gorm:"embedded"`
}
