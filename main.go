package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xor111xor/wifi-api/cmd"
)

func main() {
	// Configure flags
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Api for wifi statistic\n")
		fmt.Fprintf(os.Stderr, "Using variable PORT for change port, default 8080\n")
		flag.PrintDefaults()
	}
	flagVersion := flag.Bool("v", false, "show version")
	flagCron := flag.String("c", "*/1 * * * *", "Set cron format, default for per minute")
	flag.Parse()

	var AppVersion string

	if *flagVersion {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	if err := cmd.Run(flagCron); err != nil {
		panic(err)
	}
}
