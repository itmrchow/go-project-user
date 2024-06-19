package respdto

type ApiErrorResp struct {
	// Type     string `json:"type,omitempty"`
	// Status   int    `json:"status,omitempty"`   // http status
	Title  string `json:"title,omitempty"`  // 可讀的標示符
	Detail string `json:"detail,omitempty"` // 問題描述
	// Error  map[string]interface{} `json:"Error,omitempty"`  // 內部標示符 -> 對應到內部error的相關info
}
