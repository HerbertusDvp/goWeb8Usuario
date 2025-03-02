package ruta

import (
	"encoding/base64"
	"fmt"
	"goweb1/modelos"
	"goweb1/pkg/utils"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/jung-kurt/gofpdf"
	qrcode "github.com/skip2/go-qrcode"
	excelize "github.com/xuri/excelize/v2"
)

func Recursos(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/recursos.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}

func RecursosEmail(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/recursoEmail.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	utils.EnviarCorreo()

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}

func RecursosQR(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/recursoQR.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	// Generación de codigo QR
	imagenQR, err := qrcode.Encode("https://estructuradedatos.com/", qrcode.High, 256)

	if err != nil {
		log.Fatal("Error al enerar codifo QR", err)
	}
	imagen := base64.StdEncoding.EncodeToString(imagenQR)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
		"imagen":  imagen,
	}
	template.Execute(response, data)

}

func RecursosExcel(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/recursosExcel.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)
	// Excel
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetCellValue("Sheet1", "A1", "id")
	f.SetCellValue("Sheet1", "B1", "Nombre")
	f.SetCellValue("Sheet1", "C1", "Correo")
	f.SetActiveSheet(index)

	// Anexion de datos en el excel
	cliente := modelos.Clientes{
		modelos.Cliente{1, "Cesar Cancino", "info@gmail.com"},
		modelos.Cliente{2, "Juan Perez", "jaun@gmail.com"},
	}
	contador := 2
	i := 0

	for _, service := range cliente {
		fila := strconv.Itoa(contador)

		f.SetCellValue("Sheet1", "A"+fila, service.Id)
		f.SetCellValue("Sheet1", "B"+fila, service.Nombre)
		f.SetCellValue("Sheet1", "C"+fila, service.Correo)
		contador++
		i++
	}

	//COnstruccion del documento
	time := strings.Split(time.Now().String(), " ")
	nombre := string(time[4][6:14]) + ".xlsx"
	if err := f.SaveAs("web/static/excel/" + nombre); err != nil {
		fmt.Println(err)
	}

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
		"nombre":  nombre,
	}
	template.Execute(response, data)
}

func RecursosGeneraExcel(response http.ResponseWriter, request *http.Request) {

}

// ----------------------PDF --------------------------
func RecursosPdf(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/recursoPDF.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
	}

	template.Execute(response, data)
}

func RecursosGeneraPDF(response http.ResponseWriter, request *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("Hello.pdf")

	if err != nil {
		utils.CrearMensaje(response, request, "danger", "Error al crear el pdf")
		http.Redirect(response, request, "/rescursos/pdf", http.StatusSeeOther)
		return
	}

	utils.CrearMensaje(response, request, "success", "PDF creado")
	http.Redirect(response, request, "/recursos/pdf", http.StatusSeeOther)

}

// Funciones para pdfs

func ImageFile(fileStr string) string {
	return filepath.Join(gofpdfDir, "web/static/images", fileStr)
}

var gofpdfDir string

func Filename(baseStr string) string {
	return PdfFile(baseStr + ".pdf")
}
func PdfFile(fileStr string) string {
	return filepath.Join(PdfDir(), fileStr)
}
func PdfDir() string {
	return filepath.Join(gofpdfDir, "web/static/pdf")
}
func Summary(err error, fileStr string) {
	if err == nil {
		fileStr = filepath.ToSlash(fileStr)
		fmt.Printf("Successfully generated %s\n", fileStr)
	} else {
		fmt.Println(err)
	}
}

func RecursosGeneraPDF2(response http.ResponseWriter, request *http.Request) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	// First page: manual local link
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 20)
	_, lineHt := pdf.GetFontSize()
	pdf.Write(lineHt, "To find out what's new in this tutorial, click ")
	pdf.SetFont("", "U", 0)
	link := pdf.AddLink()
	pdf.WriteLinkID(lineHt, "here", link)
	pdf.SetFont("", "", 0)
	// Second page: image link and basic HTML with link
	pdf.AddPage()
	pdf.SetLink(link, 0, -1)
	pdf.Image(ImageFile("logoStructure.jpeg"), 10, 12, 30, 0, false, "", 0, "http://www.fpdf.org")
	pdf.SetLeftMargin(45)
	pdf.SetFontSize(14)
	_, lineHt = pdf.GetFontSize()
	htmlStr := `You can now easily print text mixing different styles: <b>bold</b>, ` +
		`<i>italic</i>, <u>underlined</u>, or <b><i><u>all at once</u></i></b>!<br><br>` +
		`<center>You can also center text.</center>` +
		`<right>Or align it to the right.</right>` +
		`You can also insert links on text, such as ` +
		`<a href="http://www.fpdf.org">www.fpdf.org</a>, or on an image: click on the logo.`
	html := pdf.HTMLBasicNew()
	html.Write(lineHt, htmlStr)
	time := strings.Split(time.Now().String(), " ")

	nombre := string(time[4][6:14])
	fileStr := Filename(nombre)
	err := pdf.OutputFileAndClose(fileStr)
	Summary(err, fileStr)

	mensaje := "Se creó el documento PDF " + nombre + ".pdf de forma correcta"

	utils.CrearMensaje(response, request, "success", mensaje)
	http.Redirect(response, request, "/recursos/pdf", http.StatusSeeOther)

}
