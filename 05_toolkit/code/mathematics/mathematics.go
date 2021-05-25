package mathematics

import "fmt"

func print_current_number(num int) {
	fmt.Println("Current number: ", num)
}

func Add(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		
		total += nums[i]
		//print_current_number(total)
	}

	return total
}

func DoSomeOtherStuff(){
	fmt.Println("I am usefull boi.")
}