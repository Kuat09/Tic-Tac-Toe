package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var x int = 0
var o int = 0
var draw int = 0
var Move int = 1
var Cells = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func Winner() {
	if Cells[0] == "X" && Cells[1] == "X" && Cells[2] == "X" {
		fmt.Println("Победил X")
		x++
		os.Exit(0)
	} else if Cells[3] == "X" && Cells[4] == "X" && Cells[5] == "X" {
		fmt.Println("Победил X")
		x++
		os.Exit(0)
	} else if Cells[6] == "X" && Cells[7] == "X" && Cells[8] == "X" {
		fmt.Println("Победил X")
		x++
		os.Exit(0)
	} else if Cells[0] == "X" && Cells[3] == "X" && Cells[6] == "X" {
		fmt.Println("Победил X")
		x++
		os.Exit(0)
	} else if Cells[1] == "X" && Cells[4] == "X" && Cells[7] == "X" {
		fmt.Println("Победил X")
		x++
		os.Exit(0)
	} else if Cells[2] == "X" && Cells[5] == "X" && Cells[8] == "X" {
		fmt.Println("Победил X")
		x++
		os.Exit(0)
	} else if Cells[0] == "X" && Cells[4] == "X" && Cells[8] == "X" {
		fmt.Println("Победил X")
		x++
		os.Exit(0)
	} else if Cells[2] == "X" && Cells[4] == "X" && Cells[6] == "X" {
		fmt.Println("Победил X")
		x++
		os.Exit(0)
	} else if Cells[0] == "O" && Cells[1] == "O" && Cells[2] == "O" {
		fmt.Println("Победил O")
		o++
		os.Exit(0)
	} else if Cells[3] == "O" && Cells[4] == "O" && Cells[5] == "O" {
		fmt.Println("Победил O")
		o++
		os.Exit(0)
	} else if Cells[6] == "O" && Cells[7] == "O" && Cells[8] == "O" {
		fmt.Println("Победил O")
		o++
		os.Exit(0)
	} else if Cells[0] == "O" && Cells[3] == "O" && Cells[6] == "O" {
		fmt.Println("Победил O")
		o++
		os.Exit(0)
	} else if Cells[1] == "O" && Cells[4] == "O" && Cells[7] == "O" {
		fmt.Println("Победил O")
		o++
		os.Exit(0)
	} else if Cells[2] == "O" && Cells[5] == "O" && Cells[8] == "O" {
		fmt.Println("Победил O")
		o++
		os.Exit(0)
	} else if Cells[0] == "O" && Cells[4] == "O" && Cells[8] == "O" {
		fmt.Println("Победил O")
		o++
		os.Exit(0)
	} else if Cells[2] == "O" && Cells[4] == "O" && Cells[6] == "O" {
		fmt.Println("Победил O")
		o++
		os.Exit(0)
	} else {
		fmt.Println("Ничья")
		draw++
		os.Exit(0)
	}
}

func statistics() {
	fmt.Println("Статистика:")
	fmt.Printf("Победы X: %d\n", x)
	fmt.Printf("Победы O: %d\n", o)
	fmt.Printf("Ничьи: %d\n", draw)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(" %s | %s | %s", Cells[0], Cells[1], Cells[2])
		fmt.Println()
		fmt.Println("---+---+---")
		fmt.Printf(" %s | %s | %s", Cells[3], Cells[4], Cells[5])
		fmt.Println()
		fmt.Println("---+---+---")
		fmt.Printf(" %s | %s | %s", Cells[6], Cells[7], Cells[8])
		fmt.Println()
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		choice := strings.TrimSpace(input)
		switch choice {
		case "1":
			if Cells[0] == "X" || Cells[0] == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[0] = "X"
				Move++
			} else {
				Cells[0] = "O"
				Move++
			}
		case "2":
			if Cells[1] == "X" || Cells[1] == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[1] = "X"
				Move++
			} else {
				Cells[1] = "O"
				Move++
			}
		case "3":
			if Cells[2] == "X" || Cells[2] == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[2] = "X"
				Move++
			} else {
				Cells[2] = "O"
				Move++
			}
		case "4":
			if Cells[3] == "X" || Cells[3] == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[3] = "X"
				Move++
			} else {
				Cells[3] = "O"
				Move++
			}
		case "5":
			if Cells[4] == "X" || Cells[4] == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[4] = "X"
				Move++
			} else {
				Cells[4] = "O"
				Move++
			}
		case "6":
			if Cells[5] == "X" || Cells[5] == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[5] = "X"
				Move++
			} else {
				Cells[5] = "O"
				Move++
			}
		case "7":
			if Cells[6] == "X" || Cells[6] == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[6] = "X"
				Move++
			} else {
				Cells[6] = "O"
				Move++
			}
		case "8":
			if Cells[7] == "X" || Cells[7] == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[7] = "X"
				Move++
			} else {
				Cells[7] = "O"
				Move++
			}
		case "9":
			if Cells[8] == "X" || Cells[8] == "O" {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[8] = "X"
				Move++
			} else {
				Cells[8] = "O"
				Move++
			}
		default:
			fmt.Println("Неверный ввод. Введите число от 1 до 9.")
		}
		Winner()
	}
}
