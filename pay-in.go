package Pockyt

import (
	"fmt"
)

// CreateQRCode 创建二维码的函数，接收参数
func CreateQRCode(merchantNo, storeNo, storeAdminNo, amount, currency, settleCurrency, vendor, reference, ipnUrl, needQrcode, timeout, verifySign string) {
	url := "https://mapi.yuansfer.com/app-instore/v3/create-trans-qrcode"

	payload := QRCodeRequest{
		MerchantNo:     merchantNo,
		StoreNo:        storeNo,
		StoreAdminNo:   storeAdminNo,
		Amount:         amount,
		Currency:       currency,
		SettleCurrency: settleCurrency,
		Vendor:         vendor,
		Reference:      reference,
		IpnUrl:         ipnUrl,
		NeedQrcode:     needQrcode,
		Timeout:        timeout,
		VerifySign:     verifySign,
	}

	response, err := SendRequest(url, payload)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response)
}
