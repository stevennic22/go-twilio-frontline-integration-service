# Frontline Integration Service Example

## !!THIS IS A COMMUNITY DEVELOPED APPLICATION AND IT IS NOT SUPPORTED BY TWILIO!!
## !!IF YOU HAVE ANY ISSUES RELATED TO THIS INTEGRATION SERVICE, PLEASE CREATE A GITHUB ISSUE FIRST!!

This repository contains an example server-side web application that is required to use [Twilio Frontline](https://www.twilio.com/frontline).

It creates the following routes that you will then need to add to your Twilio Frontline Console:

- `/callbacks/crm`
- `/callbacks/outgoing-conversation`
- `/callbacks/templates`
- `/callbacks/routing`
- `/callbacks/twilio-conversations`

## Prerequisites
- A Twilio Account. Don't have one? [Sign up](https://www.twilio.com/try-twilio) for free!
- A golang or Docker installation

## How to start development service (without Docker)

```shell script
# install dependencies
go mod vendor

# build application
go build -v -o ./bin/ ./...

# copy environment variables
cp .env.example .env

# run service
bin/TwFIS
```

## How to build and run Docker image

```shell script
docker build ./ --cpu-shares 512 \
-t stevennic/go-twilio-frontline-integration-service:0.0.3

docker run --name=go-twilio-frontline-integration-service \
--workdir=/usr/src/app \
-p 8082:8082 \
stevennic/go-twilio-frontline-integration-service:0.0.3 TwFIS
```

## Environment variables

```
# Service variables
LISTEN_PORT=8082# default 8082

# Twilio account variables
TWILIO_ACCOUNT_SID=ACXXXX...
TWILIO_AUTH_TOKEN=w2x5y2z6

# Worker Identity variables
Worker_Identity=user@example.com

# Variables for chat configuration
TWILIO_SMS_NUMBER      # Twilio number for incoming/outgoing SMS
TWILIO_WHATSAPP_NUMBER # Twilio number for incoming/outgoing Whatsapp

# Variables for customer configuration
CUSTOMER_PHONE_1=+15557778888
CUSTOMER_PHONE_2=+14443339999
```

## Setting up customers and mapping
The customer data can be configured in [providers/customers.go](providers/customers.go).

### Map between customer address + worker identity pair.
For inbound routing: Used to determine to which worker a new conversation with a particular customer should be routed to.

```js
{
    customerAddress: workerIdentity
}
```

Example:
```golang
var customersToWorkersMap = map[string]string{
	"+15557778888": "user@example.com",
}
```


### Customers list
In the CRM callback response, each [customer object](https://www.twilio.com/docs/frontline/data-transfer-objects#customer) should look like this: 

Example:
```golang

var Customer2 = &Customer{
	Avatar:      "https://example.com/image.jpeg",
	CustomerID:  2,
	DisplayName: "Customer 2",
	Channels: []Channel{
		{
			Type:  "sms",
			Value: "+15557778888",
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
	Worker: "worker@example.com",
}
```

Response format:
```json
objects: {
    customers: customers
}
```
