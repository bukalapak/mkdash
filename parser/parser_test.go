package parser_test

import (
	"io/ioutil"
	"testing"

	pr "github.com/bukalapak/mkdash/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParserSuite struct {
	suite.Suite
	parser *pr.Parser
}

func TestParserSuite(t *testing.T) {
	suite.Run(t, &ParserSuite{})
}

func (p *ParserSuite) SetupSuite() {
	p.parser = pr.New()
}

func (p *ParserSuite) TestParseTemplate() {
	res, err := p.parser.ParseTemplate(&pr.InputData{
		Project: "chronos",
		Service: "api",
	})
	assert.Nil(p.T(), err, "err should be nil")

	exp, err := ioutil.ReadFile("../test/template.json")
	assert.Nil(p.T(), err, "no test parse expectation file")
	assert.Equal(p.T(), res, string(exp), "should be the same")
}
