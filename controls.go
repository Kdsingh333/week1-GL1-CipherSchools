package main

import "fmt"

func main() {
	// data := 10
	// fmt.Println(data)

	// var input int
	// fmt.Scanln(&input)

	// if input%2 == 0{
	// 	fmt.Println("hey you are even")
	// } else {
	// 	fmt.Println("Hey you are odd")
	// }
	data := 100

	switch(data){
	case 10:
		data = 100
		fmt.Println(data);
	case 100:
		data = 103
		fmt.Println(data);
		fallthrough // execute the next case also
	case 1000:
		data = 104
		fmt.Println(data);
	case 100000:
		data = 105
		fmt.Println(data);
	default :
		data = 30
		fmt.Println(data);
	}
    

}