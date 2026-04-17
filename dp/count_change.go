package dp

import "fmt"

// coinType starts from 1 to 5, ignore coinType 0.
var firstNomination = []int{0, 1, 5, 10, 25, 50}

// CountChange count change in recursive way
func CountChange(amount int, coinType int) int {

	// one kind of change
	if amount == 0 {
		return 1
	}

	// can't be changed
	if amount < 0 || coinType == 0 {
		return 0
	}

	return CountChange(amount, coinType-1) + CountChange(amount-firstNomination[coinType], coinType)
}

// CountChange2 count change in iterative way
func CountChange2(amount int, coinType int) int {

	// state: amount, coinType
	// let state[coinType][amount] to accumulate change count
	// then state[coinType][amount] = state[cointType-1][amount] + state[coinType][amount - firstNomination[coinType]]

	// init, all zeros
	state := make([][]int, coinType+1)
	for i := 0; i <= coinType; i++ {
		state[i] = make([]int, amount+1)

		// first path: one coin
		state[i][firstNomination[i]] = 1
	}

	fmt.Println("-------------------------")
	for i := 1; i <= coinType; i++ {
		for j := 1; j <= amount; j++ {
			remainAmount := max(0, j-firstNomination[i])
			state[i][j] = state[i][j] + state[i-1][j] + state[i][remainAmount]
		}
	}

	// answer: state[coinType][amount]
	return state[coinType][amount]
}
