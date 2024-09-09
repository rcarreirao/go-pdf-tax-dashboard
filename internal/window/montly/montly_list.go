package monthly

import (
	"bytes"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/trading_note"
	trading_note_repository "github.com/rcarreirao/pdf_balance_parser/pkg/repository/trading_note"
	"github.com/rivo/tview"
)

var header = []string{"Valor dos Negócios",
	"Tx Reg BM&F",
	"Tx BM&F",
	"Ajuste Day Trade",
	"Custo Operacional",
	"Total Liq.",
	"Total nota"}

var TradingNoteRepository trading_note_repository.TradingNoteRepository

type MontlhyList struct {
	app                  *tview.Application
	tviewTable           *tview.Table
	selectedAuctionDayId int
}

func (ml *MontlhyList) GetTable() *tview.Table {
	return ml.tviewTable
}

func (ml *MontlhyList) SetSelectedAuctionDayId(auctionDayId int) *MontlhyList {
	ml.selectedAuctionDayId = auctionDayId
	return ml
}

func (ml *MontlhyList) Start(app *tview.Application) *MontlhyList {
	ml.app = app
	ml.tviewTable = tview.NewTable()
	return ml
}

func (ml *MontlhyList) RenderTableList() {
	var rowTable int
	var buffer bytes.Buffer
	TradingNoteRepository.New()
	list := TradingNoteRepository.Search(trading_note.TradingNoteSummaries{AuctionDayID: ml.selectedAuctionDayId})
	ml.printTableHeader()
	rows := len(list)
	for row := 0; row < rows; row++ {
		rowTable = row + 1
		buffer.Reset()
		buffer.WriteString(strconv.FormatFloat(float64(list[row].BusinessValue), 'f', 2, 32))
		buffer.WriteString(list[row].BusinessValueOp)
		ml.printLine(rowTable, 0, buffer.String())
		ml.printLine(rowTable, 1, strconv.FormatFloat(float64(list[row].TaxRegisterBmf), 'f', 2, 32))
		buffer.Reset()
		buffer.WriteString(strconv.FormatFloat(float64(list[row].TaxBmf), 'f', 2, 32))
		buffer.WriteString(list[row].TaxBmfOp)
		ml.printLine(rowTable, 2, buffer.String())

		buffer.Reset()
		buffer.WriteString(strconv.FormatFloat(float64(list[row].DayTradeAdjustment), 'f', 2, 32))
		buffer.WriteString(list[row].DayTradeAdjustmentOp)
		ml.printLine(rowTable, 3, buffer.String())

		buffer.Reset()
		buffer.WriteString(strconv.FormatFloat(float64(list[row].TotalOperationalCosts), 'f', 2, 32))
		buffer.WriteString(list[row].TotalOperationalCostsOp)
		ml.printLine(rowTable, 4, buffer.String())

		buffer.Reset()
		buffer.WriteString(strconv.FormatFloat(float64(list[row].TotalNet), 'f', 2, 32))
		buffer.WriteString(list[row].TotalNetOp)
		ml.printLine(rowTable, 5, buffer.String())

		buffer.Reset()
		buffer.WriteString(strconv.FormatFloat(float64(list[row].TotalNetInvoice), 'f', 2, 32))
		buffer.WriteString(list[row].TotalNetInvoiceOp)
		ml.printLine(rowTable, 6, buffer.String())
	}
}

func (ml *MontlhyList) printTableHeader() {
	cols, rows := len(header), 1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorYellow
			ml.tviewTable.SetCell(r, c,
				tview.NewTableCell(header[c]).
					SetTextColor(color).
					SetAlign(tview.AlignCenter))
		}
	}
}

func (ml *MontlhyList) printLine(row int, col int, value string) {
	color := tcell.ColorWhite
	ml.tviewTable.SetCell(row, col,
		tview.NewTableCell(value).
			SetTextColor(color).
			SetAlign(tview.AlignCenter))

}
