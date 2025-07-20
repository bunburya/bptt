package main

import (
	"flag"
	"fmt"
	"log"
	"ltt/output"
	"ltt/tfl"
	"os"
	"strings"
)

func main() {

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: ltt COMMAND\navailable commands: status, arrivals\n")
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "status":
		statusCmd := flag.NewFlagSet("status", flag.ExitOnError)
		var statusColor bool
		statusCmd.BoolVar(&statusColor, "color", false, "use colored output")
		statusCmd.Usage = func() {
			w := flag.CommandLine.Output()
			fmt.Fprintf(w, "Usage: ltt status [options] LINE_ID...\noptions:\n")
			statusCmd.PrintDefaults()
		}
		err := statusCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		lineIds := statusCmd.Args()
		if len(lineIds) == 0 {
			statusCmd.Usage()
			os.Exit(1)
		}
		lines, err := tfl.GetLineStatuses(lineIds)
		if err != nil {
			log.Fatal(err)
		}
		table := output.Table{}
		for _, line := range lines {
			row, err := line.ToRow(statusColor)
			if err != nil {
				log.Fatal(err)
			}
			table.AddRow(row)
		}
		table.Print("\t", true, statusColor)
	case "arrivals":
		arrivalsCmd := flag.NewFlagSet("arrivals", flag.ExitOnError)
		var lineFilterStr string
		arrivalsCmd.StringVar(&lineFilterStr, "lines", "",
			"comma-delimited list of lines/routes to display")
		var count int
		arrivalsCmd.IntVar(&count, "count", 0, "number of arrivals to display")
		arrivalsCmd.Usage = func() {
			w := flag.CommandLine.Output()
			fmt.Fprintf(w, "Usage: ltt arrivals [options] STOP_ID\noptions:\n")
			arrivalsCmd.PrintDefaults()
		}
		err := arrivalsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		var lines []string
		if len(lineFilterStr) > 0 {
			lines = strings.Split(lineFilterStr, ",")
		}
		if len(arrivalsCmd.Args()) != 1 {
			arrivalsCmd.Usage()
			os.Exit(1)
		}
		arrivals, err := tfl.GetStopArrivals(arrivalsCmd.Args()[0], lines, count)
		if err != nil {
			log.Fatal(err)
		}

		table := output.Table{}
		for _, arr := range arrivals.Arrivals {
			row := arr.ToRow()
			table.AddRow(row)
		}
		table.Print("\t", true, false)
	}
}
