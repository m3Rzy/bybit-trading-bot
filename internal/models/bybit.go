package models

type Bybit struct {
	Result struct {
		List []struct {
			IndexPrice string `json:"indexPrice"`
		} `json:"list"`
	} `json:"result"`
}