package short_url

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/short-url/internal/http"
	"github.com/kangyueyue/short-url/internal/models"
	vo2 "github.com/kangyueyue/short-url/internal/vo"
)

// Del godoc
// @Summary 删除短链
// @Description 根据短链标识和客户端ID删除对应的短链
// @Tags 短链接
// @Accept json
// @Produce json
// @Param client_id header string true "客户端ID"
// @Param request body vo2.DelVo true "删除请求参数"
// @Success 200 "删除成功"
// @Failure 500 {object} http.Response "服务器内部错误"
// @Router /short_url/del [delete]
func (s *ShortUrlSvr) Del(c *gin.Context) {
	client_id := c.GetHeader("client_id")
	var delVo *vo2.DelVo
	if err := c.ShouldBindJSON(&delVo); err != nil {
		http.Fail(c, "参数错误")
		return
	}
	db := s.store.GetDB().WithContext(c)
	if delVo.IsDisk {
		// 从磁盘删除
		if err := db.Where("short_url = ? AND client_id = ?", delVo.ShortUrl, client_id).Unscoped().Delete(&models.PShortUrlData{}).Error; err != nil {
			http.Fail(c, "删除失败")
			return
		}
	} else {
		// 软删除
		if err := db.Where("client_id = ?", client_id).Where("short_url = ?", delVo.ShortUrl).Delete(&models.PShortUrlData{}).Error; err != nil {
			http.Fail(c, "删除失败")
			return
		}
	}
	http.Success(c, "删除成功")
}
