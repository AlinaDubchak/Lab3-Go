package main

import (
	"net/http"

	"github.com/AlinaDubchak/Lab3-Go/painter"
	"github.com/AlinaDubchak/Lab3-Go/painter/lang"
	"github.com/AlinaDubchak/Lab3-Go/ui"
)

func main() {
	var (
		pv ui.Visualizer // Візуалізатор створює вікно та малює у ньому.

		// Потрібні для частини 2.
		opLoop painter.Loop // Цикл обробки команд.
		parser lang.Parser  // Парсер команд.
	)

	//pv.Debug = true
	pv.Title = "Simple painter"

	pv.OnScreenReady = opLoop.Start
	opLoop.Receiver = &pv

	go func() {
		http.Handle("/", lang.HttpHandler(&opLoop, &parser))
		_ = http.ListenAndServe("localhost:17000", nil)
	}()

	pv.Main()
	opLoop.StopAndWait()
}
