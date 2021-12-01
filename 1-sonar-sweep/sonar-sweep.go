package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
	file,err := os.Open("input.data")
	check(err)

	defer file.Close()

	//simpleIncrements(file)
	slidingWindowOfThree(file)

}

func slidingWindowOfThree(file *os.File){
	scanner := bufio.NewScanner(file)

	var count, increased int
	var window [3]int

	for scanner.Scan(){
		current, err := strconv.Atoi(scanner.Text())
		check(err)

		currentSum := sumArray(window)
		window[count % 3] = current

		count++
		if  count > 3 && currentSum < sumArray(window){
			increased++
		}
	}

	fmt.Printf("Window of 3 increased %v times\n", increased)
}

func simpleIncrements(file *os.File){
	scanner := bufio.NewScanner(file)

	var count, previous int
	started := false
	for scanner.Scan(){

		current, err := strconv.Atoi(scanner.Text())
		check(err)

		if started && current > previous {
			count++
		}

		previous = current
		started = true
	}

	fmt.Printf("Increased %v times\n", count)
}
func sumArray(arr [3] int) int{
	var sum int
	for i:=0; i < len(arr) ; i++{
		sum += arr[i]
	}
	return sum
}
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}


