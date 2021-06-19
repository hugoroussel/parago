package parago

// ExchangeRequest is the build parameters request call object
type ExchangeRequest struct {
	PriceRoute   interface{} `json:"priceRoute"`
	SrcToken     string      `json:"srcToken"`
	FromDecimals int         `json:"fromDecimals"`
	DestToken    string      `json:"destToken"`
	ToDecimals   int         `json:"toDecimals"`
	SrcAmount    string      `json:"srcAmount"`
	DestAmount   string      `json:"destAmount"`
	UserAddress  string      `json:"userAddress"`
	Referrer     string      `json:"referrer"`
	Receiver     string      `json:"receiver"`
}

type TokenResponse struct {
	Tokens *AllTokens
}

type AllTokens struct {
	Tokens []Token `json:"tokens"`
}

type Token struct {
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Img      string `json:"img"`
	Network  int    `json:"network"`
}

type PriceRoute struct {
	MultiPath bool `json:"multiPath"`
	BestRoute []struct {
		Exchange              string `json:"exchange"`
		SrcAmount             string `json:"srcAmount"`
		DestAmount            string `json:"destAmount"`
		Percent               string `json:"percent"`
		DestAmountFeeDeducted string `json:"destAmountFeeDeducted"`
	} `json:"bestRoute"`
	MultiRoute     []interface{} `json:"multiRoute"`
	BlockNumber    int           `json:"blockNumber"`
	DestAmount     string        `json:"destAmount"`
	SrcAmount      string        `json:"srcAmount"`
	AdapterVersion string        `json:"adapterVersion"`
	Others         []struct {
		Exchange string `json:"exchange"`
		Rate     string `json:"rate"`
		Unit     string `json:"unit"`
		Data     struct {
			Router    string   `json:"router"`
			Path      []string `json:"path"`
			Factory   string   `json:"factory"`
			InitCode  string   `json:"initCode"`
			TokenFrom string   `json:"tokenFrom"`
			TokenTo   string   `json:"tokenTo"`
			GasUSD    string   `json:"gasUSD"`
		} `json:"data"`
		RateFeeDeducted string `json:"rateFeeDeducted"`
		UnitFeeDeducted string `json:"unitFeeDeducted"`
	} `json:"others"`
	Side    string `json:"side"`
	Details struct {
		TokenFrom  string `json:"tokenFrom"`
		TokenTo    string `json:"tokenTo"`
		SrcAmount  string `json:"srcAmount"`
		DestAmount string `json:"destAmount"`
	} `json:"details"`
	BestRouteGas          string `json:"bestRouteGas"`
	BestRouteGasCostUSD   string `json:"bestRouteGasCostUSD"`
	ContractMethod        string `json:"contractMethod"`
	FromUSD               string `json:"fromUSD"`
	ToUSD                 string `json:"toUSD"`
	PriceWithSlippage     string `json:"priceWithSlippage"`
	Spender               string `json:"spender"`
	DestAmountFeeDeducted string `json:"destAmountFeeDeducted"`
	ToUSDFeeDeducted      string `json:"toUSDFeeDeducted"`
	MaxImpactReached      bool   `json:"maxImpactReached"`
	PriceID               string `json:"priceID"`
	Hmac                  string `json:"hmac"`
}
type Rate struct {
	PriceRoute struct {
		BestRoute []struct {
			Exchange   string `json:"exchange"`
			SrcAmount  string `json:"srcAmount"`
			DestAmount string `json:"destAmount"`
			Percent    string `json:"percent"`
			Data       struct {
				Router    string   `json:"router"`
				Path      []string `json:"path"`
				Factory   string   `json:"factory"`
				InitCode  string   `json:"initCode"`
				TokenFrom string   `json:"tokenFrom"`
				TokenTo   string   `json:"tokenTo"`
				GasUSD    string   `json:"gasUSD"`
			} `json:"data"`
			DestAmountFeeDeducted string `json:"destAmountFeeDeducted"`
		} `json:"bestRoute"`
		BlockNumber    int    `json:"blockNumber"`
		DestAmount     string `json:"destAmount"`
		SrcAmount      string `json:"srcAmount"`
		AdapterVersion string `json:"adapterVersion"`
		Others         []struct {
			Exchange string `json:"exchange"`
			Rate     string `json:"rate"`
			Unit     string `json:"unit"`
			Data     struct {
				Router    string   `json:"router"`
				Path      []string `json:"path"`
				Factory   string   `json:"factory"`
				InitCode  string   `json:"initCode"`
				TokenFrom string   `json:"tokenFrom"`
				TokenTo   string   `json:"tokenTo"`
				GasUSD    string   `json:"gasUSD"`
			} `json:"data"`
			RateFeeDeducted string `json:"rateFeeDeducted"`
			UnitFeeDeducted string `json:"unitFeeDeducted"`
		} `json:"others"`
		Side    string `json:"side"`
		Details struct {
			TokenFrom  string `json:"tokenFrom"`
			TokenTo    string `json:"tokenTo"`
			SrcAmount  string `json:"srcAmount"`
			DestAmount string `json:"destAmount"`
		} `json:"details"`
		BestRouteGas          string        `json:"bestRouteGas"`
		BestRouteGasCostUSD   string        `json:"bestRouteGasCostUSD"`
		ContractMethod        string        `json:"contractMethod"`
		FromUSD               string        `json:"fromUSD"`
		ToUSD                 string        `json:"toUSD"`
		PriceWithSlippage     string        `json:"priceWithSlippage"`
		Spender               string        `json:"spender"`
		DestAmountFeeDeducted string        `json:"destAmountFeeDeducted"`
		ToUSDFeeDeducted      string        `json:"toUSDFeeDeducted"`
		MultiRoute            []interface{} `json:"multiRoute"`
		MaxImpactReached      bool          `json:"maxImpactReached"`
		PriceID               string        `json:"priceID"`
		Hmac                  string        `json:"hmac"`
	} `json:"priceRoute"`
}

type BuildParamaters struct {
	From     string `json:"from"`
	To       string `json:"to"`
	ChainID  int    `json:"chainId"`
	Value    string `json:"value"`
	Data     string `json:"data"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
}
