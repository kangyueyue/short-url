package short_url

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/short-url/internal/http"
	"github.com/kangyueyue/short-url/internal/models"
	"github.com/kangyueyue/short-url/internal/vo"
)

// @Summary 获取长链接
// @Description 根据短链接获取原始长链接
// @Tags 短链接
// @Accept json
// @Produce json
// @Param request body vo.GetVO true "请求参数"
// @Success 200 {object} http.Response{data=models.PShortUrlData} "成功返回长链接信息"
// @Failure 500 {object} http.Response "服务器内部错误"
// @Router /short_url/get [post]
func (s *ShortUrlSvr) Get(c *gin.Context) {
	// 从请求中获取body
	body := &vo.GetVO{}
	if err := c.ShouldBindJSON(body); err != nil {
		http.Fail(c, "参数错误")
		return
	}
	// 获取短链
	biz := &models.PShortUrlData{}
	db := s.store.GetDB().WithContext(c)
	if err := db.Where("short_url = ? AND client_id = ?", body.ShortUrl, body.ClientId).First(biz).Error; err != nil {
		http.Fail(c, "短链不存在")
		return
	}
	http.Success(c, biz)
}
