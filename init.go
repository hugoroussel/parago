package parago

// Configuration is the configuration object of the parago client
type Configuration struct {
	Network     int
	Referrer    string
	UserAddress string
	Slippage    float64
}

// Client contains the configuration of the parago client as well as all the tokens supported on a given network
type Client struct {
	Configuration *Configuration
	AllTokens     *AllTokens
}

// NewConfiguration takes as parameter
// The Network : 1: Ethereum mainnet, 56: Binance Smart Chain, 137: Polygon, 3: Ropsten
// The Referrer : defaults to "paraswap.io"
// The User address : the address of the account that will broadcast the transaction
// The slippage : 0.01 for 1%, 0.05 for 5% etc..
// returns a configuration object containing the network, the referrer and the user address
func NewConfiguration(network int, referrer string, userAddress string, slippage float64) *Configuration {
	if referrer == "" {
		referrer = "paraswap.io"
	}
	return &Configuration{
		Network:     network,
		Referrer:    referrer,
		UserAddress: userAddress,
		Slippage:    slippage,
	}
}

// NewClient returns a new parago client with all the tokens corresponding to the configuration network
func NewClient(c *Configuration) (*Client, error) {
	client := &Client{Configuration: c}
	at, err := client.GetTokens()
	if err != nil {
		return nil, err
	}
	client.AllTokens = at
	return client, nil
}
