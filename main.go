package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите команду: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка чтения!")
			return
		}
		text := scanner.Text()

		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели!")
			return
		}
		cmd := fields[0]

		fmt.Println("Команда:", cmd)

		if cmd == "добавить" || cmd == "удалить" {
			str := " "
			for i := 1; i < len(fields); i++ {
				str += fields[i]

				if i != len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("Кажется, вы хотите " + cmd + str)
		} else if cmd == "help" {
			fmt.Println("Список доступных команд:")
			fmt.Println("")
			fmt.Println("добавить [то, что хотите добавить]")
			fmt.Println("добавляет что-то")
			fmt.Println("-------------------")
			fmt.Println("добавить [то, что хотите добавить]")
			fmt.Println("удаляет что-то")
			fmt.Println("-------------------")
			fmt.Println("help")
			fmt.Println("Выводит список доступных команд")
		} else if cmd == "выйти" {
			fmt.Println("До скорого!")
			break
		} else {
			fmt.Println("Вы ввели неизвестную команду!")
		}
	}
}
