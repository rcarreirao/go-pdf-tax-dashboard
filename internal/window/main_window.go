package window

import "github.com/rivo/tview"

func Execute() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	main := tview.NewList()
	app := tview.NewApplication()
	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("Tax Dashboard"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	RenderDailyList(main)

	grid.AddItem(main, 1, 0, 1, 3, 0, 0, false)
	if err := app.SetRoot(grid, true).SetFocus(main).Run(); err != nil {
		panic(err)
	}

}
