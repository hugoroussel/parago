package parago

import (
	"fmt"
)

const (
	API_URL          = "https://apiv4.paraswap.io/v2"
	NETWORK_ETHEREUM = 1
	NETWORK_ROPSTEN  = 3
	NETWORK_BINANCE  = 36
	NETWORK_POLYGON  = 137
	SELL             = "SELL"
	BUY              = "BUY"
	ONE_PERCENT      = 0.01
	TWO_PERCENT      = 0.01
	THREE_PERCENT    = 0.03
	FOUR_PERCENT     = 0.04
	FIVE_PERCENT     = 0.05
)

func (c *Client) GetTokenCall() string {
	return fmt.Sprintf("%v/tokens/%v", API_URL, c.Configuration.Network)
}

func (c *Client) GetRateCall(from string, to string, df string, dt string, amount int64, side string) string {
	const pricesURL = `%v/prices?from=%v&to=%v&amount=%v&fromDecimals=%v&toDecimals=%v&side=%v&network=%v`
	eurl := fmt.Sprintf(pricesURL, API_URL, from, to, amount, df, dt, side, c.Configuration.Network)
	return eurl
}

func (c *Client) BuildParametersCall() string {
	return fmt.Sprintf("%v/transactions/%v", API_URL, c.Configuration.Network)
}

func (c *Client) GetToken(symbol string) (*Token, error) {
	for _, v := range c.AllTokens.Tokens {
		if v.Symbol == symbol {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("no token found")
}
