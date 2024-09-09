package daily

import (
	monthly "go_pdf_tax_dashboard/internal/window/montly"

	"github.com/rcarreirao/pdf_balance_parser/pkg/misc/parser"
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/auction"
	"github.com/rivo/tview"
)

type DailyList struct {
	app         *tview.Application
	grid        *tview.Grid
	TviewList   *tview.List
	MontlhyList *monthly.MontlhyList
	results     []auction.AuctionDays
}

func (dl *DailyList) Start(grid *tview.Grid) *DailyList {
	dl.TviewList = tview.NewList()
	dl.grid = grid
	dl.MontlhyList = new(monthly.MontlhyList)
	return dl
}

func (dl *DailyList) RenderMonthlyList() {
	dl.results = parser.ListAuctionDays()
	for _, result := range dl.results {
		dl.TviewList.SetSelectedFunc(dl.ShowAuctionMontly)
		dl.TviewList.AddItem("Auction month "+result.AuctionDay.Format("Y-m-d"), result.CustomerCode, 'a', nil)
	}
}

func (dl *DailyList) RenderDailyList() *tview.List {
	dl.TviewList.Clear()
	dl.results = parser.ListAuctionDays()
	for _, result := range dl.results {
		dl.TviewList.AddItem("Auction day "+result.AuctionDay.Format("Y-m-d"), result.CustomerCode, 'a', nil)
	}
	return dl.TviewList
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

func (dl *DailyList) updateEntries() {
	dl.app.Suspend(func() {})
	dl.RenderMonthlyList()
	dl.app.SetFocus(dl.TviewList).Run()
}

func (dl *DailyList) ShowAuctionMontly(index int, mainText string, secondaryText string, shortcut rune) {
	dl.MontlhyList.Start(dl.app).SetSelectedAuctionDayId(int(dl.results[index].ID)).RenderTableList()
	dl.grid.RemoveItem(dl.TviewList)
	dl.grid.AddItem(dl.MontlhyList.GetTable(), 1, 0, 1, 3, 0, 0, false)
}

func (dl *DailyList) ShowAuctionDay() {
	dl.updateEntries()
}
