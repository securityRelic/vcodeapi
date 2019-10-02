package vcodeapi

// ParseDetailedReportPdf retrieves detailed PDF report for supplied build ID
// sR - need to re-work this into one function - no need for two (detailedReportPdf)
// quick hack to get this working, while following orginal flow/structure of code.
func ParseDetailedReportPdf(credsFile, buildID string) error {

	err := detailedReportPdf(credsFile, buildID)
	if err != nil {
		return err
	}

	return nil

}
