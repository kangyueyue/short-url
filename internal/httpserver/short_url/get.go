package short_url

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/short-url/internal/http"
	"github.com/kangyueyue/short-url/internal/models"
	"github.com/kangyueyue/short-url/internal/vo"
)

// Get godoc
// @Summary 获取长链接
// @Description 根据短链接获取原始长链接
// @Tags 短链接
// @Accept json
// @Produce json
// @Param request body vo.GetVO true "请求参数"
// @Param client_id header string true "客户端ID"
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
	client_id := c.GetHeader("client_id")
	if client_id == "" {
		http.Fail(c, "缺少请求头")
		return
	}
	filters := body.Filter
	db := s.store.GetDB().WithContext(c).Where("client_id = ?", client_id)
	for _, filter := range filters {
		db = db.Where(filter.Field+" "+filter.Operator+" ?", filter.Value)
	}
	// 获取短链
	var biz *models.PShortUrlData
	if err := db.First(&biz).Error; err != nil {
		http.Fail(c, "短链不存在")
		return
	}
	http.Success(c, biz)
}
