package ui

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var (
	status *walk.Label
)

type DashboardWindow struct {
	*walk.MainWindow
}

func RunDashboard() error {
	var err error
	dashboardWindow = new(DashboardWindow)

	if err = (MainWindow{
		Title:    "Systray",
		AssignTo: &dashboardWindow.MainWindow,
		Icon:     "app.ico",
		Layout:   VBox{SpacingZero: true, Alignment: AlignHNearVNear},
		Size:     Size{Width: 360, Height: 300},
		Children: []Widget{
			Composite{
				Layout: HBox{SpacingZero: true},
				Children: []Widget{
					Composite{
						Layout:    VBox{SpacingZero: true},
						Alignment: AlignHCenterVCenter,
						Children: []Widget{
							Composite{
								Alignment: AlignHCenterVCenter,
								Layout:    VBox{SpacingZero: true, Margins: Margins{Bottom: 20}},
								Children: []Widget{
									Label{
										Text:      "HELLO",
										Font:      Font{PointSize: 14, Bold: true},
										Alignment: AlignHCenterVCenter,
										AssignTo:  &status,
									},
								},
							},
						},
					},
				},
			},
		},
	}).Create(); err != nil {
		fmt.Print(err)
		return err
	}

	dashboardWindow.MainWindow.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		*canceled = true
		dashboardWindow.Hide()
	})

	dashboardWindow.Run()

	return nil
}
