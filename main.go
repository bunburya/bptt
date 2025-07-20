package main

import (
	"flag"
	"log"
	"ltt/output"
	"ltt/tfl"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "status":
		statusCmd := flag.NewFlagSet("status", flag.ExitOnError)
		var statusColor bool
		statusCmd.BoolVar(&statusColor, "color", false, "use colored output")
		err := statusCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		lineIds := statusCmd.Args()
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
