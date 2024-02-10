package paygo

import (
	"net/http"
)

type Session struct {
	Client *http.Client

	AccessToken string

	Debug    bool
	logLevel int

	AppVersion string
	ClientName string

	DeviceUuid string
	ClientUuid string
	UserAgent  string
}
type Devices struct {
	Samsung []DeviceInfo `json:"samsung"`
}

type DeviceInfo struct {
	ModelName        string `json:"modelName"`
	OsVersion        string `json:"osVersion"`
	OsReleaseVersion string `json:"osReleaseVersion"`
	HardwareName     string `json:"hardwareName"`
}

type P2PUserProfile struct {
	CustomDisplayName *string `json:"customDisplayName"`
	DisplayName       string  `json:"displayName"`
	ExternalUserId    string  `json:"externalUserId"`
	IconImageUrl      string  `json:"iconImageUrl"`
}

type P2PTheme struct {
	Id                           string  `json:"id"`
	Title                        *string `json:"title"`
	IconImageUrl                 string  `json:"iconImageUrl"`
	PendingThumbnailAnimationUrl *string `json:"pendingThumbnailAnimationUrl"`
	ThumbnailAnimationUrl        string  `json:"thumbnailAnimationUrl"`
	BackgroundAnimationUrl       string  `json:"backgroundAnimationUrl"`
	OpenAnimationUrl             *string `json:"openAnimationUrl"`
	ShouldHideAmount             bool    `json:"shouldHideAmount"`
}

type P2PBankTransferData struct {
	Amount         int    `json:"amount"`
	ChargeAmount   int    `json:"chargeAmount"`
	OrderId        string `json:"orderId"`
	OrderType      string `json:"orderType"`
	PayoutMethodId string `json:"payoutMethodId"`
	RecipientName  string `json:"recipientName"`
	Status         string `json:"status"`
}
type P2PMessageGroupPayData struct {
	GroupPayStatus        string   `json:"groupPayStatus"`
	IconImageUrls         []string `json:"iconImageUrls"`
	Name                  string   `json:"name"`
	ParticipantAmount     *int     `json:"participantAmount"`
	ParticipantStatus     *string  `json:"participantStatus"`
	SplitBillId           int64    `json:"splitBillId"`
	TotalAmount           int64    `json:"totalAmount"`
	TotalAmountPaid       int64    `json:"totalAmountPaid"`
	TotalAmountUnpaid     int64    `json:"totalAmountUnpaid"`
	TotalParticipantCount int      `json:"totalParticipantCount"`
	UserType              string   `json:"userType"`
}

type P2PLocalized struct {
	En string `json:"en"`
	Ja string `json:"ja"`
}
type P2PSubWalletSplit struct {
	ReceiverEmoneyAmount  int64 `json:"receiverEmoneyAmount"`
	ReceiverPrepaidAmount int64 `json:"receiverPrepaidAmount"`
	SenderEmoneyAmount    int64 `json:"senderEmoneyAmount"`
	SenderPrepaidAmount   int64 `json:"senderPrepaidAmount"`
}

type P2PMessageData struct {
	Amount                *int64                  `json:"amount"`
	BankTransferData      *P2PBankTransferData    `json:"bankTransferData"`
	Expiry                *string                 `json:"expiry"`
	FailureType           *string                 `json:"failureType"`
	GroupPayData          *P2PMessageGroupPayData `json:"groupPayData"`
	IsLink                *bool                   `json:"isLink"`
	IsQr                  *bool                   `json:"isQr"`
	IsResendable          *bool                   `json:"isResendable"`
	MyCodeLink            *string                 `json:"myCodeLink"`
	NotificationText      *P2PLocalized           `json:"notificationText"`
	OrderId               *string                 `json:"orderId"`
	OrderType             *string                 `json:"orderType"`
	RequestId             *string                 `json:"requestId"`
	RequestMoneyId        *string                 `json:"requestMoneyId"`
	SendMoneyLink         *string                 `json:"sendMoneyLink"`
	SendMoneyLinkPasscode *string                 `json:"sendMoneyLinkPasscode"`
	Sender                *P2PUserProfile         `json:"sender"`
	Status                *string                 `json:"status"`
	SubWalletSplit        *P2PSubWalletSplit      `json:"subWalletSplit"`
	Theme                 *P2PTheme               `json:"theme"`
	TransactionAt         *string                 `json:"transactionAt"`
	TransactionType       *string                 `json:"transactionType"`
	Type                  *string                 `json:"type"`
}

type P2PMessageInfo struct {
	Message               string         `json:"message"`
	MessageId             string         `json:"messageId"`
	CustomType            string         `json:"customType"`
	User                  P2PUserProfile `json:"user"`
	AndroidMinimumVersion string         `json:"androidMinimumVersion"`
	IosMinimumVersion     string         `json:"iosMinimumVersion"`
}

type P2PAmount struct {
	Amount int64  `json:"amount"`
	Label  string `json:"label"`
	Type   string `json:"type"`
}

type PendingP2PInfo struct {
	AcceptedAt      *string    `json:"acceptedAt"`
	Amount          int64      `json:"amount"`
	AmountList      *P2PAmount `json:"amountList"`
	CreatedAt       string     `json:"createdAt"`
	ExpiredAt       string     `json:"expiredAt"`
	ImageUrl        string     `json:"imageUrl"`
	IsSetPasscode   bool       `json:"isSetPasscode"`
	Link            *string    `json:"link"`
	Message         *string    `json:"message"`
	MoneyPriority   *any       `json:"moneyPriority"`
	OrderId         string     `json:"orderId"`
	OrderType       string     `json:"orderType"`
	Passcode        *string    `json:"passcode"`
	SubDescription  *string    `json:"subDescription"`
	TextDescription *string    `json:"textDescription"`
}
type P2PLinkInfo struct {
	Message        *P2PMessageInfo `json:"message"`
	OrderStatus    string          `json:"orderStatus"`
	PendingP2PInfo PendingP2PInfo  `json:"pendingP2PInfo"`
	Receiver       *P2PUserProfile `json:"receiver"`
	Sender         *P2PUserProfile `json:"sender"`
}
