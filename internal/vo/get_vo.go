package vo

// GetVO 获取
type GetVO struct {
	ShortUrl string `json:"shortUrl" binding:"required"`
}
