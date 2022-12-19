package callbacks

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stevennic22/TwFIS/providers"
)

// Get list of Templates specifically for Customer ID
func HandleGetTemplatesByCustomerIdCallback(c *gin.Context, params map[string]string) {
	customerId, _ := strconv.Atoi(params["CustomerId"])
	log.Printf("Getting templates for: %d", customerId)

	worker := params["Worker"]

	customer := providers.GetCustomerById(customerId)

	if customer.DisplayName == "" {
		c.String(http.StatusNotFound, "404-FILENOTFOUND")
	}

	var openersCategory = providers.Category{
		DisplayName: "Openers",
		Templates: []providers.Content{
			{Content: providers.OPENER_NEXT_STEPS.CompileTemplate(customer, worker), WhatsAppApproved: providers.OPENER_NEXT_STEPS.WAApproved},
			{Content: providers.OPENER_NEW_PRODUCT.CompileTemplate(customer, worker), WhatsAppApproved: providers.OPENER_NEW_PRODUCT.WAApproved},
			{Content: providers.OPENER_ON_MY_WAY.CompileTemplate(customer, worker), WhatsAppApproved: providers.OPENER_ON_MY_WAY.WAApproved},
		},
	}

	var repliesCategory = providers.Category{
		DisplayName: "Replies",
		Templates: []providers.Content{
			{Content: providers.REPLY_SENT.CompileTemplate(customer, worker), WhatsAppApproved: providers.REPLY_SENT.WAApproved},
			{Content: providers.REPLY_RATES.CompileTemplate(customer, worker), WhatsAppApproved: providers.REPLY_RATES.WAApproved},
			{Content: providers.REPLY_OMW.CompileTemplate(customer, worker), WhatsAppApproved: providers.REPLY_OMW.WAApproved},
			{Content: providers.REPLY_OPTIONS.CompileTemplate(customer, worker), WhatsAppApproved: providers.REPLY_OPTIONS.WAApproved},
			{Content: providers.REPLY_ASK_DOCUMENTS.CompileTemplate(customer, worker), WhatsAppApproved: providers.REPLY_ASK_DOCUMENTS.WAApproved},
		},
	}

	var closingCategory = providers.Category{
		DisplayName: "Closing",
		Templates: []providers.Content{
			{Content: providers.CLOSING_ASK_REVIEW.CompileTemplate(customer, worker), WhatsAppApproved: providers.CLOSING_ASK_REVIEW.WAApproved},
		},
	}

	var categories []providers.Category

	categories = append(categories, openersCategory)
	categories = append(categories, repliesCategory)
	categories = append(categories, closingCategory)

	c.JSON(http.StatusOK, categories)
}
