package dto

type CampaignVersionDTO struct {
	CampaignID     int64   `json:"campaignId"`
	CampaignName   string  `json:"campaignName"`
	CampaignTypeID int64   `json:"campaignTypeId"`
	Active         bool    `json:"active"`
	Locked         bool    `json:"locked"`
	DateCreated    int64   `json:"dateCreated"`
	DateModified   int64   `json:"dateModified"`
	UserCreated    string  `json:"userCreated"`
	UserModified   string  `json:"userModified"`
	Stats          StatDTO `json:"statDtos"`
	LastActivated  int64  `json:"lastActivated"`
}
