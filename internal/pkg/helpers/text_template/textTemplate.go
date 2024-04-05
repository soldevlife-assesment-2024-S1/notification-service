package texttemplate

import (
	"bytes"
	"html/template"
)

var (
	TemplateHtmlMessge = `
		<!DOCTYPE html>
		<html lang="en">
		<head>	
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h1>{{.Message}}</h1>
		</body>
		</html>
	`
	TemplateHtmlInvoice = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h1>Invoice</h1>
			<p>Booking ID: {{.BookingID}}</p>
			<p>Payment Amount: {{.PaymentAmount}}</p>
			<p>Payment Expiration: {{.PaymentExpiration}}</p>
		</body>
		</html>
	`
	TemplateHtmlPayment = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h1>Payment Receipt</h1>
			<p>Booking ID: {{.BookingID}}</p>
			<p>Message: {{.Message}}</p>
			<p>Payment Method: {{.PaymentMethod}}</p>
		</body>
		</html>
	`
)

func PopulateTemplate(templateHtml string, data interface{}) (string, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(templateHtml))
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
