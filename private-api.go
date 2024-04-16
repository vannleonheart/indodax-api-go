package indodax

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

func (c *Client) PrivateApiCall(method string, data *map[string]interface{}) (*ResponseBody, error) {
	var respBody ResponseBody

	if err := c.PrivateApiCallWithCustomResult(method, data, &respBody); err != nil {
		return nil, err
	}

	return &respBody, nil
}

func (c *Client) PrivateApiCallWithCustomResult(method string, data *map[string]interface{}, result interface{}) error {
	if c.Config == nil {
		return errors.New("config is not set")
	}

	if len(c.Config.PrivateApiBaseUrl) == 0 {
		return errors.New("private api base url can not be empty")
	}

	if c.Credential == nil {
		return errors.New("credential is required")
	}

	reqBody := map[string]interface{}{
		"method":     method,
		"timestamp":  time.Now().UnixMilli(),
		"recvWindow": DefaultRecvWindow,
	}

	if data != nil {
		for k, v := range *data {
			reqBody[k] = v
		}
	}

	reqHeader := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Key":          c.Credential.TradeApiKey,
		"Sign":         generateSign(reqBody, c.Credential.TradeApiSecret),
	}

	targetUrl := fmt.Sprintf("%s/tapi", c.Config.PrivateApiBaseUrl)

	resp, err := sendRequest(targetUrl, &reqBody, &reqHeader)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(*resp, result); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetInfo() (*GetInfoResponseBody, error) {
	resp, err := c.PrivateApiCall(MethodGetInfo, nil)
	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)
	if err != nil {
		return nil, err
	}

	var ret GetInfoResponseBody

	if err = json.Unmarshal(jsonString, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) GetTransactionHistory(fromDate, toDate string) (*GetTransactionHistoryResponseBody, error) {
	resp, err := c.PrivateApiCall(MethodGetTransactionHistory, &map[string]interface{}{
		"start": fromDate,
		"end":   toDate,
	})

	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)
	if err != nil {
		return nil, err
	}

	var ret GetTransactionHistoryResponseBody

	if err = json.Unmarshal(jsonString, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) GetTradeHistory(pair string, fromId, toId, order *string, since, end, count *int64, orderId *string) (*GetTradeHistoryResponseBody, error) {
	reqBody := map[string]interface{}{
		"pair": pair,
	}

	if fromId != nil {
		reqBody["from_id"] = *fromId
	}

	if toId != nil {
		reqBody["end_id"] = *toId
	}

	if order != nil {
		reqBody["order"] = *order
	}

	if since != nil {
		reqBody["since"] = *since
	}

	if end != nil {
		reqBody["end"] = *end
	}

	if count != nil && *count > 0 && *count < 1000 {
		reqBody["count"] = *count
	}

	if orderId != nil {
		reqBody["order_id"] = *orderId
	}

	resp, err := c.PrivateApiCall(MethodGetTradeHistory, &reqBody)

	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)
	if err != nil {
		return nil, err
	}

	var ret GetTradeHistoryResponseBody

	if err = json.Unmarshal(jsonString, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) GetOpenOrders(pair *string) (interface{}, error) {
	reqBody := map[string]interface{}{}

	if pair != nil {
		reqBody["pair"] = *pair
	}

	resp, err := c.PrivateApiCall(MethodGetOpenOrders, &reqBody)

	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)
	if err != nil {
		return nil, err
	}

	if pair != nil {
		var ret GetPairOpenOrdersResponseBody

		if err = json.Unmarshal(jsonString, &ret); err != nil {
			return nil, err
		}

		return &ret, nil
	} else {
		var ret GetOpenOrdersResponseBody

		if err = json.Unmarshal(jsonString, &ret); err != nil {
			return nil, err
		}

		return &ret, nil
	}
}

func (c *Client) GetOrderHistory(pair string, count, from *int) (*GetOrderHistoryResponseBody, error) {
	reqBody := map[string]interface{}{
		"pair": pair,
	}

	if count != nil && *count > 0 && *count < 1000 {
		reqBody["count"] = *count
	}

	if from != nil {
		reqBody["from"] = *from
	}

	resp, err := c.PrivateApiCall(MethodGetOrderHistory, &reqBody)

	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)
	if err != nil {
		return nil, err
	}

	var ret GetOrderHistoryResponseBody

	if err = json.Unmarshal(jsonString, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) GetOrder(pair, orderId string) (*GetOrderResponseBody, error) {
	resp, err := c.PrivateApiCall(MethodGetOrder, &map[string]interface{}{
		"pair":     pair,
		"order_id": orderId,
	})

	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)

	if err != nil {
		return nil, err
	}

	var ret GetOrderResponseBody

	if err = json.Unmarshal(jsonString, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) GetOrderByClientOrderId(clientOrderId string) (*GetOrderResponseBody, error) {
	resp, err := c.PrivateApiCall(MethodGetOrderByClientOrderId, &map[string]interface{}{
		"client_order_id": clientOrderId,
	})

	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)
	if err != nil {
		return nil, err
	}

	var ret GetOrderResponseBody

	if err = json.Unmarshal(jsonString, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) Trade(tradeType, pair, orderType string, price, amount float64, timeInForce, clientOrderId *string) (*ResponseBody, error) {
	slPair := strings.Split(pair, "_")
	if len(slPair) != 2 {
		return nil, errors.New("invalid pair")
	}

	coinId := slPair[0]
	currencyId := slPair[1]
	reqBody := map[string]interface{}{
		"pair":       pair,
		"type":       tradeType,
		"price":      price,
		"order_type": orderType,
	}

	switch tradeType {
	case TradeTypeSell:
		reqBody[coinId] = amount
	case TradeTypeBuy:
		if orderType == OrderTypeMarket {
			reqBody[currencyId] = amount
		} else {
			reqBody[coinId] = amount
		}
	}

	if timeInForce != nil {
		reqBody["time_in_force"] = *timeInForce
	}

	if clientOrderId != nil {
		reqBody["client_order_id"] = clientOrderId
	}

	resp, err := c.PrivateApiCall(MethodTrade, &reqBody)
	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	return resp, nil
}

func (c *Client) CancelOrder(pair, orderId, tradeType string, orderType *string) (*map[string]interface{}, error) {
	reqBody := map[string]interface{}{
		"pair":     pair,
		"order_id": orderId,
		"type":     tradeType,
	}

	if orderType != nil {
		reqBody["order_type"] = *orderType
	}

	resp, err := c.PrivateApiCall(MethodCancelOrder, &reqBody)

	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)
	if err != nil {
		return nil, err
	}

	var ret map[string]interface{}

	if err = json.Unmarshal(jsonString, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) CancelOrderByClientOrderId(clientOrderId string) (*map[string]interface{}, error) {
	reqBody := map[string]interface{}{
		"client_order_id": clientOrderId,
	}

	resp, err := c.PrivateApiCall(MethodCancelOrderByClientOrderId, &reqBody)

	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)
	if err != nil {
		return nil, err
	}

	var ret map[string]interface{}

	if err = json.Unmarshal(jsonString, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) Withdraw(requestId, currency, address, network, amount, memo string) (*WithdrawCoinResponseBody, error) {
	reqBody := map[string]interface{}{
		"request_id":       requestId,
		"currency":         currency,
		"withdraw_address": address,
		"withdraw_amount":  amount,
	}

	if len(network) > 0 {
		switch network {
		case "bsc":
			network = "bep20"
		case "homestead":
			network = "erc20"
		case "matic":
			network = "polygon"
		}

		reqBody["network"] = network
	}

	if len(memo) > 0 {
		reqBody["withdraw_memo"] = memo
	}

	var result WithdrawCoinResponseBody

	if err := c.PrivateApiCallWithCustomResult(MethodWithdrawCoin, &reqBody, &result); err != nil {
		return nil, err
	}

	if result.Success != 1 {
		if result.Error != nil {
			return nil, errors.New(*result.Error)
		}

		return nil, errors.New("api call failed")
	}

	return &result, nil
}

func (c *Client) GetWithdrawFee(coinId string, coinNetwork *string) (*map[string]interface{}, error) {
	reqBody := map[string]interface{}{
		"currency": coinId,
	}

	if coinNetwork != nil {
		reqBody["network"] = *coinNetwork
	}

	resp, err := c.PrivateApiCall(MethodWithdrawFee, &reqBody)
	if err != nil {
		return nil, err
	}

	if resp.Success != 1 {
		if resp.Error != nil {
			return nil, errors.New(*resp.Error)
		}

		return nil, errors.New("api call failed")
	}

	jsonString, err := json.Marshal(resp.Return)
	if err != nil {
		return nil, err
	}

	var ret map[string]interface{}

	if err = json.Unmarshal(jsonString, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
