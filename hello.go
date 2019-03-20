package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a positive integer:  ")
	numInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	numInput = strings.Replace(numInput, "\n", "", -1)
	num, err := strconv.ParseFloat(numInput, 64)
	if err != nil {
		log.Fatal(err)
	}
	squareRoot, err := squareRoot(num)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Square root of ", numInput, ": ", squareRoot)

	fmt.Println("Sales by month for 2018")
	salesByMonth()

	fmt.Println("Searching for hello world from randomly selected parts of the same text...")
	helloWorld()
}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("can't take square root of negative number")
	}
	return math.Sqrt(x), nil
}

func salesByMonth() {
	months := [12]string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	salesByMonth := [12]float64{
		1284.20, 1710.26, 2245.97, 3032.40, 2956.43, 3215.16,
		3165.98, 3420.10, 3687.61, 3678.73, 3712.10, 3799.35,
	}
	for i := 0; i < len(months); i++ {
		fmt.Println(months[i], salesByMonth[i])
	}
}

func helloWorld() {
	t := "Hello World!"
	s := []rune(t)
	for {
		rand.Shuffle(len(s), func(i int, j int) {
			s[i], s[j] = s[j], s[i]
		})
		time.Sleep(1 * time.Second)
		fmt.Println(string(s))
		if string(s) == string(t) {
			break
		}
	}
}
