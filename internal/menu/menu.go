package menu

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/s4bb4t/malendar/internal/calendar"
	"github.com/s4bb4t/malendar/internal/commands"
)

func Greeting() {
	fmt.Println("Привет, это твой календарь!\nЧто тебя интересует?\n")
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

func weekCommand(e *bool, t *time.Time, scanner *bufio.Scanner) {
	fmt.Println(commands.Week(*t))
	WeekFunctions()
	fmt.Print("\nДействие: ")
	scanner.Scan()
	switch scanner.Text() {
	case "n":
		*t = t.AddDate(0, 0, 7)
	case "p":
		*t = t.AddDate(0, 0, -7)
	case "q":
		*e = true
	}
}

func monthCommand(e *bool, t *time.Time, scanner *bufio.Scanner) {
	fmt.Println(commands.Month(*t))
	MonthFunctions()
	fmt.Print("\nДействие: ")
	scanner.Scan()
	switch scanner.Text() {
	case "n":
		*t = t.AddDate(0, 1, 0)
	case "p":
		*t = t.AddDate(0, -1, 0)
	case "q":
		*e = true
	}
}

func Menu(Info map[string]string) {
	Greeting()
	Commands()
	scanner := bufio.NewScanner(os.Stdin)
	r := regexp.MustCompile(`^month\s+(\d{2})\.(\d{2})$`)
	r2 := regexp.MustCompile(`^set\s+(\d{2})\.(\d{2})\s+(.+)$`)

	Period := make(map[string]commands.Note)

	for {
		fmt.Print("\nВаша команда: ")
		scanner.Scan()

		switch {
		case scanner.Text() == "info":
			fmt.Println("Информация о пользователе:\n")
			fmt.Println(commands.Info(Info))

		case scanner.Text() == "week":
			t := time.Now()
			exit := false
			for !exit {
				weekCommand(&exit, &t, scanner)
			}

		case scanner.Text() == "info set":

		case r.MatchString(scanner.Text()):
			t, err := time.Parse("01.02", r.FindString(scanner.Text())[6:])
			if err != nil {
				fmt.Println(err)
				return
			}
			t = t.AddDate(time.Now().Year(), 0, -(t.Day() - 1))
			exit := false
			for !exit {
				monthCommand(&exit, &t, scanner)
			}

		case scanner.Text() == "month":
			t := time.Now()
			t = t.AddDate(0, 0, -(t.Day() - 1))
			exit := false
			for !exit {
				monthCommand(&exit, &t, scanner)
			}

		case scanner.Text() == "q":
			fmt.Println("\nЧто ж...Пришло время прощаться...")
			return

		case r2.MatchString(scanner.Text()):
			number := r2.FindString(scanner.Text())[4:9]
			phase := r2.FindString(scanner.Text())[10:]
			var p calendar.Phase
			switch phase {
			case "M", "Menstrual":
				p = calendar.MenstrualPhase
			}
			commands.Set(Period, number, p)

		default:
			fmt.Println("Ерунда! Что-то не то!")
		}
	}

}
