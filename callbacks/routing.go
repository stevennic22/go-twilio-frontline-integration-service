package callbacks

import (
	"log"

	"github.com/stevennic22/TwFIS/providers"
	"github.com/stevennic22/TwFIS/web"
	"github.com/twilio/twilio-go"
	conversations "github.com/twilio/twilio-go/rest/conversations/v1"
)

// Route Conversation to specific worker identity
func routeConversationToWorker(client *twilio.RestClient, conversationSid string, workerIdentity string) {
	params := &conversations.CreateConversationParticipantParams{}
	params.SetIdentity(workerIdentity)

	resp, err := client.ConversationsV1.CreateConversationParticipant(conversationSid, params)
	if !web.CheckError(err, false) {
		log.Printf("Create agent participant: \n%v", err.Error())
	} else {
		if resp.Sid != nil {
			log.Println(*resp.Sid)
		} else {
			log.Println(resp.Sid)
		}
	}
}

// Attempt to route conversation
func RouteConversation(client *twilio.RestClient, conversationSid string, customerNumber string) {
	var workerIdentity string

	fWFC := providers.FindWorkerForCustomer(customerNumber)

	if fWFC != "" {
		workerIdentity = fWFC
	} else {
		fRW := providers.FindRandomWorker()

		if fRW != "" {
			workerIdentity = fRW
		} else {
			log.Printf("Routing failed, please add workers to customersToWorkersMap or define a default worker | { conversationSid: %s }", conversationSid)
		}
	}

	if workerIdentity != "" {
		routeConversationToWorker(client, conversationSid, workerIdentity)
	}
}
