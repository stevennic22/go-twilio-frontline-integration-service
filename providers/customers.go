package providers

import (
	"math/rand"
	"strings"
	"time"
)

type Link struct {
	Type        string `json:"type"`
	Value       string `json:"value"`
	DisplayName string `json:"display_name"`
}

type Channel struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Detail struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Customer struct {
	CustomerID  int       `json:"customer_id"`
	DisplayName string    `json:"display_name"`
	Channels    []Channel `json:"channels"`
	Links       []Link    `json:"links"`
	Details     Detail    `json:"details"`
	Worker      string    `json:"worker"`
	Avatar      string    `json:"avatar"`
}

type TrimmedCustomer struct {
	CustomerID  int    `json:"customer_id"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
}

func (c *Customer) ReturnTrimmed() *TrimmedCustomer {
	return (&TrimmedCustomer{
		CustomerID:  c.CustomerID,
		DisplayName: c.DisplayName,
		Avatar:      c.Avatar,
	})
}

func (c *Customer) FirstName() string {
	return (c.DisplayName[:strings.Index(c.DisplayName, " ")])
}

func ValInStringSlice(argList []string, x string) bool {
	for _, y := range argList {
		if y == x {
			return true
		}
	}
	return false
}

var CustomersToWorkersMap = map[string]string{}

func FindWorkerForCustomer(customerNumber string) string {
	return (CustomersToWorkersMap[customerNumber])
}

func FindRandomWorker() string {
	// Use customerToWorkersMap to return list of unique workers
	// Use built workers list to return a random index to randomly assign
	var workers []string

	for _, v := range CustomersToWorkersMap {
		if !ValInStringSlice(workers, v) {
			workers = append(workers, v)
		}
	}

	if len(CustomersToWorkersMap) > 0 {
		rand.Seed(time.Now().UnixNano())
		randomWorkerIndex := rand.Intn(len(workers))
		return (workers[randomWorkerIndex])
	} else {
		return ("")
	}
}

// Get list of customers and returned trimmed details
func GetCustomersList(worker string, pageSize int, anchor int) []*TrimmedCustomer {
	var trimmedcustomers []*TrimmedCustomer

	var limit = (anchor + pageSize)

	if limit > len(Customers) {
		limit = len(Customers)
	}

	for x := anchor; x < limit; x++ {
		trimmedcustomers = append(trimmedcustomers, Customers[x].ReturnTrimmed())
	}
	return (trimmedcustomers)
}

// Get customer by contact number
func GetCustomerByNumber(customerNumber string) *Customer {
	var customer = &Customer{}
	for _, v := range Customers {
		for _, n := range v.Channels {
			if n.Value == customerNumber {
				customer = v
			}
		}
	}

	return (customer)
}

// Get customer by ID
func GetCustomerById(cID int) *Customer {
	var customer = &Customer{}
	for _, v := range Customers {
		if v.CustomerID == cID {
			customer = v
		}
	}

	return (customer)
}

// Generation of integrated customer list
var Customers []*Customer

func SetUpCustomerList(passedWorker string, customerNumber1 string, customerNumber2 string) {

	var customer1 = &Customer{
		Avatar:      "https://example.com/image.jpeg",
		CustomerID:  1,
		DisplayName: "Customer 1",
		Channels: []Channel{
			{
				Type:  "sms",
				Value: customerNumber1,
			},
			{
				Type:  "chat",
				Value: passedWorker,
			},
		},
		Links: []Link{
			{
				Type:        "Facebook",
				Value:       "https://meta.facebook.com/",
				DisplayName: "Meta",
			},
		},
		Details: Detail{
			Title:   "Purchase History",
			Content: "Product: Unobtanium\n\nDate: 2021-07-07\nQuantity: 1000 units\n\nSales rep: Ash Williams",
		},
		Worker: passedWorker,
	}

	var customer2 = &Customer{
		Avatar:      "https://example.com/image.jpeg",
		CustomerID:  2,
		DisplayName: "Customer 2",
		Channels: []Channel{
			{
				Type:  "sms",
				Value: customerNumber2,
			},
		},
		Links: []Link{
			{
				Type:        "Facebook",
				Value:       "https://meta.facebook.com/",
				DisplayName: "Facebook",
			},
		},
		Details: Detail{
			Title:   "Purchase History",
			Content: "Product: Unobtanium\n\nDate: 2021-07-07\nQuantity: 1000 units\n\nSales rep: Ash Williams",
		},
		Worker: passedWorker,
	}

	Customers = append(Customers, customer1)
	Customers = append(Customers, customer2)

	CustomersToWorkersMap[customerNumber1] = passedWorker
	CustomersToWorkersMap[customerNumber2] = passedWorker
}
