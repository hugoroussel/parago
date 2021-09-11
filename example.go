package parago

import (
	"log"
	"math/big"
)

// Simple test to show how the client works
func main() {

	log.Println("Welcome to parago")

	// Setup your address here
	c := NewConfiguration(NETWORK_ROPSTEN, "paraswap.io", "0x047E7375215af92E9b0Af86e47835Fe181f3C8af", ONE_PERCENT)

	client, err := NewClient(c)
	if err != nil {
		log.Println("error creating client", err)
		return
	}

	weth, err := client.GetTokenWithSymbol("WETH")
	if err != nil {
		log.Println("Wrong token symbol: ", err)
		return
	}

	dai, err := client.GetTokenWithSymbol("DAI")
	if err != nil {
		log.Println("Wrong token symbol: ", err)
		return
	}

	log.Println(weth, dai)

	rate, err := client.GetRate(weth, dai, big.NewInt(1e16), "SELL")
	if err != nil {
		log.Println("error getting rate", err)
		return
	}

	log.Println(rate)

	bp, err := client.BuildParameters(weth, dai, rate, NULL_ADDRESS)
	if err != nil {
		log.Println("something went wrong", err)
		return
	}

	log.Println("Add this data to your transaction", bp.Data)

}
