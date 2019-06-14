package output

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

// Table struct
type Table struct {
	t *simpletable.Table
}

// Option option func
type Option func(*Table)

// NewTable constructor
func NewTable(options ...Option) *Table {
	table := &Table{
		t: simpletable.New(),
	}

	table.t.SetStyle(simpletable.StyleUnicode)

	for _, opt := range options {
		opt(table)
	}

	return table
}

// Header header option
func Header(align int, headerTexts ...string) Option {
	return func(table *Table) {
		cells := make([]*simpletable.Cell, 0, len(headerTexts))
		for _, headerTxt := range headerTexts {
			cells = append(cells, &simpletable.Cell{
				Align: align,
				Text:  headerTxt,
			})
		}

		table.t.Header = &simpletable.Header{
			Cells: cells,
		}
	}
}

// Footer footer option
func Footer(align int, footerTxt string) Option {
	return func(table *Table) {
		table.t.Footer = &simpletable.Footer{
			Cells: []*simpletable.Cell{
				{},
				{},
				{Align: align, Text: footerTxt},
			},
		}
	}
}

// AddCells add table cells
func (t *Table) AddCells(star int, packageURL, description string) {
	r := []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", star)},
		{Text: packageURL},
		{Text: description},
	}

	t.t.Body.Cells = append(t.t.Body.Cells, r)
}

// Println output contents stdout
func (t *Table) Println() {
	fmt.Println()
	t.t.Println()
}
