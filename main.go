package main

import (
	"fmt"
	"github.com/sne-ot-research/peerscanner/ipfsx"
	"github.com/evilsocket/islazy/tui"
	"github.com/sne-ot-research/peerscanner/scanner"
	"os"
	"strconv"
)

func main()  {
	//typePtr := flag.String("type", "cos", "type of introspection to perform on peers")
	//flag.Parse()

	scanType := os.Args[1]

	peerIPs := ipfsx.GetPeers()

	if scanType == "cors" {
		scanRes := scanner.CorsScan(peerIPs)
		tableRes := [][]string{}
		for k, v := range scanRes {
			entry := []string{k, strconv.FormatBool(!isEmpty(v))}
			tableRes = append(tableRes, entry)
		}

		tui.Table(os.Stdout, []string{"IP Address", "CORS Exposed"}, tableRes)

	} else {
		fmt.Println("Unsupported scan type", scanType)
	}
}

func isEmpty(arrays [][]string) bool {
	for _, v := range arrays {
		if len(v) != 0 {
			return false
		}
	}
	return true
}