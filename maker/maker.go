// package maker implement a request sender to create dashboard
package maker

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// Maker is a struct that connect to mkdash
type Maker struct {
	grafanaAuth string
	grafanaURL  string
	timeOut     time.Duration
}

// New returns a maker struct for creating dashboard
func New(url, auth string, timeout time.Duration) *Maker {
	return &Maker{
		grafanaAuth: auth,
		grafanaURL:  url,
		timeOut:     timeout,
	}
}

// CreateDashboard send a request to create dashboard on grafana
func (m *Maker) CreateDashboard(body string) error {

	auth := "Bearer " + m.grafanaAuth
	url := m.grafanaURL + "/api/dashboards/db"

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	client := &http.Client{
		Timeout: m.timeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	} else {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return errors.New("fail to create dashboard: " + string(respBody))
	}
}
