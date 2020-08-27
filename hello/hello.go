package hello

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
	squareRoot, err := SquareRoot(num)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Square root of ", numInput, ": ", squareRoot)

	salesByMonth := SalesByMonth()
	fmt.Println("Sales by month for 2018")
	for month, sales := range salesByMonth {
		fmt.Println(month, sales)
	}

	fmt.Println("Searching for hello world from randomly selected parts of the same text...")
	HiWorld()
}

func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("can't take square root of negative number")
	}
	return math.Sqrt(x), nil
}

func SalesByMonth() map[string]float64 {
	salesByMonth := map[string]float64{}
	salesByMonth["Jan"] = 1284.20
	salesByMonth["Feb"] = 1710.26
	salesByMonth["Mar"] = 2245.97
	salesByMonth["Apr"] = 3032.40
	salesByMonth["May"] = 2956.43
	salesByMonth["Jun"] = 3215.16
	salesByMonth["Jul"] = 3165.98
	salesByMonth["Aug"] = 3420.10
	salesByMonth["Sep"] = 3687.61
	salesByMonth["Oct"] = 3678.73
	salesByMonth["Nov"] = 3712.10
	salesByMonth["Dec"] = 3799.35

	return salesByMonth
}

func HiWorld() {
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
