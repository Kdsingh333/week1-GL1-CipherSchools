package main

import "fmt"

func main() {
	number := 10
	fmt.Println(number)
	var getInt func(int) int
	getInt = func(x int ) int{
		fmt.Println("In a 1st function")
		return 20 + x
	}
 
	g(getInt,8)
	//  getInt = func (x int) int{
	// 	fmt.Println("In a function")
	// 	return 10 + x
	// }

	j := func(x int) int{
		fmt.Println("In a last function")
		return 20 + x
	}(10)
   
	
	 fmt.Println(j)
}

func g(getInt func(int) int,u int){
	getInt(78)
}

// function is a first class member in golang
