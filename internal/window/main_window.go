package window

import "github.com/rivo/tview"

var (
	app *tview.Application
)

func Execute() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	app = tview.NewApplication()
	grid := tview.NewGrid().
		SetRows(3, 0).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(newPrimitive("Tax Dashboard"), 0, 0, 1, 3, 0, 0, false)

	dailyList := new(DailyList).Start(grid)

	dailyList.RenderMonthlyList()
	grid.AddItem(dailyList.tviewList, 1, 0, 1, 3, 0, 0, false)
	//grid.AddItem(dailyList.tviewTable, 1, 0, 1, 3, 0, 0, false)
	if err := app.SetRoot(grid, true).SetFocus(dailyList.tviewList).Run(); err != nil {
		panic(err)
	}

}
