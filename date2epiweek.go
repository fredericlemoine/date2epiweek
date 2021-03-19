package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jmeekhof/epiweek"
)

func main() {
	os.Exit(mainWithExit())
}

func mainWithExit() (exitstatus int) {
	var err error

	var metadatafile *string = flag.String("metadata", "stdin", "Meta data file")
	var sep *string = flag.String("sep", "\t", "Column seprator, default tab")
	var column *int = flag.Int("column", 0, "Column with date (YYYY-MM-DD)")
	var header *bool = flag.Bool("header", false, "If input has a header")

	var help *bool = flag.Bool("help", false, "help")
	var helpmessage string = `Converts dates in an input files into epi_weeks and prints them on stdout.
`
	exitstatus = 0

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, helpmessage)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help {
		flag.Usage()
		exitstatus = 1
		return
	}

	var metaf *os.File
	if metaf, err = os.Open(*metadatafile); err != nil {
		exitstatus = 1
		return
	}

	var line string
	var cols []string
	var date time.Time
	var scanner *bufio.Scanner = bufio.NewScanner(bufio.NewReader(metaf))
	var first bool = true
	for scanner.Scan() {
		line = scanner.Text()
		cols = strings.Split(line, *sep)

		if len(cols) <= *column {
			fmt.Fprintf(os.Stderr, "Column %d does not exist", column)
			exitstatus = 1
			return
		}

		if first && *header {
			first = false
			fmt.Printf("epi_week\n")
			continue
		}

		if date, err = time.Parse("2006-01-02", cols[*column]); err != nil {
			fmt.Fprintf(os.Stderr, "Cannot parse date : %s", cols[*column])
			exitstatus = 1
			return
		}
		epi := epiweek.NewEpiweek(date)
		_, week := epi.Epiweek()
		fmt.Printf("%s - %d\n", date, week)
	}
	if err = scanner.Err(); err != nil {
		exitstatus = 1
		return
	}
	return
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
