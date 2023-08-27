package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQR(c *gin.Context) {
	nomor := c.Query("nomor")
	nama := c.Query("nama")

	header := map[string]string{
		"nama":  strings.ToUpper(nama),
		"nomor": strings.ToUpper(nomor),
	}

	//get data
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)
	//m.SetFirstPageNb(1)

	Body_QR(m, header)

	file, err := m.Output()

	if err != nil {
		return
	}
	var reader io.Reader
	nRead := int64(len(file.Bytes()))
	reader = bytes.NewReader(file.Bytes())
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="qrcode_` + nomor + `.pdf"`,
	}

	c.DataFromReader(http.StatusOK, nRead, "application/pdf", reader, extraHeaders)
}

func Body_QR(m pdf.Maroto, header map[string]string) {

	m.SetBorder(false)
	// m.Row(100, func() {
	// 	m.ColSpace(1)
	// 	m.Col(10, func() {
	// 		m.QrCode(header["nomor"], props.Rect{
	// 			Left:    5,
	// 			Top:     5,
	// 			Center:  true,
	// 			Percent: 100,
	// 		})
	// 	})
	// 	m.ColSpace(1)
	// })

	var png []byte
	png, err := qrcode.Encode(header["nomor"], qrcode.Medium, 256)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	var encodedString = base64.StdEncoding.EncodeToString(png)
	m.Row(180, func() {
		m.Col(12, func() {
			_ = m.Base64Image(encodedString, consts.Png, props.Rect{
				//Left:    5,
				//Top:     5,
				Center:  true,
				Percent: 120,
			})
		})
	})

	m.Row(30, func() {
		m.ColSpace(1)
		m.Col(10, func() {
			m.Text(header["nomor"], props.Text{
				Top:   5,
				Left:  3,
				Size:  50.0,
				Align: consts.Center,
				Style: consts.Bold,
			})
		})
		m.ColSpace(1)
	})
	m.Row(20, func() {
		//m.ColSpace(1)
		m.Col(12, func() {
			m.Text(header["nama"], props.Text{
				Top:   5,
				Left:  3,
				Size:  30.0,
				Align: consts.Center,
				//Style: consts.Bold,
			})
		})
		//m.ColSpace(1)
	})

	m.Row(10, func() {
		m.ColSpace(12)
	})
}
