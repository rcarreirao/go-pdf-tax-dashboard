package monthly

import (
	"bytes"
	"go_pdf_tax_dashboard/internal/window/daily"
	"strconv"

	"github.com/rcarreirao/pdf_balance_parser/pkg/misc/parser"
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/auction"
	"github.com/rivo/tview"
)

type MonthlyList struct {
	app           *tview.Application
	grid          *tview.Grid
	TviewList     *tview.List
	TableList     *tview.Table
	ActionButtons *tview.Flex
	DailyList     *daily.DailyList
	results       []auction.AuctionMonths
}

func (ml *MonthlyList) Start(app *tview.Application, grid *tview.Grid) *MonthlyList {
	ml.TviewList = tview.NewList()
	ml.app = app
	ml.grid = grid
	ml.DailyList = new(daily.DailyList)
	return ml
}

func (ml *MonthlyList) RenderMonthlyList() {
	ml.results = parser.ListAuctionMonthly()
	var buffer bytes.Buffer
	ml.TviewList.Clear()
	for _, result := range ml.results {
		ml.TviewList.SetSelectedFunc(ml.ShowAuctionMontly)
		buffer.Reset()
		buffer.WriteString(strconv.FormatUint(uint64(result.Month), 10))
		buffer.WriteString("/")
		buffer.WriteString(strconv.FormatUint(uint64(result.Year), 10))
		ml.TviewList.AddItem("Month: "+buffer.String(), strconv.FormatUint(uint64(result.ID), 10), 'a', nil)
	}
	ml.grid.AddItem(ml.TviewList, 1, 0, 1, 3, 0, 0, false)

}

func (ml *MonthlyList) updateMonthlyEntries() {
	ml.app.Suspend(func() {})
	ml.RenderMonthlyList()
	ml.app.SetFocus(ml.TviewList).Run()
}

func (ml *MonthlyList) ShowAuctionMontly(index int, mainText string, secondaryText string, shortcut rune) {
	ml.DailyList.Start(ml.app, ml.grid).SetSelectedAuctionDayId(int(ml.results[index].ID)).RenderTableList()
	ml.grid.RemoveItem(ml.TviewList)
	ml.TableList = ml.DailyList.GetTable()
	ml.initActionButtons()
	ml.grid.AddItem(ml.TableList, 1, 0, 1, 3, 0, 0, false)
	ml.grid.AddItem(ml.ActionButtons, 2, 0, 1, 3, 0, 0, false)
	ml.app.SetFocus(ml.ActionButtons)
	//ml.grid.AddItem(ml.DailyList.PrintExitButton(), 2, 0, 1, 1, 0, 0, false)
}

func (ml *MonthlyList) initActionButtons() *MonthlyList {
	ml.ActionButtons = tview.NewFlex().AddItem(
		tview.NewList().
			AddItem("Return", "return to monthly list", 'r', func() {
				ml.ShowAuctionDay()
			}).
			AddItem("Quit", "Press to exit", 'q', func() {
				ml.app.Stop()
			}), 50, 1, true)
	return ml

}

func (ml *MonthlyList) ShowAuctionDay() {
	ml.grid.RemoveItem(ml.TableList)
	ml.grid.RemoveItem(ml.ActionButtons)
	ml.updateMonthlyEntries()
}
