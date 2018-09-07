package dto

type SendTriggerDTO struct {
	DeliveryTime                 string            `json:"deliveryTime,omitempty"` // When is the trigger to be sent? Trigger will be sent immediately if the value is omitted, format shall be yyyy-MM-ddTHH:mm:ss-zzzz -zzzz is timezone difference from UCT sweden is +01000 , e.g. 2016-02-09T11:00:00+0100,
	OptOut                       bool              `json:"optOut,omitempty"`       // if true the user will unsubscribed,
	SaveProps                    bool              `json:"saveProps,omitempty"`    // True to update propertydata on the contact,
	OriginalId                   string            `json:"originalId"`             // originalId of the contact the trigger is to be sent to,
	CustomerId                   int64             `json:"customerId"`
	CampaignId                   int64             `json:"campaignId,omitempty"`                   // If supplied the trigger willsend this delivery instead of the triggers active delivery,
	Properties                   map[string]string `json:"properties"`                             // the data for the trigger,
	ForceCreateMissingProperties bool              `json:"forceCreateMissingProperties,omitempty"` // Automatically create missing properties. If false, missing property names will throw exception
}
