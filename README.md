### Public API

Public API, is a publicly available INDODAX API that is called without using credentials and does not contain any related account information.

```go
package main

import "indodax"

func main()  {
	
	idx := indodax.New(&indodax.Config{
            PublicApiBaseUrl:  "https://indodax.com",
            PrivateApiBaseUrl: "https://indodax.com",
	})

	result, err := idx.GetTicker("btcidr")

	if err != nil {
            panic(err)
	}

	fmt.Printf("%+v", *result)

}
```

#### List of Public API Functions
```
func (c *Client) GetServerTime() (*GetServerTimeResponseBody, error)
func (c *Client) GetPairs() (*[]Pair, error)
func (c *Client) GetPriceIncrements() (*GetPriceIncrementsResponseBody, error)
func (c *Client) GetSummaries() (*GetSummariesResponseBody, error)
func (c *Client) GetTicker(pairId string) (*GetTickerResponseBody, error)
func (c *Client) GetTickerAll() (*GetTickerAllResponseBody, error)
func (c *Client) GetTrades(pairId string) (*[]Trade, error)
func (c *Client) GetDepth(pairId string) (*GetDepthResponseBody, error)
func (c *Client) GetOHLCHistory(pairId, timeFrame string, from, to int64) (*[]OHLC, error)
```


### Private API

Private API, is INDODAX API that is called using credentials and contains any related account information.

```go
package main

import "indodax"

func main() {
	
	tradeApiKey := "xxxx"
	tradeApiSecret := "yyyy"

	idx := indodax.New(&indodax.Config{
            PublicApiBaseUrl:  "https://indodax.com",
            PrivateApiBaseUrl: "https://indodax.com",
	})

	result, err := idx.WithCredential(tradeApiKey, tradeApiSecret).GetInfo()

	if err != nil {
            panic(err)
	}

	fmt.Printf("%+v", *result)
	
}
```

#### List of Private API Functions
```
func (c *Client) GetInfo() (*GetInfoResponseBody, error)
func (c *Client) GetTransactionHistory(fromDate, toDate string) (*GetTransactionHistoryResponseBody, error)
func (c *Client) GetTradeHistory(pair string, fromId, toId, order *string, since, end, count *int64, orderId *string) (*GetTradeHistoryResponseBody, error)
func (c *Client) GetOpenOrders(pair *string) (interface{}, error)
func (c *Client) GetOrderHistory(pair string, count, from *int) (*GetOrderHistoryResponseBody, error)
func (c *Client) GetOrder(pair, orderId string) (*GetOrderResponseBody, error)
func (c *Client) GetOrderByClientOrderId(clientOrderId string) (*GetOrderResponseBody, error)
func (c *Client) Trade(tradeType, pair, orderType string, price, amount float64, timeInForce, clientOrderId *string) (*ResponseBody, error)
func (c *Client) CancelOrder(pair, orderId, tradeType string, orderType *string) (*map[string]interface{}, error)
func (c *Client) CancelOrderByClientOrderId(clientOrderId string) (*map[string]interface{}, error)
func (c *Client) Withdraw(requestId, currency, address, network, amount, memo string)
func (c *Client) GetWithdrawFee(coinId string, coinNetwork *string)
```

### References
- [INDODAX API](https://github.com/btcid/indodax-official-api-docs/blob/master/Marketdata-websocket.md) 