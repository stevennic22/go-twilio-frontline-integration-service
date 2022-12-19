package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevennic22/TwFIS/callbacks"
	"github.com/twilio/twilio-go"
)

// Return error for unknown callback specified
func handleUnknown(c *gin.Context, description string, unknownItem string) {
	log.Printf("Unknown Callback: %s | %s", description, unknownItem)
	c.String(
		http.StatusUnprocessableEntity,
		fmt.Sprintf("Unknown %s: %s", description, unknownItem))
}

func ConversationsCallbackHandler(c *gin.Context, client *twilio.RestClient, params map[string]string) {
	log.Printf("Conversations Callback: %s", params["EventType"])
	switch params["EventType"] {

	case "onConversationAdd":
		callbacks.UpdateConversation(c, params)

	case "onParticipantAdded":
		callbacks.UpdateParticipantAttributes(c, client, params)

	default:
		handleUnknown(c, "EventType", params["EventType"])
	}

}

func CRMCallbackHandler(c *gin.Context, params map[string]string) {
	log.Printf("CRM Callback: %s", params["Location"])
	switch params["Location"] {

	case "GetCustomersList":
		callbacks.HandleGetCustomersListCallback(c, params)

	case "GetCustomerDetailsByCustomerId":
		callbacks.HandleGetCustomerDetailsByCustomerIdCallback(c, params)

	default:
		handleUnknown(c, "Location", params["Location"])
	}
}

func OutgoingConversationCallbackHandler(c *gin.Context, params map[string]string) {
	log.Printf("Outgoing Conversation Callback: %s", params["Location"])
	switch params["Location"] {

	case "GetProxyAddress":
		callbacks.HandleGetProxyAddress(c, params)

	default:
		handleUnknown(c, "Location", params["Location"])
	}
}

func RoutingCallbackHandler(c *gin.Context, client *twilio.RestClient, params map[string]string) {
	log.Printf("Routing Callback: %s", params["ConversationSid"])
	callbacks.RouteConversation(client, params["ConversationSid"], params["MessagingBinding.Address"])
}

func TemplatesCallbackHandler(c *gin.Context, params map[string]string) {
	log.Printf("Templates Callback: %s", params["Location"])
	switch params["Location"] {

	case "GetTemplatesByCustomerId":
		callbacks.HandleGetTemplatesByCustomerIdCallback(c, params)

	default:
		handleUnknown(c, "Location", params["Location"])
	}
}
