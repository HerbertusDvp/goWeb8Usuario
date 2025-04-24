package ruta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goweb1/modelos"
	"goweb1/pkg/utils"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/joho/godotenv"
)

func PasarelaHomePay(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/PagoHome.html", utils.Frontend))

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}

func PasarelaWebPay(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/PagoWebPay.html", utils.Frontend))

	// Comunicacion a webpay
	errorVariables := godotenv.Load("internal/config/.env")
	if errorVariables != nil {
		fmt.Println(errorVariables)
		return
	}

	client := &http.Client{}

	datos := map[string]string{
		"buy_order":  "ordenCompra12345678",
		"session_id": "sesion1234557545",
		"amount":     "10000",
		"return_url": "http://localhost:8080/pasarela/webpay/respuesta",
	}

	jsonValue, _ := json.Marshal(datos)

	req, err := http.NewRequest("POST", os.Getenv("WEBPAY_URL"), bytes.NewBuffer(jsonValue))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Tbk-Api-Key-Id", os.Getenv("WEBPAY_ID"))
	req.Header.Set("Tbk-Api-Key-Secret", os.Getenv("WEBPAY_SECRET"))

	if err != nil {
		fmt.Println(err)
	}

	reg, err2 := client.Do(req)
	defer reg.Body.Close()

	if err2 != nil {
		return
	}

	body, err := io.ReadAll(reg.Body)
	webpay := modelos.WebPayModel{}
	errJson := json.Unmarshal(body, &webpay)

	if errJson != nil {
		return
	}

	//retorno
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
		"url":     webpay.Url,
		"token":   webpay.Token,
	}

	template.Execute(response, data)
}

func WebPayRespuesta(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/PagoWebPayRespuesta.html", utils.Frontend))

	errorVariables := godotenv.Load("internal/config/.env")

	if errorVariables != nil {
		fmt.Println(errorVariables)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", os.Getenv("WEBPAY_URL")+"/"+request.URL.Query().Get("token_ws"), nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Tbk-Api-Key-Id", os.Getenv("WEBPAY_ID"))
	req.Header.Set("Tbk-Api-Key-Secret", os.Getenv("WEBPAY_SECRET"))

	if err != nil {
		fmt.Println(err)
	}

	reg, err2 := client.Do(req)

	if err2 != nil {
		fmt.Println()
	}

	defer reg.Body.Close()

	body, err := io.ReadAll(reg.Body)
	webpay := modelos.WebpayRespuestaModel{}

	errJson := json.Unmarshal(body, &webpay)

	if errJson != nil {
		fmt.Println(err)
	}

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)
	amount := strconv.Itoa(webpay.Amount)

	data := map[string]string{

		"css":                 cssSesion,
		"mensaje":             cssMensaje,
		"token_ws":            request.URL.Query().Get("token_ws"),
		"vci":                 webpay.Vci,
		"amount":              amount,
		"status":              webpay.Status,
		"buy_order":           webpay.Buy_order,
		"session_id":          webpay.Session_id,
		"card_detail":         webpay.Card_detail["card_number"],
		"accounting_date":     webpay.Accounting_date,
		"transaction_date":    webpay.Transaction_date,
		"authorization_code":  webpay.Authorization_code,
		"payment_type_code":   webpay.Payment_type_code,
		"response_code":       webpay.Response_code,
		"installments_number": webpay.Installments_number,
	}

	template.Execute(response, data)
}

func PasarelaPayPal(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/PagoPayPal.html", utils.Frontend))

	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}
