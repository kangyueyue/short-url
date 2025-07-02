package short_url

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/short-url/internal/http"
	"github.com/kangyueyue/short-url/internal/models"
	"github.com/kangyueyue/short-url/internal/vo"
	"time"
)

// @Summary 创建短链接
// @Description 批量创建短链接，将长链接转换为短链接
// @Tags 短链接
// @Accept json
// @Produce json
// @Param request body vo.SetVo true "请求参数"
// @Success 200 {object} http.Response{data=[]models.PShortUrlData} "成功返回短链接信息"
// @Failure 500 {object} http.Response "服务器内部错误"
// @Router /short_url/set [post]
func (s *ShortUrlSvr) Set(c *gin.Context) {
	// 从请求中获取body
	body := &vo.SetVo{}
	if err := c.ShouldBindJSON(body); err != nil {
		http.Fail(c, "参数错误")
		return
	}
	// 创建短链
	var bizs []*models.PShortUrlData
	for _, longUrl := range body.LongUrls {
		// TODO 生成short url
		biz := &models.PShortUrlData{
			ShortUrlData: &models.ShortUrlData{
				LongUrl:  longUrl,
				ShortUrl: "aaabbb",
				ClientId: body.ClientId,
				ExpireAt: time.Now().AddDate(0, 0, 1), // 一年以后过期
			},
		}
		bizs = append(bizs, biz)
	}
	db := s.store.GetDB().WithContext(c)
	if err := db.CreateInBatches(bizs, len(bizs)).Error; err != nil { // 注意添加 .Error
		http.Fail(c, "存储短链失败: "+err.Error())
		return
	}
	http.Success(c, bizs)
}
