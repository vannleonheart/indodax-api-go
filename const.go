package indodax

const (
	MethodGetInfo                    = "getInfo"
	MethodGetTransactionHistory      = "transHistory"
	MethodGetTradeHistory            = "tradeHistory"
	MethodGetOpenOrders              = "openOrders"
	MethodGetOrderHistory            = "orderHistory"
	MethodGetOrder                   = "getOrder"
	MethodGetOrderByClientOrderId    = "getOrderByClientOrderId"
	MethodTrade                      = "trade"
	MethodCancelOrder                = "cancelOrder"
	MethodCancelOrderByClientOrderId = "cancelByClientOrderId"
	MethodWithdrawFee                = "withdrawFee"
	MethodWithdrawCoin               = "withdrawCoin"

	TradeTypeBuy  = "buy"
	TradeTypeSell = "sell"

	OrderTypeMarket    = "market"
	OrderTypeLimit     = "limit"
	OrderTypeStop      = "stop"
	OrderTypeStopLimit = "stoplimit"

	TimeInForceGoodTillCancel = "GTC"
	TimeInForceMakerOrCancel  = "MOC"

	DefaultRecvWindow = 5000

	TimeFrame1Minute   = "1"
	TimeFrame15Minutes = "15"
	TimeFrame30Minutes = "30"
	TimeFrame1Hour     = "60"
	TimeFrame4Hours    = "240"
	TimeFrame1Day      = "1D"
	TimeFrame3Days     = "3D"
	TimeFrame1Week     = "1W"
)
