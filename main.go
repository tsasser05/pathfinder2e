package main

import (
	"fmt"
	"github.com/tsasser05/dieroll"
	"github.com/tsasser05/pathfinder2e/abilities"
	"os"
	"strconv"
)

func main() {

	if argCount != 2 {
		fmt.Println("Incorrect number of args.  Call this program with max, num.")
		fmt.Println("  max = max of the die.  4, 6, 8, 10, 12, 20 or 100.")
		fmt.Println("  num = number of dice to roll.")
		fmt.Println("go run main.go 6 3")
		os.Exit(3)
	} // if

	max_str := os.Args[1]
	num_str := os.Args[2]

	max, _ := strconv.Atoi(max_str)
	num, _ := strconv.Atoi(num_str)

	argCount := len(os.Args[1:])

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Test 1: dieroll.Roll()")
	fmt.Println("-------------------------------------------------------")
	roll1 := dieroll.Roll(max, 1)
	fmt.Println("The roll1  was:  ", roll1, "\n")

	roll2 := dieroll.Roll(max, 2)
	fmt.Println("The roll2 total for two rolls with max =", max, " was:  ", roll2, "\n")

	roll3 := dieroll.Roll(max, num)
	fmt.Println("max = ", max, "num = ", num, "\tSum of rolls:  ", roll3)

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Test 2: dieroll.RollStat()")
	fmt.Println("   Tests the rolling of stats using 4d6.")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("The stat rolled was:  ", dieroll.RollStat())

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Test 3: dieroll.RandomKey() DOES NOTHING NOW")
	fmt.Println("-------------------------------------------------------")

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("Test 5: Test Abilities")
	fmt.Println("-------------------------------------------------------")

	abil := abilities.Abilities{}
	abil.Roll()
	abil.Display()

	fmt.Println("\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("End testing")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("\n")

} // main()
