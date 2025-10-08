package main

import (
	"github.com/s4bb4t/malendar/internal/commands"
	"github.com/s4bb4t/malendar/internal/menu"
)

func main() {

	M := map[string]string{}

	commands.InfoSet(M, "Name", "Jora")

	menu.Menu(M)
}
