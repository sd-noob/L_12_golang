package main

import (
	"fmt"
)

func main(){
	
	fmt.Print("Enter your age:")
	var age int
	fmt.Scanf("%d",& age)
	//fmt.Println( age)
	
	switch age {
	case 2:
 		fmt.Println("Infant")
	case 3,4,5,6,7,8,9,10,11,12:
		fmt.Println("Children")
	case 13,14,15,16,17,18,19:
		fmt.Println("Teen age")
	default:
		fmt.Println("Adult")
	}
	
}