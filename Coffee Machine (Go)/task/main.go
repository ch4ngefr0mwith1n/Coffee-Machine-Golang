package main

import (
	"fmt"
	"math"
)

func main() {
	drink := Drink{
		water:       200,
		milk:        50,
		coffeeBeans: 15,
	}

	coffeeMachine := CoffeeMachine{}
	coffeeMachine.fillMachine()

	coffeeMachine.howManyCups(drink)
}

type Drink struct {
	water       int
	milk        int
	coffeeBeans int
}

type CoffeeMachine struct {
	availableWater       int
	availableMilk        int
	availableCoffeeBeans int
	availableCups        int
}

func (coffeeMachine *CoffeeMachine) fillMachine() {
	var water, milk, coffee int

	fmt.Println("Write how many ml of water the coffee machine has:")
	_, _ = fmt.Scan(&water)

	fmt.Println("Write how many ml of milk the coffee machine has:")
	_, _ = fmt.Scan(&milk)

	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	_, _ = fmt.Scan(&coffee)

	coffeeMachine.availableWater = water
	coffeeMachine.availableMilk = milk
	coffeeMachine.availableCoffeeBeans = coffee
}

func (coffeeMachine *CoffeeMachine) howManyCups(drink Drink) {
	fmt.Println("Write how many cups of coffee you will need:")

	var neededCups int
	_, _ = fmt.Scan(&neededCups)

	cups := int(
		math.Min(
			float64(coffeeMachine.availableWater/drink.water), math.Min(float64(coffeeMachine.availableMilk/drink.milk), float64(coffeeMachine.availableCoffeeBeans/drink.coffeeBeans))))

	if cups == neededCups {
		fmt.Println("Yes, I can make that amount of coffee")
	} else if cups > neededCups {
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", cups-neededCups)
	} else if cups < neededCups {
		fmt.Printf("No, I can make only %d cups of coffee", cups)
	}
}
