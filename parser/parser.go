// package parser implement parse template function for mkdash
package parser

import (
	"bytes"
	"text/template"

	"github.com/bukalapak/mkdash/resource"
	"github.com/pkg/errors"
)

// Parser is a struct that connect to mkdash
type Parser struct {
	tmpl *template.Template
}

// InputData is a struct that hold input data
type InputData struct {
	// Project is project name for the dashboard
	Project string
	// Service is a service name for the dashboard
	Service string
}

// New return a parser struct to pare data provided
func New() *Parser {
	return &Parser{
		tmpl: template.New(""),
	}
}

// ParseTemplate will parse the data to intended parse
func (p *Parser) ParseTemplate(data *InputData) (string, error) {
	t, err := p.tmpl.Parse(resource.DashboardTemplate)
	if err != nil {
		return "", errors.Wrap(err, "fail to parse template")
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		err = errors.Wrap(err, "fail to execute template: ")
	}
	return tpl.String(), err
}
