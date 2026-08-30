package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func name() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите имя игрока номер 1:")
	name1, _ := reader.ReadString('\n')
	fmt.Println("Введите имя игрока номер 2:")
	name2, _ := reader.ReadString('\n')
	fmt.Printf("Игрок 1: %s", name1)
	fmt.Printf("Игрок 2: %s", name2)
	First = strings.TrimSpace(name1)
	Second = strings.TrimSpace(name2)
}

var first int = 0
var second int = 0
var First string = "X"
var Second string = "O"
var win_First string = "Победил X"
var win_Second string = "Победил O"
var draw bool = true
var draws int = 0
var Game int = 0
var Move int = 1
var Cells = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func reset() {
	Cells = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	Move = 1
}

func color() {
	First = "\033[91mX\033[0m"
	Second = "\033[94mO\033[0m"
	win_First = "\033[32mПобедил X\033[0m"
	win_Second = "\033[32mПобедил O\033[0m"
}

func Winner() {
	if Cells[0] == First && Cells[1] == First && Cells[2] == First {
		fmt.Println(win_First)
		draw = false
		first++
		Game++
	} else if Cells[3] == First && Cells[4] == First && Cells[5] == First {
		fmt.Println(win_First)
		draw = false
		first++
		Game++
	} else if Cells[6] == First && Cells[7] == First && Cells[8] == First {
		fmt.Println(win_First)
		draw = false
		first++
		Game++
	} else if Cells[0] == First && Cells[3] == First && Cells[6] == First {
		fmt.Println(win_First)
		draw = false
		first++
		Game++
	} else if Cells[1] == First && Cells[4] == First && Cells[7] == First {
		fmt.Println(win_First)
		draw = false
		first++
		Game++
	} else if Cells[2] == First && Cells[5] == First && Cells[8] == First {
		fmt.Println(win_First)
		draw = false
		first++
		Game++
	} else if Cells[0] == First && Cells[4] == First && Cells[8] == First {
		fmt.Println(win_First)
		draw = false
		first++
		Game++
	} else if Cells[2] == First && Cells[4] == First && Cells[6] == First {
		fmt.Println(win_First)
		draw = false
		first++
		Game++
	} else if Cells[0] == Second && Cells[1] == Second && Cells[2] == Second {
		fmt.Println(win_Second)
		draw = false
		second++
		Game++
	} else if Cells[3] == Second && Cells[4] == Second && Cells[5] == Second {
		fmt.Println(win_Second)
		draw = false
		second++
		Game++
	} else if Cells[6] == Second && Cells[7] == Second && Cells[8] == Second {
		fmt.Println(win_Second)
		draw = false
		second++
		Game++
	} else if Cells[0] == Second && Cells[3] == Second && Cells[6] == Second {
		fmt.Println(win_Second)
		draw = false
		second++
		Game++
	} else if Cells[1] == Second && Cells[4] == Second && Cells[7] == Second {
		fmt.Println(win_Second)
		draw = false
		second++
		Game++
	} else if Cells[2] == Second && Cells[5] == Second && Cells[8] == Second {
		fmt.Println(win_Second)
		draw = false
		second++
		Game++
	} else if Cells[0] == Second && Cells[4] == Second && Cells[8] == Second {
		fmt.Println(win_Second)
		draw = false
		second++
		Game++
	} else if Cells[2] == Second && Cells[4] == Second && Cells[6] == Second {
		fmt.Println(win_Second)
		draw = false
		second++
		Game++
	}
	if draw == true && Move > 9 {
		fmt.Println("Ничья")
		draws++
		Game++
	}
}

func statistics() {
	fmt.Println("Статистика:")
	fmt.Printf("Всего игр: %d\n", Game)
	fmt.Printf("Победы %s: %d\n", First, first)
	fmt.Printf("Победы %s: %d\n", Second, second)
	fmt.Printf("Ничьи: %d\n", draws)
}

func verbose() {
	fmt.Println("Ход игры:")
	fmt.Printf("Ход %d: Игрок %s\n", Move, First)
	fmt.Printf("Ход %d: Игрок %s\n", Move, Second)
}

func help() {
	fmt.Println("Использование:")
	fmt.Println("--name : Ввести имена игроков")
	fmt.Println("--color : Включить цветной вывод")
	fmt.Println("--help : Показать справку")
}

func main() {
	for _, arg := range os.Args {
		if len(os.Args) > 1 && arg == "--name" {
			name()
		}
		if len(os.Args) > 1 && arg == "--color" {
			color()
		}
		if len(os.Args) > 1 && arg == "--help" {
			help()
			os.Exit(0)
		}
	}
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
			if Cells[0] == First || Cells[0] == Second {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[0] = First
				Move++
			} else {
				Cells[0] = Second
				Move++
			}
		case "2":
			if Cells[1] == First || Cells[1] == Second {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[1] = First
				Move++
			} else {
				Cells[1] = Second
				Move++
			}
		case "3":
			if Cells[2] == First || Cells[2] == Second {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[2] = First
				Move++
			} else {
				Cells[2] = Second
				Move++
			}
		case "4":
			if Cells[3] == First || Cells[3] == Second {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[3] = First
				Move++
			} else {
				Cells[3] = Second
				Move++
			}
		case "5":
			if Cells[4] == First || Cells[4] == Second {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[4] = First
				Move++
			} else {
				Cells[4] = Second
				Move++
			}
		case "6":
			if Cells[5] == First || Cells[5] == Second {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[5] = First
				Move++
			} else {
				Cells[5] = Second
				Move++
			}
		case "7":
			if Cells[6] == First || Cells[6] == Second {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[6] = First
				Move++
			} else {
				Cells[6] = Second
				Move++
			}
		case "8":
			if Cells[7] == First || Cells[7] == Second {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[7] = First
				Move++
			} else {
				Cells[7] = Second
				Move++
			}
		case "9":
			if Cells[8] == First || Cells[8] == Second {
				fmt.Println("Эта клетка уже занята")
				continue
			}
			if Move%2 == 1 {
				Cells[8] = First
				Move++
			} else {
				Cells[8] = Second
				Move++
			}
		default:
			fmt.Println("Неверный ввод. Введите число от 1 до 9.")
		}
		Winner()
		if Move > 9 {
			fmt.Println("Хочешь сыграть еще раз? (y/n)")
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				continue
			}
			switch strings.TrimSpace(input) {
			case "y":
				reset()
			case "n":
				statistics()
				os.Exit(0)
			default:
				fmt.Println("Неверный ввод. Введите 'y' или 'n'.")
			}
		}
	}
}
