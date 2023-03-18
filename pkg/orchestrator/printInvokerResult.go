package orchestrator

import (
	"fmt"
	"text/tabwriter"
)

func printResults(writer *tabwriter.Writer, results map[int]int) {
	printNonVerbose(writer, results)
}

func printNonVerbose(writer *tabwriter.Writer, results map[int]int) {
	exceptions := 0
	informational := 0
	success := 0
	redirection := 0
	clientError := 0
	serverError := 0

	for result, count := range results {
		if result == -1 {
			exceptions += count
		} else if result < 200 {
			informational += count
		} else if result < 300 {
			success += count
		} else if result < 400 {
			redirection += count
		} else if result < 500 {
			clientError += count
		} else if result < 600 {
			serverError += count
		}
	}
	fmt.Fprintf(writer, "%d\t%d\t%d\t%d\t%d\t%d\n", informational, success, redirection, clientError, serverError, exceptions)
}
