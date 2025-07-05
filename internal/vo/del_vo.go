package vo

// DelVo 删除短链接VO
type DelVo struct {
	ShortUrl string `json:"shortUrl" binding:"required"`
	IsDisk   bool   `json:"isDisk"` // 是否删除磁盘文件
}
