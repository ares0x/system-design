package service

import (
	"short-url/internal/constvar"
	"short-url/internal/data"
	"strings"
	"time"
)

type UrlService struct{}

func NewUrlService() *UrlService {
	return new(UrlService)
}

// Encode 将长链接转为短链接
func (u *UrlService) Encode(originalUrl string, tm int) (string, error) {
	// 根据某种算法获取短链接
	// 将长链接，短链接存储，设置过期时间
	uuid := time.Now().UnixNano() // uuid 可以通过多种方式生成，如，redis incr，独立的发号器等
	var shortUrl string
	for uuid > 0 {
		shortUrl = string(constvar.Seed[uuid%62]) + shortUrl
		uuid = uuid / 62
	}
	url := &data.Url{
		Id:          uuid,
		OriginalUrl: originalUrl,
		ShortUrl:    shortUrl,
	}
	urlQuery := data.New(nil)
	if err := urlQuery.Create(url); err != nil {
		return "", err
	}
	return shortUrl, nil
}

// Decode 将短链接转为长链接
func (u *UrlService) Decode(shortUrl string) (string, error) {
	// 根据短链接获取
	var id int64
	n := len(shortUrl)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(constvar.Seed, shortUrl[i])
		id = id*62 + int64(pos)
	}
	urlQuery := data.New(nil)
	url, err := urlQuery.GetById(id)
	if err != nil {

	}
	return url.OriginalUrl, nil
}
