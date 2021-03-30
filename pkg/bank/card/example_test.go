package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
    fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_0,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
	  {
		Balance: 1_000_0,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
	  {
		Balance: 2_000_0,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
	  {
		Balance: 1_000_0,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
	  {
		Balance: -1_000_0,
			Active:  true,
		},
	}))

	// Output:
	// 10000
	// 10000
	// 20000
	// 0
	// 0
}