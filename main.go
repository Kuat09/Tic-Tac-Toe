package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var a int = 1
var a1 string = "1"
var a2 string = "2"
var a3 string = "3"
var a4 string = "4"
var a5 string = "5"
var a6 string = "6"
var a7 string = "7"
var a8 string = "8"
var a9 string = "9"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(" %s | %s | %s", a1, a2, a3)
		fmt.Println()
		fmt.Println("---+---+---")
		fmt.Printf(" %s | %s | %s", a4, a5, a6)
		fmt.Println()
		fmt.Println("---+---+---")
		fmt.Printf(" %s | %s | %s", a7, a8, a9)
		fmt.Println()
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		choice := strings.TrimSpace(input)
		switch choice {
		case "1":
			if a1 == "X" || a1 == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if a%2 == 1 {
				a1 = "X"
				a++
			} else {
				a1 = "O"
				a++
			}
		case "2":
			if a2 == "X" || a2 == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if a%2 == 1 {
				a2 = "X"
				a++
			} else {
				a2 = "O"
				a++
			}
		case "3":
			if a3 == "X" || a3 == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if a%2 == 1 {
				a3 = "X"
				a++
			} else {
				a3 = "O"
				a++
			}
		case "4":
			if a4 == "X" || a4 == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if a%2 == 1 {
				a4 = "X"
				a++
			} else {
				a4 = "O"
				a++
			}
		case "5":
			if a5 == "X" || a5 == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if a%2 == 1 {
				a5 = "X"
				a++
			} else {
				a5 = "O"
				a++
			}
		case "6":
			if a6 == "X" || a6 == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if a%2 == 1 {
				a6 = "X"
				a++
			} else {
				a6 = "O"
				a++
			}
		case "7":
			if a7 == "X" || a7 == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if a%2 == 1 {
				a7 = "X"
				a++
			} else {
				a7 = "O"
				a++
			}
		case "8":
			if a8 == "X" || a8 == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if a%2 == 1 {
				a8 = "X"
				a++
			} else {
				a8 = "O"
				a++
			}
		case "9":
			if a9 == "X" || a9 == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if a%2 == 1 {
				a9 = "X"
				a++
			} else {
				a9 = "O"
				a++
			}
		default:
			return
		}
	}
}
