package providers

import "strings"

type Template struct {
	Body       string `json:"body"`
	WAApproved string `json:"whatsappapproved"`
}

type Content struct {
	Content          string `json:"content"`
	WhatsAppApproved string `json:"whatsappapproved"`
}

type Category struct {
	DisplayName string    `json:"display_name"`
	Templates   []Content `json:"templates"`
}

// Return template with templated strings replaced
func (t *Template) CompileTemplate(c *Customer, author string) string {
	r := strings.NewReplacer(
		"{{Name}}", c.DisplayName,
		"{{FirstName}}", c.FirstName(),
		"{{Author}}", author,
	)

	return (r.Replace(t.Body))
}

// Opener Templates
var OPENER_NEXT_STEPS = Template{
	Body:       "Hello {{Name}} we have now processed your documents and would like to move you on to the next step. Drop me a message. - {{Author}}.",
	WAApproved: "false",
}
var OPENER_NEW_PRODUCT = Template{
	Body:       "Hello {{Name}} we have a new product out which may be of interest to your business. Drop me a message. - {{Author}}.",
	WAApproved: "false",
}
var OPENER_ON_MY_WAY = Template{
	Body:       "Just to confirm I am on my way to your office. {{Name}}.",
	WAApproved: "true",
}

// Reply Templates
var REPLY_SENT = Template{
	Body:       "This has now been sent. - {{Author}}.",
	WAApproved: "false",
}
var REPLY_RATES = Template{
	Body:       "Our rates for any loan are 20% or 30% over $30,000. You can read more at https://example.com. - {{Author}}.",
	WAApproved: "false",
}
var REPLY_OMW = Template{
	Body:       "Just to confirm I am on my way to your office. - {{Author}}.",
	WAApproved: "false",
}
var REPLY_OPTIONS = Template{
	Body:       "Would you like me to go over some options with you {{FirstName}}? - {{Author}}.",
	WAApproved: "false",
}
var REPLY_ASK_DOCUMENTS = Template{
	Body:       "We have a secure drop box for documents. Can you attach and upload them here: https://example.com. - {{Author}}",
	WAApproved: "false",
}

// Closing Templates
var CLOSING_ASK_REVIEW = Template{
	Body:       "Happy to help, {{FirstName}}. If you have a moment could you leave a review about our interaction at this link: https://example.com. - {{Author}}.",
	WAApproved: "false",
}
