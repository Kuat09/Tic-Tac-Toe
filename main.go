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
	if strings.TrimSpace(name1) == "" {
		fmt.Println("Имя игрока номер 1 не может быть пустым. Пожалуйста, введите имя.")
		name()
	}
	if len([]rune(strings.TrimSpace(name1))) > 1 {
		fmt.Println("Имя игрока номер 1 слишком длинное. Пожалуйста, введите имя не более 1 символа.")
		name()
	}
	fmt.Println("Введите имя игрока номер 2:")
	name2, _ := reader.ReadString('\n')
	if strings.TrimSpace(name2) == "" {
		fmt.Println("Имя игрока номер 2 не может быть пустым. Пожалуйста, введите имя.")
		name()
	}
	if len([]rune(strings.TrimSpace(name2))) > 1 {
		fmt.Println("Имя игрока номер 2 слишком длинное. Пожалуйста, введите имя не более 1 символа.")
		name()
	}
	if strings.TrimSpace(name1) == strings.TrimSpace(name2) {
		fmt.Println("Имена игроков не могут быть одинаковыми. Пожалуйста, введите разные имена.")
		name()
	}
	fmt.Printf("Игрок 1: %s", name1)
	fmt.Printf("Игрок 2: %s", name2)
	First = strings.TrimSpace(name1)
	Second = strings.TrimSpace(name2)
}

var first int = 0
var second int = 0
var First string = "X"
var Second string = "O"
var win_First string = "Победил 1 игрок"
var win_Second string = "Победил 2 игрок"
var end bool = false
var draw bool = true
var draws int = 0
var Game int = 0
var Move int = 1
var Verbose bool = false
var Cells = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func reset() {
	Cells = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	Move = 1
}

func color() {
	First = "\033[91m" + First + "\033[0m"
	Second = "\033[94m" + Second + "\033[0m"
	win_First = "\033[32mПобедил " + First + "\033[0m"
	win_Second = "\033[32mПобедил " + Second + "\033[0m"
}

func Winner() {
	if Cells[0] == First && Cells[1] == First && Cells[2] == First {
		fmt.Println(win_First)
		draw = false
		end = true
		first++
		Game++
	} else if Cells[3] == First && Cells[4] == First && Cells[5] == First {
		fmt.Println(win_First)
		draw = false
		end = true
		first++
		Game++
	} else if Cells[6] == First && Cells[7] == First && Cells[8] == First {
		fmt.Println(win_First)
		draw = false
		end = true
		first++
		Game++
	} else if Cells[0] == First && Cells[3] == First && Cells[6] == First {
		fmt.Println(win_First)
		draw = false
		end = true
		first++
		Game++
	} else if Cells[1] == First && Cells[4] == First && Cells[7] == First {
		fmt.Println(win_First)
		draw = false
		end = true
		first++
		Game++
	} else if Cells[2] == First && Cells[5] == First && Cells[8] == First {
		fmt.Println(win_First)
		draw = false
		end = true
		first++
		Game++
	} else if Cells[0] == First && Cells[4] == First && Cells[8] == First {
		fmt.Println(win_First)
		draw = false
		end = true
		first++
		Game++
	} else if Cells[2] == First && Cells[4] == First && Cells[6] == First {
		fmt.Println(win_First)
		draw = false
		end = true
		first++
		Game++
	} else if Cells[0] == Second && Cells[1] == Second && Cells[2] == Second {
		fmt.Println(win_Second)
		draw = false
		end = true
		second++
		Game++
	} else if Cells[3] == Second && Cells[4] == Second && Cells[5] == Second {
		fmt.Println(win_Second)
		draw = false
		end = true
		second++
		Game++
	} else if Cells[6] == Second && Cells[7] == Second && Cells[8] == Second {
		fmt.Println(win_Second)
		draw = false
		end = true
		second++
		Game++
	} else if Cells[0] == Second && Cells[3] == Second && Cells[6] == Second {
		fmt.Println(win_Second)
		draw = false
		end = true
		second++
		Game++
	} else if Cells[1] == Second && Cells[4] == Second && Cells[7] == Second {
		fmt.Println(win_Second)
		draw = false
		end = true
		second++
		Game++
	} else if Cells[2] == Second && Cells[5] == Second && Cells[8] == Second {
		fmt.Println(win_Second)
		draw = false
		end = true
		second++
		Game++
	} else if Cells[0] == Second && Cells[4] == Second && Cells[8] == Second {
		fmt.Println(win_Second)
		draw = false
		end = true
		second++
		Game++
	} else if Cells[2] == Second && Cells[4] == Second && Cells[6] == Second {
		fmt.Println(win_Second)
		draw = false
		end = true
		second++
		Game++
	}
	if draw == true && Move > 9 {
		fmt.Println("Ничья")
		end = true
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
	if Verbose == true {
		verbose()
	}
}

func verbose() {
	if end == true {
		fmt.Println("Ходов в игре сделано:", Move-1)
		fmt.Println("Процент победы игрока номер 1:", float64(first)/float64(Game)*100, "%")
		fmt.Println("Процент победы игрока номер 2:", float64(second)/float64(Game)*100, "%")
	}
}

func FirstPlayer() {
	fmt.Println("Выберите, кто будет ходить первым:")
	fmt.Println("1. Игрок 1")
	fmt.Println("2. Игрок 2")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	switch strings.TrimSpace(input) {
	case "1":
		fmt.Printf("Первым ходит игрок %s\n", First)
	case "2":
		fmt.Printf("Первым ходит игрок %s\n", Second)
		First, Second = Second, First
	default:
		fmt.Println("Неверный ввод. Игрок 1 будет ходить первым по умолчанию.")
	}
}

func help() {
	fmt.Println("Использование:")
	fmt.Println("--name : Ввести имена игроков")
	fmt.Println("--color : Включить цветной вывод")
	fmt.Println("--help : Показать справку")
	fmt.Println("--verbose : Включить подробный вывод (не реализовано)")
	fmt.Println("--first : Выбрать, кто будет ходить первым")
}

func main() {
	for _, arg := range os.Args {
		if len(os.Args) > 1 && arg == "--name" {
			name()
		}
		if len(os.Args) > 1 && arg == "--color" {
			color()
		}
		if len(os.Args) > 1 && arg == "--verbose" {
			Verbose = true
		}
		if len(os.Args) > 1 && arg == "--first" {
			FirstPlayer()
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
		if end == true {
			if Verbose == true {
				verbose()
			}
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
