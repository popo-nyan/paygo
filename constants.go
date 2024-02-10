package paygo

// PayPayAppVersion Android版のPayPayアプリのバージョン
const PayPayAppVersion string = "4.31.0"

// HmacKey https://push-config.paypay.ne.jp//v1/config等のリクエストに必要なHMAC-SHA256を計算する為の鍵
const HmacKey string = "717198a8bcf3405384abcb4815d1efb9"

// ClientName x-requesterの値
const ClientName string = "android-pp"

type OrderStatus string

const (
	OrderStatusSuccess    OrderStatus = "SUCCESS"
	OrderStatusFailed     OrderStatus = "FAILED"
	OrderStatusPending    OrderStatus = "PENDING"
	OrderStatusCreated    OrderStatus = "CREATED"
	OrderStatusRejected   OrderStatus = "REJECTED"
	OrderStatusExpired    OrderStatus = "EXPIRED"
	OrderStatusCanceled   OrderStatus = "CANCELED"
	OrderStatusAuthorized OrderStatus = "AUTHORIZED"
	OrderStatusUnknown    OrderStatus = "AUTHORIZED"
)

// PayPayResultCode app4.paypay.ne.jpのレスポンスで返ってくるコード
type PayPayResultCode string

const (
	PayPayResultSuccess                                            PayPayResultCode = "S0000"
	PayPayResultRefreshAccessTokenRequired                         PayPayResultCode = "S0001"
	PayPayResultAppVersionError                                    PayPayResultCode = "S1000"
	PayPayResultUnderMaintenance                                   PayPayResultCode = "S1001"
	PayPayResultYidLinkRequired                                    PayPayResultCode = "S1002"
	PayPayResultSessionFailure                                     PayPayResultCode = "S1003"
	PayPayResultSmsAuthRequired                                    PayPayResultCode = "S1004"
	PayPayResultNeedChargeForP2P                                   PayPayResultCode = "S1005"
	PayPayResultPaymentFailedToSettleTimeout                       PayPayResultCode = "S2004"
	PayPayResultPaymentFailedInsufficientWalletBalance             PayPayResultCode = "S2013"
	PayPayResultMerchantScanPaymentFailedInsufficientWalletBalance PayPayResultCode = "S2014"
	PayPayResultChargeFailedTimeout                                PayPayResultCode = "S2106"
	PayPayResultP2pFailedTimeout                                   PayPayResultCode = "S2206"
	PayPayResultShowToast                                          PayPayResultCode = "S2017"
	PayPayResultChargeFailedGiftCardTimeOut                        PayPayResultCode = "S2120"
	PayPayResultP1pFailedTimeout                                   PayPayResultCode = "S2206"
	PayPayResultP2pSearchFailedNotFound                            PayPayResultCode = "S2210"
	PayPayResultP2pSearchFailedMyself                              PayPayResultCode = "S2212"
	PayPayResultPailedToFetchPreAuthDisplayInfo                    PayPayResultCode = "S2700"
	PayPayResultSettingPasswordFailed                              PayPayResultCode = "S3001"
	PayPayResultSettingNicknameFailed                              PayPayResultCode = "S3003"
	PayPayResultSettinglastNameFailed                              PayPayResultCode = "S3004"
	PayPayResultSettinglastKanaNameFailed                          PayPayResultCode = "S3005"
	PayPayResultSettingfirstNameFailed                             PayPayResultCode = "S3006"
	PayPayResultSettingfirstKanaNameFailed                         PayPayResultCode = "S3007"
	PayPayResultRegisterPayPayIdFailed                             PayPayResultCode = "S3009"
	PayPayResultRegisterPayPayIdValidateFailed                     PayPayResultCode = "S3010"
	PayPayResultAuthSmsFailedOTPMismatch                           PayPayResultCode = "S3103"
	PayPayResultRiskNG                                             PayPayResultCode = "S4002"
	PayPayResultNotTpointLinked                                    PayPayResultCode = "S4032"
	PayPayResultSendBirdMessageFailed                              PayPayResultCode = "S4053"
	PayPayResultFacingServerError                                  PayPayResultCode = "S5000"
	PayPayResultServerTimeoutError                                 PayPayResultCode = "S5001"
	PayPayResultCircuitBreakerError                                PayPayResultCode = "S5003"
	PayPayResultGenericErrorWithDeeplink                           PayPayResultCode = "S5004"
	PayPayResultCustomErrorResponse                                PayPayResultCode = "S9999"
	PayPayResultInsufficientHeaderParameter                        PayPayResultCode = "C0000"
	PayPayResultInvalidRequestHeader                               PayPayResultCode = "C0001"
	PayPayResultProhibitionOverseasAccessError                     PayPayResultCode = "C1002"
	PayPayResultIgnoreProhibitionOverseasAccess                    PayPayResultCode = "C1003"
	PayPayResultServerError                                        PayPayResultCode = "E0000"
	PayPayResultKycAddressParseError                               PayPayResultCode = "S03200006"
	PayPayResultKycUserInfoParseError                              PayPayResultCode = "S03200007"
	PayPayResultKycNameExternalCharacterFoundError                 PayPayResultCode = "S03200008"
	PayPayResultSoftLimitUpdateError                               PayPayResultCode = "S9998"
)
