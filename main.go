package main

import (
	"startupwin/ui"

	"github.com/energye/systray"
)

func main() {
	go startSystray()
	ui.ShowMainWindow()
}

func startSystray() {
	systray.Run(ui.RunSystray, ui.ExitSystray)
}
