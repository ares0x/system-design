package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short-url/internal/router/model"
	"short-url/internal/service"
	"short-url/pkg/errno"
)

type UrlHandler struct{}

func NewUrlHandler() *UrlHandler {
	return new(UrlHandler)
}

// SetOriginalUrl 将长链接转换为短链接
func (u *UrlHandler) SetOriginalUrl(c *gin.Context) {
	param := new(model.SetOriginalUrlReq)
	if err := c.ShouldBind(param); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	shortUrl, err := service.NewUrlService().Encode(param.OriginalUrl, param.Expire)
	if err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	}
	reply := &model.SetOriginalUrlReply{
		ShortUrl: shortUrl,
	}

	SendResponse(c, nil, reply)
}

// RedirectToOriginalUrl 重定向到长链接
func (u *UrlHandler) RedirectToOriginalUrl(c *gin.Context) {
	shotUrl := c.Param("shorturl")
	if shotUrl == "" {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	originalUrl, err := service.NewUrlService().Decode(shotUrl)
	if err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	c.Redirect(http.StatusMovedPermanently, originalUrl)
}
