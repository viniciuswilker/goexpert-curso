package main

import "fmt"

func main(){

	fmt.Println("Primeiro")
	defer fmt.Println("Segundo")
	fmt.Println("Terceiro")
}