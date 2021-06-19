package parago

import (
	"log"
)

// Simple test to show how the client works
func main() {

	log.Println("Welcome to parago")

	// Setup your address here
	c := NewConfiguration(NETWORK_POLYGON, "paraswap.io", "0xAccountAddress", ONE_PERCENT)

	client, err := NewClient(c)
	if err != nil {
		log.Println("error creating client", err)
		return
	}

	matic, err := client.GetToken("MATIC")
	if err != nil {
		log.Println("Wrong token symbol: ", err)
		return
	}

	usdc, err := client.GetToken("USDC")
	if err != nil {
		log.Println("Wrong token symbol: ", err)
		return
	}

	rate, err := client.GetRate(matic, usdc, 1e18, "SELL")
	if err != nil {
		log.Println("error getting rate", err)
		return
	}

	bp, err := client.BuildParameters(matic, usdc, rate, NULL_ADDRESS)
	if err != nil {
		log.Println("something went wrong", err)
		return
	}

	log.Println("Add this data to your transaction", bp.Data)
}
