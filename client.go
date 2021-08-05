package parago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

const (
	NULL_ADDRESS = "0x0000000000000000000000000000000000000000"
)

// GetTokens returns all tokens supported on a given network
func (c *Client) GetTokens() (*AllTokens, error) {
	resp, err := http.Get(c.GetTokenCall())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	at := &AllTokens{}
	err = json.Unmarshal(data, at)
	if err != nil {
		return nil, err
	}
	return at, nil
}

// GetRate
// returns the rate object
// parameters
// two tokens object,
// the amount desired
// "SELL" or "BUY" side
func (c *Client) GetRate(tokenA *Token, tokenB *Token, amount *big.Int, side string) (*Rate, error) {
	return c.getRate(tokenA.Address, tokenB.Address, strconv.Itoa(tokenA.Decimals), strconv.Itoa(tokenB.Decimals), amount, side)
}

// getRate call the rate api
func (c *Client) getRate(from string, to string, df string, dt string, amount *big.Int, side string) (*Rate, error) {
	resp, err := http.Get(c.GetRateCall(from, to, df, dt, amount, side))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errorBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("api returned a non 200 code : %v. Error : %v", resp.StatusCode, string(errorBody))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &Rate{}
	err = json.Unmarshal(data, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// BuildParameters
// returns the BuildParameter object containing the transaction data
// Parameters
// two token objects
// the rate object from the GetRate function
// a receiver if different from the transaction originator
func (c *Client) BuildParameters(tokenA *Token, tokenB *Token, rate *Rate, receiver string) (*BuildParamaters, error) {

	pr := rate.PriceRoute

	multiplier := big.NewFloat(1 - c.Configuration.Slippage)

	// Convert the result into a big float
	da, check := big.NewFloat(0).SetString(pr.BestRoute[0].DestAmount)
	if !check {
		return nil, fmt.Errorf("error setting destination amount")
	}

	// Multiply by the slippage and convert to big int
	finalres, acc := da.Mul(da, multiplier).Int(big.NewInt(0))
	if acc != 0 {
		log.Println("non exact conversion", finalres.String(), acc)
	}

	er := &ExchangeRequest{
		PriceRoute:   rate.PriceRoute,
		SrcToken:     tokenA.Address,
		FromDecimals: tokenA.Decimals,
		DestToken:    tokenB.Address,
		ToDecimals:   tokenB.Decimals,
		SrcAmount:    pr.BestRoute[0].SrcAmount,
		DestAmount:   finalres.String(),
		UserAddress:  c.Configuration.UserAddress,
		Referrer:     c.Configuration.Referrer,
		Receiver:     receiver,
	}

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(er)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.BuildParametersCall(), "application/json", payloadBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errorBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("api returned a non 200 code : %v. Error : %v", resp.StatusCode, string(errorBody))
	}

	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bp := &BuildParamaters{}
	err = json.Unmarshal(dat, bp)
	if err != nil {
		return nil, err
	}

	return bp, nil
}
