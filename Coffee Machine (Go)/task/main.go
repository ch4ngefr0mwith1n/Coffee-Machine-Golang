package main

import (
	"fmt"
	"os"
)

func main() {
	coffeeMachine := CoffeeMachine{}
	coffeeMachine.fillMachine(400, 540, 120, 9, 550)

	for {
		coffeeMachine.actionMenu()
	}
}

type Drink struct {
	water       int
	milk        int
	coffeeBeans int
	price       int
}

type CoffeeMachine struct {
	availableWater       int
	availableMilk        int
	availableCoffeeBeans int
	availableCups        int
	availableMoney       int
}

func (coffeeMachine *CoffeeMachine) fillMachine(water int, milk int, coffeeBeans int, disposableCups int, money int) {
	coffeeMachine.availableWater = water
	coffeeMachine.availableMilk = milk
	coffeeMachine.availableCoffeeBeans = coffeeBeans
	coffeeMachine.availableCups = disposableCups
	coffeeMachine.availableMoney = money
}

func (coffeeMachine *CoffeeMachine) actionMenu() {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	var action string
	fmt.Scan(&action)

	switch action {
	case "buy":
		coffeeMachine.buyCoffee()
		break
	case "fill":
		coffeeMachine.fillMenu()
		break
	case "take":
		coffeeMachine.takeMoney()
		break
	case "remaining":
		coffeeMachine.showState()
	case "exit":
		os.Exit(0)
	}
}

func (coffeeMachine *CoffeeMachine) showState() {
	fmt.Printf("\nThe coffee machine has:\n%d ml of water\n%d ml of milk\n%d g of coffee beans\n%d disposable cups\n$%d of money\n\n", coffeeMachine.availableWater, coffeeMachine.availableMilk, coffeeMachine.availableCoffeeBeans, coffeeMachine.availableCups, coffeeMachine.availableMoney)
}

func (coffeeMachine *CoffeeMachine) takeMoney() {
	fmt.Printf("I gave you $%d\n", coffeeMachine.availableMoney)
	coffeeMachine.availableMoney = 0
}

func (coffeeMachine *CoffeeMachine) buyCoffee() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var option string
	fmt.Scan(&option)

	switch option {
	case "1":
		espresso := Drink{
			water:       250,
			milk:        0,
			coffeeBeans: 16,
			price:       4,
		}

		coffeeMachine.makeTransaction(espresso)
		break
	case "2":
		latte := Drink{
			water:       350,
			milk:        75,
			coffeeBeans: 20,
			price:       7,
		}

		coffeeMachine.makeTransaction(latte)
		break
	case "3":
		cappuccino := Drink{
			water:       200,
			milk:        100,
			coffeeBeans: 12,
			price:       6,
		}

		coffeeMachine.makeTransaction(cappuccino)
		break
	case "back":
		return
	}
}

func (coffeeMachine *CoffeeMachine) fillMenu() {
	var water, milk, coffee, disposableCups int

	fmt.Println("Write how many ml of water you want to add:")
	_, _ = fmt.Scan(&water)

	fmt.Println("Write how many ml of milk you want to add:")
	_, _ = fmt.Scan(&milk)

	fmt.Println("Write how many grams of coffee beans you want to add:")
	_, _ = fmt.Scan(&coffee)

	fmt.Println("Write how many disposable cups you want to add:")
	_, _ = fmt.Scan(&disposableCups)

	coffeeMachine.availableWater += water
	coffeeMachine.availableMilk += milk
	coffeeMachine.availableCoffeeBeans += coffee
	coffeeMachine.availableCups += disposableCups
}

func (coffeeMachine *CoffeeMachine) checkAvailability(drink Drink) bool {
	if coffeeMachine.availableWater < drink.water {
		fmt.Printf("Sorry, not enough water!\n\n")
		return false
	} else if coffeeMachine.availableMilk < drink.milk {
		fmt.Printf("Sorry, not enough milk!\n\n")
		return false
	} else if coffeeMachine.availableCoffeeBeans < drink.coffeeBeans {
		fmt.Printf("Sorry, not enough coffee beans!\n\n")
		return false
	} else if coffeeMachine.availableCups < 1 {
		fmt.Printf("Sorry, not enough cups!\n\n")
		return false
	}
	return true
}

func (coffeeMachine *CoffeeMachine) makeTransaction(drink Drink) {
	if !coffeeMachine.checkAvailability(drink) {
		return
	}

	coffeeMachine.availableWater -= drink.water
	coffeeMachine.availableMilk -= drink.milk
	coffeeMachine.availableCoffeeBeans -= drink.coffeeBeans
	coffeeMachine.availableCups -= 1

	coffeeMachine.availableMoney += drink.price
	fmt.Printf("I have enough resources, making you a coffee!\n\n")
}
