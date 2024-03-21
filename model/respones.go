package model

type TransRes struct {
	From        string         `json:"from"`
	To          string         `json:"to"`
	TransResult []*TransResult `json:"trans_result"`
}

type TransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
