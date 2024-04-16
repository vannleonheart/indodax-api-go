package indodax

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) PublicApiCall(uri string) (*[]byte, error) {
	uri = strings.TrimSpace(strings.Trim(uri, "/"))

	if len(uri) == 0 {
		return nil, errors.New("uri can not be empty")
	}

	if c.Config == nil {
		return nil, errors.New("config is not set")
	}

	if len(c.Config.PublicApiBaseUrl) == 0 {
		return nil, errors.New("public api base url can not be empty")
	}

	httpClient := http.Client{}

	endpoint := fmt.Sprintf("%s/%s", c.Config.PublicApiBaseUrl, uri)

	httpRequest, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

func (c *Client) PublicApiCallWithCustomResult(uri string, result interface{}) error {
	resp, err := c.PublicApiCall(uri)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(*resp, result); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetServerTime() (*GetServerTimeResponseBody, error) {
	var result GetServerTimeResponseBody

	if err := c.PublicApiCallWithCustomResult("/api/server_time", &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPairs() (*[]Pair, error) {
	var result []Pair

	if err := c.PublicApiCallWithCustomResult("/api/pairs", &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPriceIncrements() (*GetPriceIncrementsResponseBody, error) {
	var result GetPriceIncrementsResponseBody

	if err := c.PublicApiCallWithCustomResult("/api/price_increments", &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSummaries() (*GetSummariesResponseBody, error) {
	var result GetSummariesResponseBody

	if err := c.PublicApiCallWithCustomResult("/api/summaries", &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTicker(pairId string) (*GetTickerResponseBody, error) {
	var result GetTickerResponseBody

	if err := c.PublicApiCallWithCustomResult(fmt.Sprintf("/api/ticker/%s", pairId), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTickerAll() (*GetTickerAllResponseBody, error) {
	var result GetTickerAllResponseBody

	if err := c.PublicApiCallWithCustomResult("/api/ticker_all", &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTrades(pairId string) (*[]Trade, error) {
	var result []Trade

	if err := c.PublicApiCallWithCustomResult(fmt.Sprintf("/api/trades/%s", pairId), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDepth(pairId string) (*GetDepthResponseBody, error) {
	var result GetDepthResponseBody

	if err := c.PublicApiCallWithCustomResult(fmt.Sprintf("/api/depth/%s", pairId), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetOHLCHistory(pairId, timeFrame string, from, to int64) (*[]OHLC, error) {
	var result []OHLC

	if err := c.PublicApiCallWithCustomResult(fmt.Sprintf("/tradingview/history_v2?symbol=%s&tf=%s&from=%d&to=%d", pairId, timeFrame, from, to), &result); err != nil {
		return nil, err
	}

	return &result, nil
}
