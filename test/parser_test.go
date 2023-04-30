package test

import (
	"image/color"
	"reflect"
	"strings"
	"testing"

	"github.com/AlinaDubchak/Lab3-Go/painter"
	"github.com/AlinaDubchak/Lab3-Go/painter/lang"

	"github.com/stretchr/testify/assert"
)

func Test_parse_func(t *testing.T) {
	tests := []struct {
		command    string
		operations painter.Operation
	}{
		{
			command:    "white",
			operations: painter.OperationFunc(painter.WhiteFill),
		},
		{
			command:    "green",
			operations: painter.OperationFunc(painter.GreenFill),
		},
		{
			command:    "black",
			operations: painter.OperationFunc(painter.BlackFill),
		},
		{
			command:    "reset",
			operations: painter.OperationFunc(painter.Reset),
		},
		{
			command:    "changeBackground 0",
			operations: painter.OperationFunc(painter.Reset),
		},
		{
			command:    "changeBackground 1",
			operations: painter.OperationFunc(painter.ResetWhite),
		},
	}
	parser := &lang.Parser{}
	for _, test := range tests {
		t.Run(test.command, func(t *testing.T) {
			operations, _ := parser.Parse(strings.NewReader(test.command))
			expectedType := reflect.TypeOf(test.operations)
			actualType := reflect.TypeOf(operations[0])
			assert.Equal(t, expectedType, actualType)
		})
	}
}

func Test_parse_struct(t *testing.T) {
	tests := []struct {
		command    string
		operations painter.Operation
	}{
		{
			command:    "figure 200 200",
			operations: &painter.FigureOp{X: 200, Y: 200, C: color.RGBA{R: 219, G: 208, B: 48, A: 1}},
		},
		{
			command:    "move 250 250",
			operations: &painter.MoveOp{X: 250, Y: 250},
		},
		{
			command:    "update",
			operations: painter.UpdateOp,
		},
		{
			command:    "bgrect 150 150 300 300 0",
			operations: &painter.BgRectOp{X1: 150, Y1: 150, X2: 300, Y2: 300, ST: 0},
		},
		{
			command:    "bgrect 200 200 400 400 1",
			operations: &painter.BgRectOp{X1: 200, Y1: 200, X2: 400, Y2: 400, ST: 1},
		},
		{
			command:    "Incorrect command: lya lya topolya(for example)",
			operations: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.command, func(t *testing.T) {
			parser := &lang.Parser{}
			operations, err := parser.Parse(strings.NewReader(test.command))
			if err != nil {
				assert.Nil(t, test.operations)
			} else {
				expectedType := reflect.TypeOf(test.operations)
				actualType := reflect.TypeOf(operations[1])
				assert.Equal(t, expectedType, actualType)
				if test.operations != nil {
					assert.Equal(t, test.operations, operations[1])
				}
			}
		})
	}
}
