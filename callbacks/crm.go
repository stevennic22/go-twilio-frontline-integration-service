package callbacks

import (
	"log"
	"net/http"
	"strconv"

	"github.com/stevennic22/TwFIS/providers"

	"github.com/gin-gonic/gin"
)

// Return Customer List with trimmed details
func HandleGetCustomersListCallback(c *gin.Context, params map[string]string) {
	log.Println("Getting Customers list")

	workerIdentity := params["Worker"]

	var pageSize = 30
	pageSize, _ = strconv.Atoi(params["PageSize"])

	var anchor = 0
	anchor, _ = strconv.Atoi(params["Anchor"])

	var customersList []*providers.TrimmedCustomer
	customersList = providers.GetCustomersList(workerIdentity, pageSize, anchor)

	c.JSON(http.StatusOK, gin.H{
		"objects": gin.H{
			"customers":  customersList,
			"searchable": "true",
		},
	})
}

// Return Customer Details based on provided Customer ID
func HandleGetCustomerDetailsByCustomerIdCallback(c *gin.Context, params map[string]string) {
	customerId, _ := strconv.Atoi(params["CustomerId"])

	customer := providers.GetCustomerById(customerId)

	c.JSON(http.StatusOK, gin.H{
		"objects": gin.H{
			"customer": gin.H{
				"customer_id":  customer.CustomerID,
				"display_name": customer.DisplayName,
				"channels":     customer.Channels,
				"links":        customer.Links,
				"avatar":       customer.Avatar,
				"details":      customer.Details,
			},
		},
	})
}
