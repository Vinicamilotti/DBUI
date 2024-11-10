package components

import (
	"database/sql"
	"fmt"
	interfaces "github.com/Vinicamilotti/DBTUI/interface"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"reflect"
)

type TableManager struct {
	Title   string
	DataSet *sql.Rows
	Table   table.Model
}

func (t *TableManager) Init() tea.Cmd {
	return nil
}

func (t *TableManager) CreateColumns() {
	col := []table.Column{}
	datasetCol, derr := t.DataSet.Columns()
	datasetTypes, terr := t.DataSet.ColumnTypes()
	if derr != nil {
		panic(derr)
	}
	if terr != nil {
		panic(terr)
	}
	for i, c := range datasetCol {
		colName := fmt.Sprintf("%s (%s)", c, datasetTypes[i].DatabaseTypeName())
		createCol := table.Column{
			Title: colName,
			Width: len(colName),
		}

		col = append(col, createCol)
	}
	t.Table.SetColumns(col)
}

func (t *TableManager) FeedRows() {
	rows := []table.Row{}
	columnTypes, err := t.DataSet.ColumnTypes()
	if err != nil {
		panic(err)
	}
	rowValues := make([]reflect.Value, len(columnTypes))
	for i := 0; i < len(columnTypes); i++ {
		rowValues[i] = reflect.New(reflect.PointerTo(columnTypes[i].ScanType()))
	}

	resultList := [][]interface{}{}
	for t.DataSet.Next() {
		rowResult := make([]interface{}, len(columnTypes))
		for i := 0; i < len(columnTypes); i++ {
			rowResult[i] = rowValues[i].Interface()
		}

		if dbErr := t.DataSet.Scan(rowResult...); err != nil {
			panic(dbErr)
		}

		for i := 0; i < len(rowValues); i++ {
			if rv := rowValues[i].Elem(); rv.IsNil() {
				rowResult[i] = nil
			} else {
				rowResult[i] = rv.Elem().Interface()
			}
		}
		resultList = append(resultList, rowResult)
	}
	for _, r := range resultList {
		createRow := []string{}
		for _, v := range r {
			createRow = append(createRow, fmt.Sprintf("%v", v))
		}
		rows = append(rows, createRow)
	}

	t.Table.SetRows(rows)
}

func (t *TableManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if t.DataSet != nil {
		t.CreateColumns()
		t.FeedRows()
	}
	return t, nil
}

func (t *TableManager) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Center, t.Title, t.Table.View())
}

func (t *TableManager) Focus() {
	t.Table.Focus()
}

func (t *TableManager) SetDataset(table string, dataset *sql.Rows) {
	t.Title = table
	t.DataSet = dataset

}

func (t *TableManager) Value() interface{} {
	return nil
}

func CreateTableManager() interfaces.CustomBubble {
	return &TableManager{
		Title:   "None",
		DataSet: nil,
		Table:   table.New(),
	}
}
