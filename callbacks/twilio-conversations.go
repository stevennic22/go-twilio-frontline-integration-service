package callbacks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/stevennic22/TwFIS/providers"
	"github.com/stevennic22/TwFIS/web"

	"github.com/twilio/twilio-go"
	conversations "github.com/twilio/twilio-go/rest/conversations/v1"
)

type OnConversationAdded struct {
	Avatar string `json:"avatar"`
}

func (o OnConversationAdded) String() string {
	result, err := json.Marshal(o)

	if !web.CheckError(err, false) {
		return (string(result))
	} else {
		return ("{}")
	}
}

type WebhookOnConversationAdd struct {
	FriendlyName string `json:"friendly_name"`
	Attributes   string `json:"attributes"`
}

type WebhookOnParticipantAdded struct {
	Attributes struct {
		CustomerID  string `json:"customer_id"`
		Avatar      string `json:"avatar"`
		DisplayName string `json:"display_name"`
	} `json:"attributes"`
}

// Update Conversation with Customer-driven attributes/naming
func UpdateConversation(c *gin.Context, params map[string]string) {
	customerNumber := params["MessagingBinding.Address"]
	var isIncomingConversation bool

	if customerNumber == "" {
		isIncomingConversation = false
	} else {
		isIncomingConversation = true
	}

	if isIncomingConversation {
		var conversationFriendlyName string
		customerDetails := providers.GetCustomerByNumber(customerNumber)

		if customerDetails.DisplayName == "" {
			conversationFriendlyName = customerNumber
		} else {
			conversationFriendlyName = customerDetails.DisplayName
		}

		avatar := &OnConversationAdded{
			Avatar: customerDetails.Avatar,
		}

		var webhookDetails = &WebhookOnConversationAdd{
			FriendlyName: conversationFriendlyName,
			Attributes:   avatar.String(),
		}

		c.JSON(http.StatusOK, webhookDetails)
	} else {
		c.Status(http.StatusOK)
	}
}

// Update Participant Attributes (avatar/etc)
func UpdateParticipantAttributes(c *gin.Context, client *twilio.RestClient, incomingParams map[string]string) {
	var isCustomer bool
	var customerDetails *providers.Customer

	if incomingParams["MessagingBinding.Address"] != "" && incomingParams["Identity"] == "" {
		isCustomer = true
	} else {
		isCustomer = false
	}

	if isCustomer {
		customerDetails = providers.GetCustomerByNumber(incomingParams["MessagingBinding.Address"])

		if customerDetails.CustomerID > 0 {
			conversationParams := &conversations.UpdateConversationParticipantParams{}
			conversationParams.SetAttributes(fmt.Sprintf("{\"avatar\": \"%s\", \"customer_id\": \"%d\", \"display_name\": \"%s\"}",
				customerDetails.Avatar,
				customerDetails.CustomerID,
				customerDetails.DisplayName,
			))

			resp, err := client.ConversationsV1.UpdateConversationParticipant(incomingParams["ConversationSid"], incomingParams["ParticipantSid"], conversationParams)
			if !web.CheckError(err, false) {
				if resp.ConversationSid != nil {
					log.Println(*resp.ConversationSid)
				} else {
					log.Println(resp.ConversationSid)
				}
			}
		}
	}

	c.Status(http.StatusOK)
}
