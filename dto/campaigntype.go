package dto

type CampaignType int64

const (
	CampaignTypeEmail                   CampaignType = 1
	CampaignTypeSMS                     CampaignType = 2
	CampaignTypeDeleted                 CampaignType = 4
	CampaignTypeNormalApproved          CampaignType = 5
	CampaignTypeIncomingSMS             CampaignType = 6
	CampaignTypeSMSWithOADC             CampaignType = 8
	CampaignTypeEmailWithAttachement    CampaignType = 14
	CampaignTypeEmailSubdelivey         CampaignType = 16
	CampaignTypeEmailMasterDelivery     CampaignType = 17
	CampaignTypeSMSLinkMobilityDelivery CampaignType = 18
	CampaignTypeAppPushAndroidDelivery  CampaignType = 19
	CampaignTypeAppPushIOSDelivery      CampaignType = 20
	CampaignTypeAppPushWindowsDelivery  CampaignType = 21
)
