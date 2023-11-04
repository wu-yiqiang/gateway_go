package dao

type HttpRule struct {
	ID             int64  `json:"id" grom:"id"`
	ServiceId      int64  `json:"service_id" grom:"service_id"`
	RuleType       int64  `json:"rule_type" grom:"rule_type"`
	Rule           string `json:"rule" grom:"rule"`
	NeedHttps      int64  `json:"need_https" grom:"need_https"`
	NeedStripUri   string `json:"need_strip_uri" grom:"need_strip_uri"`
	NeedWebsocket  int64  `json:"need_websocket" grom:"need_websocket"`
	UrlRewrite     string `json:"url_rewrite" grom:"url_rewrite"`
	HeaderTransfor string `json:"header_transfor" grom:"header_transfor"`
}
