package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите команду: ")

	if ok := scanner.Scan(); !ok {
		fmt.Println("Ошибка ввода!")
		return
	}

	text := scanner.Text()
	fmt.Println("")
	fields := strings.Fields(text)

	if len(fields) == 0 {
		fmt.Println("Вы ничего не ввели")
		return
	}

	cmd := fields[0]

	if cmd == "добавить" || cmd == "удалить" {
		str := ""
		for i := 1; i < len(fields); i++ {
			str += fields[i]
			if i != len(fields)-1 {
				str += " "
			}
		}
		fmt.Println("Вы хотите " + cmd + " " + str)
	} else if cmd == "help" {
		fmt.Println("Команда 'добавить' (то, что вы хотите добавить)")
		fmt.Println("Добавляет то, что вы хотите добавить")
		fmt.Println("---")
		fmt.Println("Команда 'удалить' (то, что вы хотите удалить)")
		fmt.Println("Удаляет то, что вы хотите удалить")
		fmt.Println("---")
		fmt.Println("Команда 'help'")
		fmt.Println("Выводит список команд")
	} else {
		fmt.Println("Вы ввели неизвестную команду")
	}
}
