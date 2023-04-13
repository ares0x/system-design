package model

type (
	SetOriginalUrlReq struct {
		OriginalUrl string `json:"originalUrl"`
		Expire      int    `json:"expire"` // 过期时间，秒（0表示不过期）
	}

	SetOriginalUrlReply struct {
		ShortUrl string `json:"shortUrl"`
	}
)
