# gopos
[GO] Asseco NestPay (EST) (İş Bankası, Akbank, Finansbank, Halkbank, Anadolubank, Citibank, Ziraat Bankası, ING Bank), Garanti BBVA, POSNET (Yapı Kredi) sanal pos entegrasyonları için GO paketi.



Paketi projenize eklemeden önce aşağıdaki kodu çalıştırınız.

```bash
go get "github.com/mehmetcantas/gopos"
```

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

### NETSPAY Sanal Pos Entegrasyon Örneği

<details>
<summary>3D isteği oluşturma</summary>

Netspay sanal pos entegrasyonu için öncelikle gerekli bilgileri girerek 3D ekranına post edilecek formun oluşturulması sağlanmalıdır.
Aşağıdaki örnekte yer alan test bilgileri ile 3D ekranına post edilecek olan formu oluşturabilirsiniz.
Geriye dönen HTML form içeriğini herhangi bir HTML sayfasına eklediğinizde otomatik olarak 3D doğrulama sayfasına yönleneceksiniz.

```go
package main

import (
	"fmt"

	"github.com/mehmetcantas/gopos"
	"github.com/mehmetcantas/gopos/models"
)

func main() {
	req := models.PaymentGatewayRequest{
		CardHolderName:       "Mehmet Can Taş",
		CardNumber:           "4355084355084358", // test kredi kartı
		ExpireMonth:          "12",
		ExpireYear:           "26",
		CVV:                  "000",
		CustomerEmailAddress: "tass.mehmetcan@outlook.com",
		CompanyName:          "DOGO TASARIM SAN. ve TİC. A.Ş.",
		OrderNumber:          "12343242",
		OrderTotal:           142.54,
		InstallmentCount:     1,
		UserID:               "12312414",
		CurrencyCode:         "TL",
		LanguageCode:         "tr",
		CustomerIPAddress:    "127.0.0.1",
		CardType:             "VISA",
		SuccessURL:           "http://localhost:8090/netspay/verify",
		FailURL:              "http://localhost:8090/netspay/verify",
	}

	var netspay = netspay_provider.Netspay{
		UseSandbox:                 true,
		MerchantID:                 "100200000",
		StoreKey:                   "123456",
		BankName:                   "Akbank",
		ApiURL:                     "https://entegrasyon.asseco-see.com.tr/fim/est3Dgate",
		SecurityType:               "3D_PAY",
		UseManufacturerCardSupport: false,
	}
	res, _ := netspay.PreparePaymentGatewayForm(&req)
	
	fmt.Println(res)
}


```
 </details>

MIT License

Copyright (c) 2021 Mehmet Can Taş

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
