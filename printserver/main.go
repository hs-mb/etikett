package main

import (
	"flag"

	label "github.com/hs-mb/etikett"
)

var lprBin string
var printer string

func main() {
	lprBinArg := flag.String("b", "lpr", "lpr binary")

	flag.Parse()

	addr := flag.Arg(0)
	wsAddr := flag.Arg(1)
	printer = flag.Arg(2)

	lprBin = *lprBinArg

	go TCPServer(addr)
	go WebSocketServer(wsAddr)

	select {}
}

func makePrint(source string) error {
	return label.Print(source, printer, lprBin)
}
