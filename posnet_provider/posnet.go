package posnet_provider

import (
	"fmt"
	"github.com/mehmetcantas/gopos/internal/utils"
	"github.com/mehmetcantas/gopos/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Posnet struct {
	BankName     string
	TerminalID   string
	MerchantID   string
	PosnetID     string
	UseSandbox   bool
	ApiURL       string
	SecurityType string // 3D, 3D_PAY, 3D_PAY_HOSTING vb.
}

func (p Posnet) PreparePaymentGatewayForm(r *models.PaymentGatewayRequest) (*models.PaymentGatewayResponse, error) {

	var err error

	var orderTotal = convertOrderTotal(r.OrderTotal)
	var xmlReq string
	xmlReq = `<?xml version="1.0" encoding="utf-8"?>
				<posnetRequest>
					<mid>` + p.MerchantID + `</mid>
					<tid>` + p.TerminalID + `</tid>
					<oosRequestData>
						<posnetid>` + p.PosnetID + `</posnetid>
						<XID>` + r.OrderNumber + `</XID>
						<amount>` + orderTotal + `</amount>
						<currencyCode>` + utils.ConvertCurrencyCode(r.CurrencyCode) + `</currencyCode>
						<installment>00</installment>
						<tranType>Sale</tranType>
						<cardHolderName>` + r.CardHolderName + `</cardHolderName>
						<ccno>` + r.CardNumber + `</ccno>
						<expDate>` + r.ExpireMonth + r.ExpireYear + `</expDate>
						<cvc>` + r.CVV + `</cvc>
					</oosRequestData>
				</posnetRequest>`

	client := &http.Client{}
	data := url.Values{}
	data.Set("xmldata", xmlReq)
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", p.ApiURL, strings.NewReader(data.Encode()))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//TODO: check errors
	return &models.PaymentGatewayResponse{
		IsSuccess:       true,
		Message:         resp.Status,
		HTMLFormContent: string(body),
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
