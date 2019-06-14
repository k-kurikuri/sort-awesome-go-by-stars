package output_test

import (
	"reflect"
	"testing"

	"github.com/k-kurikuri/sort-awesome-go-by-stars/output"
)

func TestNewTable(t *testing.T) {
	table := output.NewTable()

	if table == nil {
		t.Error("table is not nil")
	}

	if w, g := "*output.Table", reflect.TypeOf(table).String(); w != g {
		t.Errorf("type error. want: %s, but %s", w, g)
	}
}
