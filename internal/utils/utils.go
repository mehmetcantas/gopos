package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

func PrepareForm(bu string, collection map[interface{}]interface{}) string {

	var sb strings.Builder
	var formID = "CreditCardPaymentForm"

	sb.WriteString(fmt.Sprintf("<form id=%[1]q name=%[1]q action=%[2]q method=\"POST\">", formID, bu))

	for item := range collection {
		sb.WriteString(fmt.Sprintf("<input type=\"hidden\" name=%q value=%q />", item, collection[item]))
	}

	sb.WriteString(fmt.Sprintf("</form><script> var v%[1]s = document.%[1]s; v%[1]s.submit(); </script>", formID))

	return sb.String()
}

func GenerateSHA1(hashString string) string {
	h := sha1.New()
	io.WriteString(h, hashString)
	hb := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(hb)
}

func ConvertCurrencyCode(ccode string) string {
	switch ccode {
	case "TL":
		return "949"
	case "EUR":
		return "978"
	case "USD":
		return "840"
	case "GBP":
		return "826"
	case "JPY":
		return "392"
	default:
		return "949"
	}
}
