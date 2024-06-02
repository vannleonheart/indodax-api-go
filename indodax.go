package indodax

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"github.com/vannleonheart/goutil"
	"strings"
)

func New(config Config) *Client {
	return &Client{Config: config}
}

func (c *Client) WithCredential(tradeApiKey, tradeApiSecret string) *Client {
	c.Credential = &Credential{
		TradeApiKey:    tradeApiKey,
		TradeApiSecret: tradeApiSecret,
	}

	return c
}

/*
 * Logging
 *
 * @param string level
 * @param interface{} data
 *
 * @return void
 */
func (c *Client) log(level string, data interface{}) {
	if c.Config.Log != nil && c.Config.Log.Enable {
		msg := map[string]interface{}{
			"level": level,
			"data":  data,
		}

		_ = goutil.WriteJsonToFile(msg, c.Config.Log.Path, c.Config.Log.Filename, c.Config.Log.Extension, c.Config.Log.Rotation)
	}
}

/*
 * Generate signature
 *
 * @param map[string]interface{} data
 *
 * @return *string
 * @return error
 */
func (c *Client) generateSign(data map[string]interface{}) (*string, error) {
	queryString, err := goutil.GenerateQueryString(data)
	if err != nil {
		return nil, err
	}

	if queryString == nil {
		return nil, errors.New("generated query string is nil")
	}

	if c.Credential == nil || len(c.Credential.TradeApiSecret) <= 0 {
		return nil, errors.New("invalid credential")
	}

	h := hmac.New(sha512.New, []byte(c.Credential.TradeApiSecret))

	h.Write([]byte(*queryString))

	signature := hex.EncodeToString(h.Sum(nil))

	return &signature, nil
}

/*
 * Adjust network name
 *
 * @param string currency
 * @param string network
 *
 * @return string
 */
func (c *Client) getNetworkName(currency, network string) string {
	network = strings.ToLower(network)

	switch network {
	case "bsc":
		network = "bep20"
	case "eth", "homestead":
		network = "erc20"
		if currency == "eth" {
			network = "eth"
		}
	case "matic", "polygon":
		network = "polygon"
	case "arbitrum":
		network = "arb"
	case "optimism":
		network = "op"
	}

	return network
}
