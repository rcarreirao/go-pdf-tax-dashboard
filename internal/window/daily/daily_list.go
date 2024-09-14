package daily

import (
	"bytes"
	"math"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/auction"
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/trading_note"
	auction_repository "github.com/rcarreirao/pdf_balance_parser/pkg/repository/auction"
	trading_note_repository "github.com/rcarreirao/pdf_balance_parser/pkg/repository/trading_note"
	"github.com/rivo/tview"
)

var header = []string{"Day", "Profit/Loss",
	"Tx Reg BM&F",
	"Tx BM&F",
	"Custo Operacional",
	"Snitch",
	"Total Liq.",
	"Total Nota"}

var TradingNoteRepository trading_note_repository.TradingNoteRepository
var AuctionDayRepository auction_repository.AuctionDayRepository

type DailyList struct {
	app                    *tview.Application
	grid                   *tview.Grid
	tviewTable             *tview.Table
	selectedAuctionDayId   int
	selectedAuctionMonthId int
	TotalSummary           TotalSummary
	TotalLines             int
}

type TotalSummary struct {
	TotalProfitLoss      float64
	TotalBmf             float64
	TotalOperationalCost float64
	TotalSnitch          float64
	TotalForTax          float64
	TotalTaxForPayment   float64
}

func (dl *DailyList) GetTable() *tview.Table {
	return dl.tviewTable
}

func (dl *DailyList) SetSelectedAuctionDayId(auctionDayId int) *DailyList {
	dl.selectedAuctionDayId = auctionDayId
	return dl
}

func (dl *DailyList) SetSelectedAuctionMonthId(auctionMonthId int) *DailyList {
	dl.selectedAuctionMonthId = auctionMonthId
	return dl
}

func (dl *DailyList) Start(app *tview.Application, grid *tview.Grid) *DailyList {
	dl.app = app
	dl.grid = grid
	dl.tviewTable = tview.NewTable()
	return dl
}

func (dl *DailyList) RenderTableList() {
	var rowTable int
	var buffer bytes.Buffer
	AuctionDayRepository.New()
	TradingNoteRepository.New()
	listAuctionDays := AuctionDayRepository.Search(auction.AuctionDays{AuctionMonthID: dl.selectedAuctionMonthId})
	dl.printTableHeader()
	for i, auctionday := range listAuctionDays {
		list := TradingNoteRepository.FindBy(trading_note.TradingNoteSummaries{AuctionDayID: int(auctionday.ID)})
		rows := len(listAuctionDays)
		dl.sumTotal(list)
		for row := 0; row < rows; row++ {
			rowTable = i + 1
			buffer.Reset()
			buffer.WriteString(auctionday.AuctionDay.Format("02/01/06"))
			dl.printLine(rowTable, 0, buffer.String())

			buffer.Reset()
			buffer.WriteString(strconv.FormatFloat(float64(list.GetBusinessValue()), 'f', 2, 32))
			buffer.WriteString(list.BusinessValueOp)
			dl.printLine(rowTable, 1, buffer.String())
			dl.printLine(rowTable, 2, strconv.FormatFloat(float64(list.TaxRegisterBmf), 'f', 2, 32))

			buffer.Reset()
			buffer.WriteString(strconv.FormatFloat(float64(list.TaxBmf), 'f', 2, 32))
			buffer.WriteString(list.TaxBmfOp)
			dl.printLine(rowTable, 3, buffer.String())

			buffer.Reset()
			buffer.WriteString(strconv.FormatFloat(float64(list.DayTradeAdjustment), 'f', 2, 32))
			buffer.WriteString(list.DayTradeAdjustmentOp)
			dl.printLine(rowTable, 4, buffer.String())

			buffer.Reset()
			buffer.WriteString(strconv.FormatFloat(float64(list.TotalOperationalCosts), 'f', 2, 32))
			buffer.WriteString(list.TotalOperationalCostsOp)
			dl.printLine(rowTable, 5, buffer.String())

			buffer.Reset()
			buffer.WriteString(strconv.FormatFloat(float64(list.IrrfDayTrade), 'f', 2, 32))
			buffer.WriteString(list.TotalNetOp)
			dl.printLine(rowTable, 6, buffer.String())

			buffer.Reset()
			buffer.WriteString(strconv.FormatFloat(float64(list.GetBusinessValue()-list.TotalOperationalCosts), 'f', 2, 32))
			buffer.WriteString(list.TotalNetInvoiceOp)
			dl.printLine(rowTable, 7, buffer.String())
		}
	}
	dl.TotalLines = len(listAuctionDays) + 1
	dl.printTableFooter()
	dl.printTableFooterTax()

}

func (dl *DailyList) printTableHeader() {
	cols, rows := len(header), 1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorYellow
			dl.tviewTable.SetCell(r, c,
				tview.NewTableCell(header[c]).
					SetTextColor(color).
					SetAlign(tview.AlignCenter))
		}
	}
}

func (dl *DailyList) printTableFooter() {
	color := tcell.ColorYellow
	var buffer bytes.Buffer
	buffer.Reset()
	buffer.WriteString(strconv.FormatFloat(float64(dl.TotalSummary.TotalProfitLoss), 'f', 2, 32))
	dl.tviewTable.SetCell(dl.TotalLines, 1,
		tview.NewTableCell(buffer.String()).
			SetTextColor(color).
			SetAlign(tview.AlignCenter))
	buffer.Reset()

	buffer.WriteString(strconv.FormatFloat(float64(dl.TotalSummary.TotalBmf), 'f', 2, 32))

	dl.tviewTable.SetCell(dl.TotalLines, 3,
		tview.NewTableCell(buffer.String()).
			SetTextColor(color).
			SetAlign(tview.AlignCenter))
	buffer.Reset()

	buffer.WriteString(strconv.FormatFloat(float64(dl.TotalSummary.TotalOperationalCost), 'f', 2, 32))

	dl.tviewTable.SetCell(dl.TotalLines, 5,
		tview.NewTableCell(buffer.String()).
			SetTextColor(color).
			SetAlign(tview.AlignCenter))
	buffer.Reset()

	buffer.WriteString(strconv.FormatFloat(float64(dl.TotalSummary.TotalSnitch), 'f', 2, 32))

	dl.tviewTable.SetCell(dl.TotalLines, 6,
		tview.NewTableCell(buffer.String()).
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

	buffer.Reset()

	buffer.WriteString(strconv.FormatFloat(float64(dl.TotalSummary.TotalForTax), 'f', 2, 32))

	dl.tviewTable.SetCell(dl.TotalLines, 7,
		tview.NewTableCell(buffer.String()).
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

}

func (dl *DailyList) printTableFooterTax() {
	var buffer bytes.Buffer

	color := tcell.ColorYellow

	dl.tviewTable.SetCell(dl.TotalLines+2, 6,
		tview.NewTableCell("Tax to be paid (DARF):").
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

	color = tcell.ColorBlue
	buffer.Reset()
	buffer.WriteString(strconv.FormatFloat(float64(dl.TotalSummary.TotalTaxForPayment), 'f', 2, 32))
	dl.tviewTable.SetCell(dl.TotalLines+2, 7,
		tview.NewTableCell(buffer.String()).
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

}

func (dl *DailyList) printLine(row int, col int, value string) {
	color := tcell.ColorWhite
	dl.tviewTable.SetCell(row, col,
		tview.NewTableCell(value).
			SetTextColor(color).
			SetAlign(tview.AlignCenter))
	dl.tviewTable.SetCell(row, col,
		tview.NewTableCell(value).
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

}

func (dl *DailyList) sumTotal(tradingNoteSummary trading_note.TradingNoteSummaries) {
	var profitLoss float64
	var taxRegisterBMF float64
	var taxBMF float64
	var operationalTax float64
	var totalOperationalCost float64
	var snitch float64
	var totalForTax float64

	profitLoss = tradingNoteSummary.GetBusinessValue()
	taxRegisterBMF = tradingNoteSummary.TaxRegisterBmf
	taxBMF = tradingNoteSummary.TaxBmf
	operationalTax = tradingNoteSummary.OperationalTax

	if profitLoss > 0 {
		snitch = math.Floor(profitLoss-(taxRegisterBMF+taxBMF)) * 0.01
	}
	totalOperationalCost = taxRegisterBMF + taxBMF + operationalTax // this should be always summed as positive
	totalForTax = (profitLoss - totalOperationalCost)
	dl.TotalSummary.TotalProfitLoss += profitLoss
	dl.TotalSummary.TotalBmf += taxBMF
	dl.TotalSummary.TotalOperationalCost += totalOperationalCost
	dl.TotalSummary.TotalSnitch += snitch
	dl.TotalSummary.TotalForTax += totalForTax
	dl.TotalSummary.TotalTaxForPayment = (dl.TotalSummary.TotalForTax * 0.20) - dl.TotalSummary.TotalSnitch
}
