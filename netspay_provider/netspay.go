package netspay_provider

import (
	"errors"
	"fmt"
	"github.com/mehmetcantas/gopos/internal/utils"
	"github.com/mehmetcantas/gopos/models"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type Netspay struct {
	BankName     string
	StoreKey     string
	MerchantID   string
	UseSandbox   bool
	ApiURL       string
	SecurityType string // 3D, 3D_PAY, 3D_PAY_HOSTING vb.
}

func (n Netspay) PreparePaymentGatewayForm(r *models.PaymentGatewayRequest) (models.PaymentGatewayResponse, error) {
	var err error
	var cardCVV string

	if n.UseSandbox {
		cardCVV = "000"
	} else {
		cardCVV = r.CVV
	}
	if r.InstallmentCount < 0 {
		err = errors.New("Taksit adeti sıfırdan küçük olamaz")
		return models.PaymentGatewayResponse{}, err
	}

	installment := convertInstallment(r.InstallmentCount)
	rndNum := strconv.Itoa(int(time.Now().UnixNano()))[:10]
	hashFormat := n.MerchantID + r.OrderNumber + fmt.Sprintf("%.2f", r.OrderTotal) + r.SuccessURL + r.FailURL + "Auth" + installment + (r.OrderNumber + rndNum) + n.StoreKey
	hashData := utils.GenerateSHA1(hashFormat)

	paymentParams := map[interface{}]interface{}{
		"clientid":                        n.MerchantID, // mağaza no
		"storetype":                       n.SecurityType,
		"oid":                             r.OrderNumber,           // sipariş numarası
		"okUrl":                           r.SuccessURL,            // 3D doğrulama başarılı olduktan sonra adres
		"failUrl":                         r.FailURL,               // 3D doğrulama başarısız olduktan sonra yönlendirilecek adres
		"rnd":                             r.OrderNumber + rndNum,  // sipariş numarası ve yukarda oluşturulan random numaranın birleşimi
		"islemtipi":                       "Auth",                  // Satış
		"lang":                            r.LanguageCode,          // Dil kodu örn : TR, EN vs.
		"ccode":                           "",                      // ülke kodu
		"cardHolderName":                  r.CardHolderName,        // kart üzerinde yazan isim
		"userid":                          r.Customer.CustomerID,   // veri tabanında kayıtlı kullanıcı id
		"email":                           r.Customer.EmailAddress, // müşteriye ait e-posta adresi
		"hash":                            hashData,                // base64 tipinde hash değeri
		"pan":                             r.CardNumber,            // kart numarası
		"cv2":                             cardCVV,                 // kart güvenlik nuamrası
		"Ecom_Payment_Card_ExpDate_Year":  r.ExpireYear,            //  son kullanma tarihi (yıl)
		"Ecom_Payment_Card_ExpDate_Month": r.ExpireMonth,           // son kullanma tarihi (ay)
		"taksit":                          installment,             // taskit adeti
		"cardType":                        r.CardType,              // Kart tipi direkt olarak VISA veya MASTERCARD olarak gönderilmediğinden dolayı dönüştürme işlemi yapılmalı
		"currency":                        fmt.Sprintf("%v", r.CurrencyCode),
		"amount":                          fmt.Sprintf("%.2f", r.OrderTotal),
		"Fismi":                           r.CardHolderName,
		"Tismi":                           r.CardHolderName,
	}

	if r.Customer.BillingAddress != "" {
		if r.Customer.IsCompany {
			paymentParams["Faturafirma"] = r.Customer.FullName
		}
		paymentParams["Fadres"] = r.Customer.BillingAddress // fatura adresi
		paymentParams["Fpostakodu"] = r.Customer.BillingAddressZipCode
	}

	if r.Customer.ShippingAddress != "" {
		paymentParams["tadres"] = r.Customer.ShippingAddress // teslimat adresi
		paymentParams["tpostakodu"] = r.Customer.ShippingAddressZipCode
	}

	return models.PaymentGatewayResponse{
		IsSuccess:       true,
		Message:         "",
		HTMLFormContent: utils.PrepareForm(n.ApiURL, paymentParams),
	}, nil
}

func (n Netspay) VerifyPayment(r *models.VerifyPaymentRequest) (models.VerifyPaymentResponse, error) {

	var err error
	successStatusCodes := []string{"1", "2", "3", "4"}
	form := r.BankParams

	mdStatus := form.Get("mdStatus")
	procReturnCode := form.Get("ProcReturnCode")

	if mdStatus == "" {
		err = errors.New("mdstatus değeri alınamadı")
		return models.VerifyPaymentResponse{}, err
	}

	if procReturnCode != "00" || sort.SearchStrings(successStatusCodes, mdStatus) > 0 {
		return models.VerifyPaymentResponse{
			IsSuccess:      false,
			BankMessage:    form.Get("ErrMsg"),
			BankStatusCode: procReturnCode,
			BankErrorCode:  procReturnCode,
			TransactionID:  form.Get("TransId"),
			OrderID:        form.Get("oid"),
			PaidAmount:     0.0,
		}, nil
	}

	ok := verifyHash(form.Get("HASH"), r.BankParams, n.StoreKey)
	if ok == false {
		err = errors.New("Güvenlik imzası doğrulanamadı. Bankanızla iletişime geçiniz. Param name : HASH")
		return models.VerifyPaymentResponse{}, err
	}

	paidAmount, _ := strconv.ParseFloat(form.Get("amount"), 64)

	return models.VerifyPaymentResponse{
		IsSuccess:      true,
		BankMessage:    "",
		BankStatusCode: procReturnCode,
		TransactionID:  form.Get("TransId"),
		OrderID:        form.Get("oid"),
		PaidAmount:     paidAmount,
	}, nil

}

func convertInstallment(installment int) string {
	if installment <= 1 {
		return ""
	} else {
		return strconv.Itoa(installment)
	}
}

func verifyHash(source string, values url.Values, storekey string) bool {
	hashFormat := values.Get("clientid") + values.Get("oid") + values.Get("AuthCode") + values.Get("ProcReturnCode") + values.Get("Response") + values.Get("mdStatus") + values.Get("cavv") + values.Get("eci") + values.Get("md") + values.Get("rnd") + storekey
	hash := utils.GenerateSHA1(hashFormat)

	if source != hash {
		return false
	}

	return true
}
