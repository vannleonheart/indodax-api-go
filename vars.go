package indodax

import "encoding/json"

type Client struct {
	Config     *Config
	Credential *Credential
}

type Config struct {
	PublicApiBaseUrl  string `json:"public_api_base_url"`
	PrivateApiBaseUrl string `json:"private_api_base_url"`
}

type Credential struct {
	TradeApiKey    string `json:"trade_api_key"`
	TradeApiSecret string `json:"trade_api_secret"`
}

type RequestBody struct {
	Method     string   `json:"method"`
	Timestamp  int64    `json:"timestamp,omitempty"`
	RecvWindow int      `json:"recvWindow,omitempty"`
	OrderId    *string  `json:"order_id,omitempty"`
	Pair       *string  `json:"pair,omitempty"`
	Type       *string  `json:"type,omitempty"`
	Price      *float64 `json:"price,omitempty"`
}

type ResponseBody struct {
	Success   uint8       `json:"success"`
	Error     *string     `json:"error,omitempty"`
	ErrorCode *string     `json:"error_code,omitempty"`
	Return    interface{} `json:"return,omitempty"`
}

type GetInfoResponseBody struct {
	UserId             json.Number                `json:"user_id"`
	Name               string                     `json:"name"`
	Email              string                     `json:"email"`
	ProfilePicture     string                     `json:"profile_picture"`
	VerificationStatus string                     `json:"verification_status"`
	GauthEnable        bool                       `json:"gauth_enable"`
	WithdrawStatus     int                        `json:"withdraw_status"`
	Balance            map[string]json.Number     `json:"balance"`
	BalanceHold        map[string]json.Number     `json:"balance_hold"`
	Network            map[string]interface{}     `json:"network"`
	MemoIsRequired     map[string]map[string]bool `json:"memo_is_required"`
	Address            map[string]string          `json:"address"`
	ServerTime         int64                      `json:"server_time"`
}

type GetTransactionHistoryResponseBody struct {
	Withdraw map[string][]map[string]interface{} `json:"withdraw"`
	Deposit  map[string][]map[string]interface{} `json:"deposit"`
}

type GetTradeHistoryResponseBody struct {
	Trades []map[string]interface{} `json:"trades"`
}

type GetPairOpenOrdersResponseBody struct {
	Orders []map[string]interface{} `json:"orders"`
}

type GetOpenOrdersResponseBody struct {
	Orders map[string][]map[string]interface{} `json:"orders"`
}

type GetOrderHistoryResponseBody struct {
	Orders []map[string]interface{} `json:"orders"`
}

type GetOrderResponseBody struct {
	Order map[string]interface{} `json:"order"`
}

type WithdrawCoinResponseBody struct {
	ResponseBody
	Status           string `json:"status"`
	WithdrawCurrency string `json:"withdraw_currency"`
	WithdrawAddress  string `json:"withdraw_address"`
	WithdrawAmount   string `json:"withdraw_amount"`
	Fee              string `json:"fee"`
	AmountAfterFee   string `json:"amount_after_fee"`
	SubmitTime       string `json:"submit_time"`
	WithdrawId       string `json:"withdraw_id"`
	TxId             string `json:"tx_id"`
}

type GetServerTimeResponseBody struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"server_time"`
}

type GetPairsResponseBody struct {
	Pairs []Pair `json:"pairs"`
}

type GetPriceIncrementsResponseBody struct {
	Increments map[string]interface{} `json:"increments"`
}

type GetSummariesResponseBody struct {
	Tickers   map[string]map[string]interface{} `json:"tickers"`
	Prices24H map[string]interface{}            `json:"prices_24h"`
	Prices7D  map[string]interface{}            `json:"prices_7d"`
}

type GetTickerResponseBody struct {
	Ticker map[string]interface{} `json:"ticker"`
}

type GetTickerAllResponseBody struct {
	Tickers map[string]map[string]interface{} `json:"tickers"`
}

type GetDepthResponseBody struct {
	Buy  [][2]json.Number `json:"buy"`
	Sell [][2]json.Number `json:"sell"`
}

type Pair struct {
	Id                     string      `json:"id"`
	Symbol                 string      `json:"symbol"`
	BaseCurrency           string      `json:"base_currency"`
	TradedCurrency         string      `json:"traded_currency"`
	TradedCurrencyUnit     string      `json:"traded_currency_unit"`
	Description            string      `json:"description"`
	TickerId               string      `json:"ticker_id"`
	VolumePrecision        json.Number `json:"volume_precision"`
	PricePrecision         json.Number `json:"price_precision"`
	PriceRound             json.Number `json:"price_round"`
	PriceScale             json.Number `json:"pricescale"`
	TradeMinBaseCurrency   json.Number `json:"trade_min_base_currency"`
	TradeMinTradedCurrency json.Number `json:"trade_min_traded_currency"`
	TradeFeePercent        json.Number `json:"trade_fee_percent"`
	TradeFeePercentTaker   json.Number `json:"trade_fee_percent_taker"`
	TradeFeePercentMaker   json.Number `json:"trade_fee_percent_maker"`
	HasMemo                bool        `json:"has_memo"`
	MemoName               interface{} `json:"memo_name"`
	UrlLogo                string      `json:"url_logo"`
	UrlLogoPng             string      `json:"url_logo_png"`
	IsMaintenance          int         `json:"is_maintenance"`
	IsMarketSuspended      int         `json:"is_market_suspended"`
	CmcId                  interface{} `json:"cmc_id"`
	CoingeckoId            string      `json:"coingecko_id"`
}

type Trade struct {
	Date   string      `json:"date"`
	Price  json.Number `json:"price"`
	Amount json.Number `json:"amount"`
	Tid    string      `json:"tid"`
	Type   string      `json:"type"`
}

type OHLC struct {
	Time   int64       `json:"Time"`
	Open   json.Number `json:"Open"`
	High   json.Number `json:"High"`
	Low    json.Number `json:"Low"`
	Close  json.Number `json:"Close"`
	Volume string      `json:"Volume"`
}
