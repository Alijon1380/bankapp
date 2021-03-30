package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 1_000_0,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))
	// Output: 10000
}
