package httputil

import (
	"anycat/global/consts"
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func TestTranslate(t *testing.T) {
	appid := consts.BaiduAppId
	secret := consts.BaiduSecret
	q := "hello"
	from := "en"
	to := "zh"
	salt := "anycat"

	signStr := appid + q + salt + secret
	hash := md5.Sum([]byte(signStr))
	sign := hex.EncodeToString(hash[:])

	r := NewClientHandle()
	payload := map[string]string{
		"q":     q,
		"from":  from,
		"to":    to,
		"appid": appid,
		"salt":  salt,
		"sign":  sign,
	}

	res, err := r.Client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded").SetFormData(payload).Post(consts.BaiduTransApi)

	if err != nil {
		t.Error(err)
	}
	t.Log(res)

}
