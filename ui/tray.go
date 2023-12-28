package ui

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/energye/systray"
)

var dashboardWindow *DashboardWindow

func RunSystray() {
	systray.SetIcon(getIcon("app.ico"))
	systray.SetTitle("Systray")
	systray.SetTooltip("Systray")
	systray.SetOnClick(func(menu systray.IMenu) {
		ShowMainWindow()
	})
	systray.SetOnDClick(func(menu systray.IMenu) {
		ShowMainWindow()
	})
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

	mShow := systray.AddMenuItem("Show Window", "Show the main window")
	mShow.Click(func() {
		ShowMainWindow()
	})
	mQuit := systray.AddMenuItem("Quit", "Quit the application")
	mQuit.Click(func() {
		ExitSystray()
	})

}

func ExitSystray() {
	systray.Quit()
	os.Exit(0)
}

func ShowMainWindow() {
	launchDashboardWindow()
}

func launchDashboardWindow() {

	if isMainWindowRunning() {
		dashboardWindow.SetVisible(true)
		dashboardWindow.BringToTop()
		dashboardWindow.Activate()
		dashboardWindow.Show()
		dashboardWindow.Focused()
	} else {
		RunDashboard()
	}
}

func isMainWindowRunning() bool {
	if dashboardWindow == nil {
		return false
	}
	return dashboardWindow.Visible()
}

func getIcon(iconName string) []byte {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Failed to get executable path: %v", err)
	}
	executableDir := filepath.Dir(executablePath)
	iconPath := filepath.Join(executableDir, iconName)
	iconData, err := os.ReadFile(iconPath)
	if err != nil {
		fmt.Printf("Failed to read icon file: %v", err)
	}
	return iconData
}
