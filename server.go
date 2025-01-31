package main

import (
	"fmt"
	"sync"
	"time"
)

func makeTea(wg * sync.WaitGroup, name string) {
	defer wg.Done()
	// fmt.Println("Boiling water...")
	time.Sleep(5 * time.Second)
	// fmt.Println("Water Boiled")
	// fmt.Println("Adding Tea powder...")
	time.Sleep(2 * time.Second)
	// fmt.Println("Added Tea powder")
	// fmt.Println("Adding Milk and Sugar...")
	time.Sleep(10 * time.Second)
	// fmt.Println("Added Milk and Sugar")
	fmt.Printf("%80s %s✅\n", "Tea is ready to be served for", name)
}
type Order struct{
	beverage int 
	name string
}

func main() {
	var wg sync.WaitGroup
	openShop := true
	for openShop{
		fmt.Println("\nEnter your Name followed By 1 for Tea🍵")
		var order Order
		fmt.Scan(&order.beverage)
		fmt.Scan(&order.name)
		switch order.beverage{
		case 1:
			fmt.Printf("%80s %s✅\n", order.name, "ordered Tea🍵")
			wg.Add(1)
			go makeTea(&wg, order.name)
		case -1:
			if order.name == "Vishnu"{
				openShop = false
			}else{
				fmt.Println("You cant close the fucking shop")
			}
		}
	}
	wg.Wait()
	fmt.Println("SHOP IS CLOSED..")

}
