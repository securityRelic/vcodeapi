Forked by securityRelic

> securityRelic@github: Added a quick hack back in 2019 to use Veracode sandboxes and retrieve detailed flaw report (pdf) in a targeted automated pipeline. Found brian1917/vcodeapi codebase easy to work with and functionality provided solid, but utimately moved in another direction. Keeping this for historical purposes.

---

# Veracode API Package
[![GoDoc](https://godoc.org/github.com/brian1917/vcodeapi?status.svg)](https://godoc.org/github.com/brian1917/vcodeapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/brian1917/vcodeapi)](https://goreportcard.com/report/github.com/brian1917/vcodeapi)

## Package Documentation
See here: https://godoc.org/github.com/brian1917/vcodeapi

## Description
Go package that provides easy access to the Veracode APIs. Each API typically has two files: one for making the http request and one for parsing the response.
For example, `detailedreport.go` calls the Veracode API and returns a `[byte]` and `detailedreportparser.go` parses the
XML response and returns usable objects such as flaws.

## Credentials File
Must be structured like the following:
```
veracode_api_key_id = ID HERE
veracode_api_key_secret = SECRET HERE
```

## Included APIs
1. Get App List (`/api/5.0/getapplist.do`)
2. Get Build List (`/api/5.0/getbuildlist.do`)
3. Get Sandbox List (`/api/5.0/getsandboxlist.do`)
4. Get Detailed Report (`/api/5.0/detailedreport.do`)
5. Get Team Info (`api/3.0/getteaminfo.do`)
6. Updated Mitigation Info (`api/updatemitigationinfo.do`)
7. Upload File (`api/5.0/uploadfile.do`)
8. Begin Prescan (`api/5.0/beginprescan.do`)

