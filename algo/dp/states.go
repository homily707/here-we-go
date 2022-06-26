package dp

import "fmt"

func subset(state int) {
	//fmt.Printf("%b \n", state)
	mask := state
	for state > 0 {
		//fmt.Printf("%b \n", state)
		state = (state - 1) & mask
	}
}

func bitcount(state int) {
	for state > 0 {
		fmt.Printf("%b %b \n", state, lowbit(state))
		state -= lowbit(state)

	}
}

func lowbit(x int) int {
	return x & (-x)
}
