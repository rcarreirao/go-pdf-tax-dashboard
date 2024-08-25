package window

import (
	"fmt"

	"github.com/rcarreirao/pdf_balance_parser/pkg/misc/parser"
	"github.com/rivo/tview"
)

func RenderMonthlyList(tviewList *tview.List) *tview.List {
	results := parser.ListMontlhyAuctions()
	for _, result := range results {
		tviewList.AddItem("Auction month "+result.AuctionMonth, "", 'a', ShowAuctionMontly)
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

func RenderDailyList(tviewList *tview.List) *tview.List {
	results := parser.ListAuctionDays()
	for _, result := range results {
		tviewList.AddItem("Auction day "+result.AuctionDay.Format("Y-m-d"), result.CustomerCode, 'a', nil)
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

func ShowAuctionMontly() {
	fmt.Println("SHow Monthly")
}
