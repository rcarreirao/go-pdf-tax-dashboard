package window

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rcarreirao/pdf_balance_parser/pkg/misc/parser"
	"github.com/rivo/tview"
)

type DailyList struct {
	grid       *tview.Grid
	tviewList  *tview.List
	tviewTable *tview.Table
}

func (dl *DailyList) Start(grid *tview.Grid) *DailyList {
	dl.tviewList = tview.NewList()
	dl.tviewTable = tview.NewTable()
	dl.grid = grid
	return dl
}

func (dl *DailyList) RenderMonthlyList() {
	results := parser.ListAuctionDays()
	for _, result := range results {
		dl.tviewList.AddItem("Auction month "+result.AuctionDay.Format("Y-m-d"), result.CustomerCode, 'a', dl.ShowAuctionMontly)
	}
}

func (dl *DailyList) RenderDailyList() *tview.List {
	dl.tviewList.Clear()
	results := parser.ListAuctionDays()
	for _, result := range results {
		dl.tviewList.AddItem("Auction day "+result.AuctionDay.Format("Y-m-d"), result.CustomerCode, 'a', dl.ShowAuctionDay)
	}
	return dl.tviewList
	/*
	   list := tview.NewList().

	   	AddItem("List item 1", "Some explanatory text", 'a', nil).
	   	AddItem("List item 2", "Some explanatory text", 'b', nil).
	   	AddItem("List item 3", "Some explanatory text", 'c', nil).
	   	AddItem("List item 4", "Some explanatory text", 'd', nil).
	   	AddItem("Quit", "Press to exit", 'q', func() {
	   		app.Stop()
	   	})
	*/
}

func (dl *DailyList) RenderTableList() {
	cols, rows := 10, 40
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			if c < 1 || r < 1 {
				color = tcell.ColorYellow
			}
			dl.tviewTable.SetCell(r, c,
				tview.NewTableCell("a").
					SetTextColor(color).
					SetAlign(tview.AlignCenter))
		}
	}
}

/*
	func (dl *DailyList) monitoringTrips(g *Gui) {
		ticker := time.NewTicker(5 * time.Minute)

LOOP:

		for {
			select {
			case <-ticker.C:
				dl.updateEntries(g)
			case <-g.state.stopChans["trips"]:
				ticker.Stop()
				break LOOP
			}
		}
	}
*/
func (dl *DailyList) updateEntries() {
	app.Suspend(func() {})
	dl.RenderMonthlyList()
	app.SetFocus(dl.tviewList).Run()
}

func (dl *DailyList) ShowAuctionMontly() {
	dl.RenderTableList()
	dl.grid.RemoveItem(dl.tviewList)
	dl.grid.AddItem(dl.tviewTable, 1, 0, 1, 3, 0, 0, false)

}

func (dl *DailyList) ShowAuctionDay() {
	dl.updateEntries()
}
