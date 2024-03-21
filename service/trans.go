package service

import (
	"anycat/global/consts"
	"anycat/global/variable"
	"anycat/model"
	"anycat/util/httputils"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func Trans(req model.TransReq) (*model.TransRes, error) {
	q := req.Content
	from := req.From
	to := req.To

	signStr := consts.BaiduAppId + q + consts.BaiduSalt + consts.BaiduSecret
	hash := md5.Sum([]byte(signStr))
	sign := hex.EncodeToString(hash[:])

	r := httputils.NewClientHandle()
	payload := map[string]string{
		"q":     q,
		"from":  from,
		"to":    to,
		"appid": consts.BaiduAppId,
		"salt":  consts.BaiduSalt,
		"sign":  sign,
	}

	res, err := r.Client.R().SetHeader("Content-Type", "application/x-www-form-urlencoded").SetFormData(payload).Post(consts.BaiduTransApi)

	if err != nil {
		variable.ZapLog.Errorf("baidu trans req failed, err: %v", err)
		return nil, err
	}

	if res.StatusCode() != 200 {
		variable.ZapLog.Errorf("baidu trans req failed, status: %d", res.StatusCode())
		return nil, err
	}

	resp := new(model.TransRes)
	err = json.Unmarshal(res.Body(), resp)

	if err != nil {
		variable.ZapLog.Errorf("baidu trans json unmarshal failed, err: %v", err)
		return nil, err
	}
	return resp, nil
}
