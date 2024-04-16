package indodax

func New(config *Config) *Client {
	return &Client{Config: config}
}

func (c *Client) WithCredential(tradeApiKey, tradeApiSecret string) *Client {
	c.Credential = &Credential{
		TradeApiKey:    tradeApiKey,
		TradeApiSecret: tradeApiSecret,
	}

	return c
}
