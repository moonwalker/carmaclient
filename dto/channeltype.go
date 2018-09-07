package dto

type ChannelType int64

const (
	ChannelTypeUnknown ChannelType = 0
	ChannelTypeEmail   ChannelType = 1
	ChannelTypeSMS     ChannelType = 2
	ChannelTypePush    ChannelType = 3
	ChannelTypeMixed   ChannelType = 4
)