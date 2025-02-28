package ruta

import (
	"fmt"
	"goweb1/pkg/utils"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/jung-kurt/gofpdf"
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

func RecursosExcel(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("web/templates/recursosExcel.html", utils.Frontend))
	cssSesion, cssMensaje := utils.RetornaMensaje(response, request)

	data := map[string]string{
		"css":     cssSesion,
		"mensaje": cssMensaje,
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
