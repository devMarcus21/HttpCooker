package httpResultslogging

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type HttpResultsLoggerBasic struct {
	writer *tabwriter.Writer
}

func (logger *HttpResultsLoggerBasic) LogHeaderMessage() {
	fmt.Fprintf(logger.writer, "%s\t%s\t%s\t%s\t%s\t%s\t\n", "1xx", "2xx", "3xx", "4xx", "5xx", "exception")
}

func (logger *HttpResultsLoggerBasic) LogHttpResults(results map[int]int) {
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
	fmt.Fprintf(logger.writer, "%d\t%d\t%d\t%d\t%d\t%d\n", informational, success, redirection, clientError, serverError, exceptions)
}

func (logger *HttpResultsLoggerBasic) LogPostGroup() {
	fmt.Fprintf(logger.writer, "\n")
}

func (logger *HttpResultsLoggerBasic) LoggerCleanUp() {
	logger.writer.Flush()
}

func BuildHttResultsLoggerBasic() *HttpResultsLoggerBasic {
	return &HttpResultsLoggerBasic{writer: tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)}
}
