package vcodeapi

import (
	"errors"
	"io"
	"net/http"
	"os"

	// use HMAC fork by securityRelicATgithub
	"github.com/securityRelic/vcodeHMAC"
)

func detailedReportPdf(credsFile, buildID string) error {

	filename := "veracode-report-" + buildID + ".pdf"

	// Create HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://analysiscenter.veracode.com/api/4.0/detailedreportpdf.do?build_id="+buildID, nil)
	if err != nil {
		return err
	}

	// Set authorization header
	authHeader, err := vcodeHMAC.GenerateAuthHeader(credsFile, req.Method, req.URL.String())
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", authHeader)

	// Make HTTP Request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.Status != "200 OK" {
		return errors.New("detailedreport.do call error: " + resp.Status)
	}
	defer resp.Body.Close()
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return nil
}
