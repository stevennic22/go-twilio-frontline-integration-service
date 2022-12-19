package callbacks

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevennic22/TwFIS/providers"
)

// Build response for channel Proxy Address requests
func HandleGetProxyAddress(c *gin.Context, params map[string]string) {
	log.Println("Getting Proxy Address")

	channelName := params["ChannelType"]

	proxyAddress := getCustomerProxyAddress(channelName)

	if proxyAddress == "" {
		log.Println("Proxy address not found!")
		c.String(http.StatusForbidden, "Proxy address not found.")
	} else {
		log.Println("Got proxy address!")
		c.JSON(http.StatusOK, gin.H{
			"proxy_address": proxyAddress,
		})
	}
}

// Return proxy address for specific channel
func getCustomerProxyAddress(channelName string) string {
	if channelName == "whatsapp" {
		return (providers.CurrentEnv.RetrieveValue("TWILIO_WHATSAPP_NUMBER"))
	} else {
		return (providers.CurrentEnv.RetrieveValue("TWILIO_SMS_NUMBER"))
	}
}
