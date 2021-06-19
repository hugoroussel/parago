package parago

import (
	"log"
)

// Simple test to show how the client works
func main() {

	log.Println("Welcome to parago")

	c := NewConfiguration(137, "paraswap.io", "", 0.005)

	client, err := NewClient(c)
	if err != nil {
		log.Println("error creating client", err)
		return
	}

	matic, _ := client.GetToken("MATIC")
	link, _ := client.GetToken("USDC")

	rate, err := client.GetRate(matic, link, 1e18, "SELL")
	if err != nil {
		log.Println("error getting rate", err)
		return
	}

	bp, err := client.BuildParameters(matic, link, rate, NULL_ADDRESS)
	if err != nil {
		log.Println("something went wrong", err)
		return
	}

	log.Println("Add this data to your transaction", bp.Data)
}
