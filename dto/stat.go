package dto

type StatDTO struct {
	ID              int64   `json:"id"`
	DeliveryID      int64   `json:"deliveryId"`
	Source          int64   `json:"source"`
	ProjectID       int64   `json:"projectId"`
	DeliveryTypeID  int64   `json:"deliveryTypeId"`
	Recipients      int64   `json:"recipients"`
	Sent            int64   `json:"sent"`
	Opened          int64   `json:"opened"`
	Clicked         int64   `json:"clicked"`
	OpenedInBrowser int64   `json:"openedInBrower"`
	NonDeliverable  int64   `json:"nonDeliverable"`
	Unsubscribe     int64   `json:"unsubscribed"`
	NotBounced      int64   `json:"notBounced"`
	SoftBounced     int64   `json:"softBounced"`
	HardBounced     int64   `json:"hardBounced"`
	Complaints      int64   `json:"complaints"`
	ShareTarget     int64   `json:"shareTarget"`
	MailClick       int64   `json:"mailClick"`
	Conversions     int64   `json:"conversions"`
	UniqueClick     int64   `json:"uniqueClick"`
	Amount          float64 `json:"amount"`
}
