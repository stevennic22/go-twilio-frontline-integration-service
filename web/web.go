package web

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	validationClient "github.com/twilio/twilio-go/client"
)

// Check if error exists and stop application if necessary
func CheckError(err error, hardStop bool) bool {
	if err != nil {
		if hardStop {
			log.Fatalf("Fatal error within application: %s", err.Error())
		}
		log.Println(err.Error())
		return (true)
	}

	return (false)
}

// Return forbidden when validation fails
func Return403(c *gin.Context) {
	log.Printf("%s%s%s", "https://", c.Request.Host, c.Request.URL)
	log.Println(c.Request.URL.Path)
	c.JSON(http.StatusForbidden, gin.H{
		"error": "403-FORBIDDEN",
	})
}

// Return whether request has valid signature
func SignatureValidation(passed_Auth_Token string, full_url_path string, params map[string]string, received_signature string) bool {
	validator := validationClient.NewRequestValidator(passed_Auth_Token)
	return validator.Validate(full_url_path, params, received_signature)
}

// Collapse URL Values/Params into a flat key/value map
func CollapseURLParams(params url.Values) map[string]string {
	formvals := make(map[string]string)

	for k, v := range params {
		if len(v) < 2 {
			formvals[k] = v[0]
		}
	}

	return (formvals)
}
