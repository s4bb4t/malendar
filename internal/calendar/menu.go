package calendar

import "fmt"

func Malendar() {
	fmt.Println("Malendar Запускает приложение.")
}
func Greeting() {
	fmt.Println("Привет, это твой календарь!\nЧто тебя интересует?")
}
func InfoCommand() {
	fmt.Println("info Вывовдит всю информацию о пользователе.")
}
func InfoSet() {
	fmt.Println("info set <название> <значение> Устанивливает информацию о пользователе.")
}

func MonthFunctions() {
	fmt.Println("n Позволяет вывсети для просмотра следующий месяц.")
	fmt.Println("p Позволяет вывсети для просмотра предыдущий месяц.")
	fmt.Println("q Позволяет выйти из просмотра месяца.")
}
func MonthAny() {
	fmt.Println("month <month> Позволяет просмотреть любой месяц.\nВ случае отсутствия числа или названия месяца, будет выведен текущий месяц.")
}
func WeekBasic() {
	fmt.Println("week Позволяет просмотреть текущую неделю.")
}
func WeekFunctions() {
	fmt.Println("n Позволяет вывсети для просмотра следующую неделю.")
	fmt.Println("p Позволяет вывсети для просмотра предыдущую неделю.")
	fmt.Println("q Позволяет выйти из просмотра недели.")
}
func Commands() {
	WeekBasic()
	MonthAny()
	InfoCommand()
	InfoSet()
}
