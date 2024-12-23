package Pockyt

// QRCodeRequest 定义创建二维码请求的结构体
type QRCodeRequest struct {
	MerchantNo     string `json:"merchantNo"`
	StoreNo        string `json:"storeNo"`
	StoreAdminNo   string `json:"storeAdminNo,omitempty"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	SettleCurrency string `json:"settleCurrency"`
	Vendor         string `json:"vendor"`
	Reference      string `json:"reference,omitempty"`
	IpnUrl         string `json:"ipnUrl"`
	NeedQrcode     string `json:"needQrcode"`
	Timeout        string `json:"timeout,omitempty"`
	VerifySign     string `json:"verifySign"`
}
