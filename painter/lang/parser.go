package lang

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/AlinaDubchak/Lab3-Go/painter"
)

type Parser struct {
	lastBackgroundColor painter.Operation
	lastBackgroundRect  *painter.BgRectOp
	figures             []*painter.FigureOp
	moveOperations      []painter.Operation
	updateOperation     painter.Operation
}

func (prs *Parser) init() {
	if prs.lastBackgroundColor == nil {
		prs.lastBackgroundColor = painter.OperationFunc(painter.Reset)
	}
	if prs.updateOperation != nil {
		prs.updateOperation = nil
	}
}

func (prs *Parser) resetParserState() {
	prs.lastBackgroundColor, prs.lastBackgroundRect, prs.figures, prs.moveOperations, prs.updateOperation = nil, nil, nil, nil, nil
}
func (prs *Parser) Parse(input io.Reader) ([]painter.Operation, error) {
	prs.init()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		if err := prs.parseCommand(line); err != nil {
			return nil, fmt.Errorf("failed to parse command: %v", err)
		}
	}

	return prs.generatePaintOperations(), nil
}
func (prs *Parser) generatePaintOperations() []painter.Operation {
	var result []painter.Operation
	if prs.lastBackgroundColor != nil {
		result = append(result, prs.lastBackgroundColor)
	}

	if prs.lastBackgroundRect != nil {
		result = append(result, prs.lastBackgroundRect)
	}

	result = append(result, prs.moveOperations...)
	prs.moveOperations = nil

	for _, figure := range prs.figures {
		result = append(result, figure)
	}

	if prs.updateOperation != nil {
		result = append(result, prs.updateOperation)
	}
	return result
}

func (prs *Parser) parseCommand(commandLine string) error {
	parts := strings.Fields(commandLine)
	instruction := parts[0]
	args := parts[1:]
	intArgs := make([]int, 0, len(args))
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			return fmt.Errorf("Unable to parse argument %v: %v", arg, err)
		}
		intArgs = append(intArgs, i)
	}

	if instruction == "white" {
		prs.lastBackgroundColor = painter.OperationFunc(painter.WhiteFill)
	} else if instruction == "black" {
		prs.lastBackgroundColor = painter.OperationFunc(painter.BlackFill)
	} else if instruction == "move" {
		moveOp := painter.MoveOp{X: intArgs[0], Y: intArgs[1], Figures: prs.figures}
		prs.moveOperations = append(prs.moveOperations, &moveOp)
	} else if instruction == "green" {
		prs.lastBackgroundColor = painter.OperationFunc(painter.GreenFill)
	} else if instruction == "bgrect" {
		prs.lastBackgroundRect = &painter.BgRectOp{X1: intArgs[0], Y1: intArgs[1], X2: intArgs[2], Y2: intArgs[3], ST: intArgs[4]}
	} else if instruction == "figure" {
		col := color.RGBA{R: 219, G: 208, B: 48, A: 1}
		figure := painter.FigureOp{X: intArgs[0], Y: intArgs[1], C: col}
		prs.figures = append(prs.figures, &figure)
	} else if instruction == "update" {
		prs.updateOperation = painter.UpdateOp
	} else if instruction == "reset" {
		prs.resetParserState()
		prs.lastBackgroundColor = painter.OperationFunc(painter.Reset)
	} else if instruction == "changeBackground" {
		prs.resetParserState()
		if intArgs[0] == 0 {
			prs.lastBackgroundColor = painter.OperationFunc(painter.Reset)
		} else if intArgs[0] == 1 {
			prs.lastBackgroundColor = painter.OperationFunc(painter.ResetWhite)
		} else if intArgs[0] == 2 {
			prs.lastBackgroundColor = painter.OperationFunc(painter.ResetGreen)
		}
	} else {
		return fmt.Errorf("Unable to parse command %v", commandLine)
	}
	return nil
}
