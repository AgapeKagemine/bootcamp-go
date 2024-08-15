package handler

import (
	"bytes"
	"consume-api/internal/domain"
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsumeSOAP(c *gin.Context) {
	client := &http.Client{}

	leftStr := c.Query("left")
	rightStr := c.Query("right")

	// left, err := strconv.Atoi(leftStr)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// right, err := strconv.Atoi(rightStr)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// soapRequest := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
	// 	<soap:Envelope
	// 		xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	// 		xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	// 		xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	// 		<soap:Body>
	// 			<Multiply
	// 				xmlns="http://tempuri.org/">
	//   				<intA>%d</intA>
	//   				<intB>%d</intB>
	// 			</Multiply>
	// 		</soap:Body>
	// 	</soap:Envelope>`, left, right,
	// )

	soapRequest := &domain.EnvelopeRequest{
		Xsi:  "http://www.w3.org/2001/XMLSchema-instance",
		Xsd:  "http://www.w3.org/2001/XMLSchema",
		Soap: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: domain.EnvelopRequestBody{
			Multiply: domain.EnvelopRequestMultiply{
				Xmlns: "http://tempuri.org/",
				IntA:  leftStr,
				IntB:  rightStr,
			},
		},
	}

	xmlRequest, err := xml.Marshal(&soapRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	req, err := http.NewRequest("POST", "http://www.dneonline.com/calculator.asmx", bytes.NewBuffer(xmlRequest))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	req.Header.Set("Content-Type", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	var envelop domain.EnvelopeResponse
	err = xml.NewDecoder(resp.Body).Decode(&envelop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, envelop)
}

// payload, err := xml.Marshal(&soapRequest)
// if err != nil {
// 	c.JSON(http.StatusInternalServerError, err)
// 	return
// }
