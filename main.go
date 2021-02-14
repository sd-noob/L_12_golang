package main

import (
	"fmt"
)

func main(){
	fmt.Print("Enter your age:")
	var age int
	fmt.Scanf("%d",& age)
	//fmt.Println( age)
	
	if age < 3 {
 		fmt.Println("Infant")
	}else if age > 2 && age < 12 {
		fmt.Println("Children")
	}else if age >= 12 && age <= 19 {
		fmt.Println("Teen age")
	}else{
		fmt.Println("Adult")
	}
	
}