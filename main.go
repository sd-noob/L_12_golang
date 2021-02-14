package main

import (
	"fmt"
)

func main(){
	
// Simple loop
/*
	for i := 1; i<=9; i++{
				fmt.Println(i)	
	                     }
*/
// Range for loop
/*	
	students :=[]string{"Sourav","Das","Hriday"}
	for i, std := range students {
					fmt.Println(i,std)
				      }
*/

// While loop
	
/*	
	for true{
			fmt.Println("Hello")
	         }
*/
	
	i := 0
	for i < 10 {
			fmt.Println(i,"Hello")
			i++
		    }		
}