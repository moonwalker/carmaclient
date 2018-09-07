package dto

type ContactDTO struct {
	ID                        int64             `json:"id,omitempty"`
	ListID                    int64             `json:"listId"`
	Country                   string            `json:"country,omitempty"`
	OriginalID                string            `json:"originalId"`
	OriginalIDHashed          int64             `json:"originalIDHashed,omitempty"`
	FirstName                 string            `json:"firstName,omitempty"`
	LastName                  string            `json:"lastName,omitempty"`
	MiddleName                string            `json:"middleName,omitempty"`
	EmailAddress              string            `json:"emailAddress,omitempty"`
	Title                     string            `json:"title,omitempty"`
	DateOfBirth               string            `json:"dateOfBirth,omitempty"`
	City                      string            `json:"city,omitempty"`
	ZipCode                   string            `json:"zipcode,omitempty"`
	Sex                       string            `json:"sex,omitempty"`
	MobileNumber              string            `json:"mobileNumber,omitempty"`
	Unsubscribed              bool              `json:"unsubscribed"`
	Bounced                   bool              `json:"bounced"`
	MobileUnsubscribed        bool              `json:"mobileUnsubscribed"`
	AudioUnsubscribed         bool              `json:"audioUnsubscribed"`
	PreferredContentVersionID int64             `json:"preferredContentVersionId,omitempty"`
	OptOutDate                int64             `json:"optOutDate"`
	DateOfInvalidation        int64             `json:"dateOfInvalidation"`
	OptOutMobileDate          int64             `json:"optOutMobileDate"`
	OptOutAudioDate           int64             `json:"optOutAudioDate"`
	Active                    bool              `json:"active"`
	DateCreated               int64             `json:"dateCreated,omitempty"`
	DateModified              int64             `json:"dateModified,omitempty"`
	BlockedUntil              int64             `json:"blockedUntil,omitempty"`
	Properties                map[string]string `json:"properties,omitempty"`
	NamedProperties           map[string]string `json:"namedProperties,omitempty"`
	MobileNumberValid         bool              `json:"mobileNumberValid,omitempty"`
	//deviceInfo (java.util.LinkedList, optional),
}
