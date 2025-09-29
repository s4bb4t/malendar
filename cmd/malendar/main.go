package main

import (
	"fmt"

	"github.com/s4bb4t/malendar/internal/calendar"
)

func main() {
	calendar.Greeting()
	fmt.Println("--------------------------------------------------------")
	calendar.Commands()
	fmt.Println("--------------------------------------------------------")
	fmt.Print("user: ")

	//scanner := bufio.NewScanner(os.Stdin)
	//if scanner.Scan() {
	//	input := scanner.Text()
	//	fmt.Printf("Команда: %s\n", input)
	//}
}
