package vo

// SetVo 设置短链请求
type SetVo struct {
	LongUrls []string `json:"longUrls" binding:"required"`
	ClientId string   `json:"clientId" binding:"required"`
}
