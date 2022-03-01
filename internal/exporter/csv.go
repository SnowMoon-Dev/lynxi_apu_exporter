package exporter

import (
	"fmt"
	"strings"
)

type table struct {
	rows          []row
	rFields       []rField
	qFieldToCells map[qField][]cell
}

type row struct {
	qFieldToCells map[qField]cell
	cells         []cell
}

type cell struct {
	qField   qField
	rField   rField
	rawValue string
}

func parseCSVIntoTable(queryResult string, qFields []qField) (table, error) {
	lines := strings.Split(strings.TrimSpace(queryResult), "\n")
	titlesLine := lines[0]
	valuesLines := lines[1:]
	rFields := toRFieldSlice(parseCSVLine(titlesLine))

	numCols := len(qFields)
	numRows := len(valuesLines)
	rows := make([]row, numRows)

	qFieldToCells := make(map[qField][]cell)
	for _, q := range qFields {
		qFieldToCells[q] = make([]cell, numRows)
	}

	for rowIndex, valuesLine := range valuesLines {
		qFieldToCell := make(map[qField]cell, numCols)
		cells := make([]cell, numCols)
		rawValues := parseCSVLine(valuesLine)
		if len(qFields) != len(rFields) {
			return table{}, fmt.Errorf("query fields (%d) and returned fields (%d) have different sizes", len(qFields), len(rFields))
		}

		for colIndex, rawValue := range rawValues {
			q := qFields[colIndex]
			r := rFields[colIndex]
			gm := cell{
				qField:   q,
				rField:   r,
				rawValue: rawValue,
			}
			qFieldToCell[q] = gm
			cells[colIndex] = gm
			qFieldToCells[q][rowIndex] = gm
		}

		gmc := row{
			qFieldToCells: qFieldToCell,
			cells:         cells,
		}

		rows[rowIndex] = gmc

	}

	return table{
		rows:          rows,
		rFields:       rFields,
		qFieldToCells: qFieldToCells,
	}, nil
}

func parseCSVLine(line string) []string {
	values := strings.Split(line, ",")
	result := make([]string, len(values))
	for i, field := range values {
		result[i] = strings.TrimSpace(field)
	}
	return result
}
