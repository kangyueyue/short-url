package redirect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	http2 "github.com/kangyueyue/short-url/internal/http"
	"github.com/kangyueyue/short-url/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// Redirect 重定向
func (r *RedirectSvr) Redirect(c *gin.Context) {
	shortUrl := c.Param("id")
	db := r.store.GetDB().WithContext(c)
	var bizs []*models.PShortUrlData
	// 查询所有匹配记录并按创建时间降序排序
	if err := db.Where("short_url = ?", shortUrl).
		Order("created_at DESC"). // 按创建时间降序
		Find(&bizs).Error; err != nil {
		http2.Fail(c, "短链不存在: "+err.Error())
		return
	}
	// 检查是否找到记录
	if len(bizs) == 0 {
		http2.Fail(c, "短链不存在")
		return
	}
	if len(bizs) > 1 {
		fmt.Print("Found multiple records for short URL: %s", shortUrl)
		logrus.Warnf("Found multiple records for short URL: %s", shortUrl)
	}
	// 取最近的一条记录（排序后的第一条）
	latestBiz := bizs[0]
	if !strings.HasPrefix(latestBiz.LongUrl, "http") {
		latestBiz.LongUrl = "http://" + latestBiz.LongUrl
	}
	// 重定向
	c.Redirect(http.StatusFound, latestBiz.LongUrl)
}
