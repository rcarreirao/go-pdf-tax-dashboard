package window

import (
	"github.com/rcarreirao/pdf_balance_parser/pkg/misc/parser"
	"github.com/rivo/tview"
)

func RenderDailyList(tviewList *tview.List) *tview.List {
	results, err := parser.ListAuctionDays().Rows()
	if err == nil {
		defer results.Close()
		for results.Next() {
			var Id int
			var AuctionDay string
			results.Scan(&Id, &AuctionDay)
			tviewList.AddItem("List item "+AuctionDay, "Some explanatory text", 'a', nil)
		}
	}

	return tviewList
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
