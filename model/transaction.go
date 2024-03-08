package model

import (
	"bytes"
	"text/template"
	"time"
)

type Transaction struct {
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Accounts []Account `json:"accounts"`
}

const txTemplate = `{{.Date.Format "2006-01-02" }} {{.Name}}
{{range .Accounts}}  {{.Name}}{{ if ne .Amount 0.0 }}  {{.Amount}} {{.Commodity}}{{end}}
{{end}}`

func (tx Transaction) Generate() (string, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("transaction").Parse(txTemplate)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(buf, tx)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
