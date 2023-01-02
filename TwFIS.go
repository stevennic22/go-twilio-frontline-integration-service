package main

import (
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"

	"github.com/stevennic22/TwFIS/providers"
	"github.com/stevennic22/TwFIS/web"
)

// Generation of Twilio REST Client
func genClient(accountSid string, authToken string) *twilio.RestClient {
	return twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

// Check for valid signature and
// route to individual callback handler functions
func callbackRouting(c *gin.Context, client *twilio.RestClient, authToken string) {

	err := c.Request.ParseForm()
	web.CheckError(err, true)

	collapsedFormVals := web.CollapseURLParams(c.Request.Form)

	var isAValidRequest = false
	if c.GetHeader("x-twilio-signature") != "" {
		urlToTest := "https://" + c.Request.Host + c.Request.URL.Path
		isAValidRequest = web.SignatureValidation(authToken, urlToTest, collapsedFormVals, c.GetHeader("x-twilio-signature"))
	}

	if !isAValidRequest {
		log.Println("Not a valid request")
		web.Return403(c)
	} else {
		switch c.Param("callback") {
		case "conversations":
			ConversationsCallbackHandler(c, client, collapsedFormVals)

		case "routing":
			RoutingCallbackHandler(c, client, collapsedFormVals)

		case "outgoing-conversation":
			OutgoingConversationCallbackHandler(c, collapsedFormVals)

		case "crm":
			CRMCallbackHandler(c, collapsedFormVals)

		case "templates":
			TemplatesCallbackHandler(c, collapsedFormVals)

		default:
			handleUnknown(c, "Callback", c.Param("callback"))
		}

	}
}

// Initialize webserver
func main() {
	providers.CurrentEnv.Location = path.Join("static/" + ".env")
	serverPort := ":" + providers.CurrentEnv.RetrieveValue("LISTEN_PORT")

	providers.CustomersToWorkersMap["+15557778888"] = providers.CurrentEnv.RetrieveValue("Worker_Identity")
	providers.SetUpCustomerList(providers.CurrentEnv.RetrieveValue("Worker_Identity"))

	accountSid := providers.CurrentEnv.RetrieveValue("TWILIO_ACCOUNT_SID")
	authToken := providers.CurrentEnv.RetrieveValue("TWILIO_AUTH_TOKEN")

	log.Printf("TWILIO_ACCOUNT_SID: %s\n", accountSid)

	client := genClient(accountSid, authToken)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/callbacks/:callback", func(c *gin.Context) {
		callbackRouting(c, client, authToken)
	})

	r.NoRoute(func(c *gin.Context) {
		log.Printf("%s%s%s", "https://", c.Request.Host, c.Request.URL)
		log.Println(c.Request.URL.Path)
		log.Println(c.Request.Method)
		c.String(http.StatusNotFound, "404-FILENOTFOUND")
	})

	r.NoMethod(func(c *gin.Context) {
		log.Printf("%s%s%s", "https://", c.Request.Host, c.Request.URL)
		log.Println(c.Request.URL.Path)
		c.String(http.StatusMethodNotAllowed, "405-METHODNOTALLOWED")
	})

	log.Printf("Starting server at %s\n", serverPort[1:])
	r.Run(serverPort) // listen and serve on 0.0.0.0:serverPort
}
