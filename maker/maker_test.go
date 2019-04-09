package maker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mk "github.com/bukalapak/mkdash/maker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MakerSuite struct {
	suite.Suite
	maker        *mk.Maker
	makerTimeout *mk.Maker
}

func TestMakerSuite(t *testing.T) {
	suite.Run(t, &MakerSuite{})
}

func (m *MakerSuite) SetupSuite() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	tsTimeout := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2000 * time.Millisecond)
	}))
	m.maker = mk.New(ts.URL, "", 1*time.Second)
	m.makerTimeout = mk.New(tsTimeout.URL, "", 1*time.Millisecond)
}

func (m *MakerSuite) TestCreateDashboard() {
	err := m.maker.CreateDashboard("")
	assert.Nil(m.T(), err, "err should be nil")
	// test timeout
	err = m.makerTimeout.CreateDashboard("")
	assert.NotNil(m.T(), err, "err should be timeout")
}
