package dto

type TriggerDTO struct {
	ID               int64                `json:"id"`
	Type             int64                `json:"type"`
	ProjectID        int64                `json:"projectId"`
	Name             string               `json:"name"`
	Description      string               `json:"description,omitempty"`
	ListID           int64                `json:"listId"`
	DateCreated      int64                `json:"dateCreated,omitempty"`
	DateModified     int64                `json:"dateModified,omitempty"`
	UserCreated      string               `json:"userCreated,omitempty"`
	UserModified     string               `json:"userModified,omitempty"`
	ChannelTypeID    ChannelType          `json:"channelTypeId"`
	ActiveCampaign   int64                `json:"activeCampaign"`
	IgnoreOptOut     bool                 `json:"ignoreOptOut,omitempty"`
	CampaignVersions []CampaignVersionDTO `json:"campaignVersions,omitempty"`
	Active           bool                 `json:"active,omitempty"`
}
