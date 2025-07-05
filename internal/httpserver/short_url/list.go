package short_url

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/short-url/internal/http"
	"github.com/kangyueyue/short-url/internal/models"
)

// List godoc
// @Summary 获取短链列表
// @Description 根据client_id获取对应的短链列表
// @Tags 短链接
// @Accept json
// @Produce json
// @Param client_id header string true "客户端ID"
// @Success 200 {array} models.PShortUrlData "短链列表"
// @Failure 400 {object} http.Response "请求参数错误"
// @Failure 500 {object} http.Response "服务器内部错误"
// @Router /short_url/list [get]
func (s *ShortUrlSvr) List(c *gin.Context) {
	// 从请求头中获取client_id
	client_id := c.GetHeader("client_id")
	db := s.store.GetDB().WithContext(c)
	var biz []*models.PShortUrlData
	if err := db.Where("client_id = ?", client_id).Find(&biz).Error; err != nil {
		http.Fail(c, "短链列表获取失败")
		return
	}
	http.Success(c, biz)
}
