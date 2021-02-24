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

	"github.com/mehmetcantas/gopos/netspay_provider"
	"github.com/mehmetcantas/gopos/models"
	"github.com/mehmetcantas/gopos/models/card_type"
	"github.com/mehmetcantas/gopos/models/currency"
)

func main() {

	customerBuilder := models.NewCustomerBuilder()

	customerBuilder.NameIs("Mehmet Can Taş").
		EmailIs("tass.mehmetcan@outlook.com").
		IpAddress("127.0.0.1").
		IsCompany(false).
		ShipTo("test shipping address", "12354").
		BillTo("test billing address", "123123").
		WithID("1234432")
	customer := customerBuilder.Build()

	paymentBuilder := models.NewPaymentRequestBuilder()
	paymentBuilder.
		Card("Mehmet Can Taş", "4355084355084358", "000").
		Type(card_type.Visa).
		ExpireAt("12", "26").
		Currency(currency.TRY).
		Language("TR").
		WithInstallment(1).
		ToCustomer(customer).
		ForOrder("23425423", 125.54).
		InSuccessReturns("http://localhost:8090/verify").
		InFailReturns("http://localhost:8090/verify")

	req := paymentBuilder.Build()

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

### TODO  

- [ ] Response objelerinin düzenlenmesi
- [ ] Eksik hata kontrollerinin yapılması
- [ ] Metot açıklamalarının eklenmesi
- [ ] Bin numaralarının eklenmesi
- [ ] Ödeme iptali
- [ ] İade
- [ ] İyzico entegrasyonu
- [ ] Sofort entegrasyonu
- [ ] UnionPay entegrasyonu
- [ ] Paypal entegrasyonu





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
