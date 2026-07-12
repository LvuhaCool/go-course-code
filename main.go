package main

/* type user struct {
	Name  string
	Age   int
	Email string
} */

func main() {
	// Первая программа:
	// fmt.Println("Hello, World!")

	// Переменные:
	/* firstVar := 1
	var secondVar string = "string"

	fmt.Println(firstVar, secondVar) */

	// Условные ветвления:
	/* variable := 67
	if variable < 100 {
		fmt.Println("Меньше 100")
	} else if variable < 200 {
		fmt.Println("Меньше 200")
	} else {
		fmt.Println("Больше или равно 200")
	} */

	// Циклы:
	/* for i := 0; i <= 5; i++ {
		fmt.Println(i)
	} */

	// Функции:
	// fmt.Println(funnyFun(60, 7))

	// defer():
	// insideMain()

	// Указатели:
	/* a := 1
	b := 3

	someFunc(&a, &b)

	fmt.Println(a, b) */

	/* user1 := user{
		Name:  "LvuhaCool",
		Age:   13,
		Email: "bla@gmail.bla",
	}

	fmt.Println(user1.getName())
	fmt.Println("-----------")
	fmt.Println(user1.getAge())
	fmt.Println("-----------")
	fmt.Println(user1.getEmail())

	// fmt.Println(user1) */

	// Модули и пакеты
	// pp.Println(67)

	// fmt.Println(paketik.Paketik2())

	// Массивы, слайсы, мапы
	/* arr := []string{"1", "2", "3", "4"}
	pp.Println(len(arr), cap(arr))
	arr = append(arr, "5")
	pp.Println(len(arr), cap(arr))

	for i, v := range arr {
		pp.Println(i, v)
	}

	mapThingy := map[string]int{
		"Зима":  1,
		"Весна": 2,
		"Лето":  3,
		"Осень": 4,
	}

	fmt.Println(mapThingy)
	fmt.Println(mapThingy["Зима"]) */
}

/* func funnyFun(a int, b int) int {
	return a + b
} */

/* func insideMain() {
	defer func() {
		fmt.Println("Defer")
	}()
	fmt.Println(67)
	anotherFunc()
}

func anotherFunc() {
	defer func() {
		fmt.Println("Defer of anotherFunc")
	}()
	fmt.Println(676767)
} */

/* func someFunc(a *int, b *int) {
	fmt.Println("someFunc working...")
	*a += 1
	*b += 1
} */

/* func (u user) getName() string {
	return u.Name
}

func (u user) getAge() int {
	return u.Age
}

func (u user) getEmail() string {
	return u.Email
} */
