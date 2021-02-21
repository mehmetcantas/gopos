package garanti_provider

import (
	// std packages

	"errors"
	"fmt"
	"github.com/mehmetcantas/gopos/internal/utils"
	"github.com/mehmetcantas/gopos/models"
	"strconv"
	"strings"
)

type Garanti struct {
	StoreKey           string
	MerchantID         string
	TerminalID         string
	TerminalUserID     string
	TerminalProvUserID string
	ProvisionPassword  string
	TransactionType    string
	UseSandbox         bool
	ApiURL             string
	SecurityType       string // 3D, 3D_PAY, 3D_PAY_HOSTING vb.
}

func (g Garanti) PreparePaymentGatewayForm(r *models.PaymentGatewayRequest) (*models.PaymentGatewayResponse, error) {

	var mode string = "PROD"
	if g.UseSandbox {
		mode = "TEST"
	}

	var orderTotal = convertOrderTotal(r.OrderTotal)
	var secureData = utils.GenerateSHA1(g.ProvisionPassword + ("0" + g.TerminalID))
	var installment = convertInstallmentCount(r.InstallmentCount)
	var hashData = utils.GenerateSHA1(g.TerminalID + r.OrderNumber + orderTotal + r.SuccessURL + r.FailURL + g.TransactionType + installment + g.StoreKey + strings.ToUpper(secureData))

	var paymentCollection = map[interface{}]interface{}{
		"mode":                  mode,
		"apiversion":            "v0.01",
		"terminalprovuserid":    g.TerminalProvUserID,
		"terminaluserid":        g.TerminalUserID,
		"terminalmerchantid":    g.MerchantID,
		"txntype":               g.TransactionType,                         // direkt satış için "sales"
		"txnamount":             orderTotal,                                // sipariş tutarı. Burada . ya da , gibi işaretler kaldırılmalı. Örn: 154,45 => 15445. Son iki karakter kuruş olarak algılanır.
		"txncurrencycode":       utils.ConvertCurrencyCode(r.CurrencyCode), // TL => 949, USD => 940, EURO => 978, GBP => 826, JPY => 392
		"txninstallmentcount":   installment,                               // peşin satış için boş gönderilmesi gerekiyor.
		"orderid":               r.OrderNumber,
		"terminalid":            g.TerminalID,   // mağaza üye numarası
		"successurl":            r.SuccessURL,   // 3D doğrulaması başarılı olması durumunda yönlendirilecek sayfa linki
		"errorurl":              r.FailURL,      // 3D doğrulaması başarısız olması durumunda yönlendirilecek sayfa linki
		"secure3dsecuritylevel": g.SecurityType, // 3D, 3D_PAY vb. güvenlik seçenekleri
		"customeripaddress":     r.CustomerIPAddress,
		"secure3dhash":          strings.ToUpper(hashData),
		"customeremailaddress":  r.CustomerEmailAddress,
		"companyname":           r.CompanyName,
		"cardnumber":            r.CardNumber,
		"cardexpiredatemonth":   r.ExpireMonth,
		"cardexpiredateyear":    r.ExpireYear,
		"cardcvv2":              r.CVV,
	}
	//TODO: check errors
	return &models.PaymentGatewayResponse{
		IsSuccess:       true,
		Message:         "",
		HTMLFormContent: utils.PrepareForm(g.ApiURL, paymentCollection),
	}, nil
}

func (g Garanti) VerifyPayment(r *models.VerifyPaymentRequest) (*models.VerifyPaymentResponse, error) {
	form := r.BankParams
	var err error

	if form == nil {
		err = errors.New("Bankadan dönen form bilgileri boş olamaz")
		return nil, err
	}

	mdstatus := form.Get("mdstatus")

	if mdstatus != "1" {
		err = errors.New(form.Get("errmsg"))
		return &models.VerifyPaymentResponse{
			IsSuccess:      false,
			BankMessage:    form.Get("hostmsg"),
			BankStatusCode: form.Get("mdstatus"),
			TransactionID:  form.Get("transid"),
			OrderID:        form.Get("orderid"),
			PaidAmount:     0.0,
		}, err
	}

	paidAmount, _ := strconv.Atoi(form.Get("txnamount"))
	return &models.VerifyPaymentResponse{
		IsSuccess:      true,
		BankMessage:    "",
		BankStatusCode: form.Get("mdstatus"),
		TransactionID:  form.Get("transid"),
		OrderID:        form.Get("orderid"),
		PaidAmount:     float64((paidAmount / 100)),
	}, nil

}

func convertInstallmentCount(icount int) string {
	if icount == 0 || icount == 1 {
		return ""
	}

	return strconv.Itoa(icount)
}

func convertOrderTotal(amount float64) string {

	s := fmt.Sprintf("%.2f", amount)
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, ",", "", -1)

	return s
}
