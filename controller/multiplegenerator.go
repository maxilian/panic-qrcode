package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"panic-qrcode/model"

	"github.com/gin-gonic/gin"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateMultipleQR(c *gin.Context) {

	var input []model.MultipleDetail

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error pak",
			"error":   true,
		})
		return
	}

	//get data
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(5, 10, 5)
	//m.SetFirstPageNb(1)

	for _, d := range input {
		//fmt.Println(d.QrString)
		Body_MultipleQR(m, d)
	}

	file, err := m.Output()

	if err != nil {
		return
	}
	var reader io.Reader
	nRead := int64(len(file.Bytes()))
	reader = bytes.NewReader(file.Bytes())
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="qrcode_multiple.pdf"`,
	}

	c.DataFromReader(http.StatusOK, nRead, "application/pdf", reader, extraHeaders)
}

func Body_MultipleQR(m pdf.Maroto, detail model.MultipleDetail) {

	m.SetBorder(false)
	//m.AddPage()
	var png []byte
	png, err := qrcode.Encode(detail.QrString, qrcode.Medium, 512)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	var encodedString = base64.StdEncoding.EncodeToString(png)
	m.Row(200, func() {
		m.Col(12, func() {
			_ = m.Base64Image(encodedString, consts.Png, props.Rect{
				//Left:    5,
				//Top:     5,
				Center:  true,
				Percent: 100,
			})
		})
	})

	m.Row(30, func() {
		m.ColSpace(1)
		m.Col(10, func() {
			m.Text(detail.Nomor, props.Text{
				Top: 5,
				//Left:  3,
				Size:  50.0,
				Align: consts.Center,
				Style: consts.Bold,
			})
		})
		m.ColSpace(1)
	})
	m.Row(20, func() {
		m.ColSpace(1)
		m.Col(10, func() {
			m.Text(detail.Detail, props.Text{
				Top: 5,
				//Left:  3,
				Size:  25.0,
				Align: consts.Center,
				//Style: consts.Bold,
			})
		})
		m.ColSpace(1)
	})

	m.Row(10, func() {
		m.ColSpace(12)
	})
}
