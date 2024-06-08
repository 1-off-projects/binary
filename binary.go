package main

//dont touch anything it works
import (
	"fmt"
	"os"
	"strconv"
)

func binaryToDecimal(bin int) {
	binaryStr := strconv.FormatInt(int64(bin), 10)
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Print("value is: ")
	fmt.Println(decimal)
}

func decimalToBinary(dec int) string {
	return strconv.FormatInt(int64(dec), 2)
}
func main() {
	for {
		fmt.Println("1. to binary")
		fmt.Println("2. from binary")
		fmt.Println("3. quit")
		var choice int
		fmt.Println("enter 1-3: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("error")
			return
		}
		// to
		if choice == 1 {
			var dec int
			fmt.Println("you selected to binary")
			fmt.Println("decimal input")
			_, err = fmt.Scanln(&dec)
			if err != nil {
				fmt.Println("error")
				return
			}
			bin := decimalToBinary(dec)
			fmt.Print("value is: ")
			fmt.Println(bin)
		} else if choice == 2 {
			var bin int
			fmt.Println("you selected from binary")
			fmt.Println("binary input")
			_, err := fmt.Scanln(&bin)
			if err != nil {
				fmt.Println("error")
				return
			}
			(binaryToDecimal(bin))
		} else if choice == 3 {
			fmt.Println("exiting")
			os.Exit(0)
		} else {
			fmt.Println("invalid. goodbye")
			os.Exit(0)
		}
	}
}
