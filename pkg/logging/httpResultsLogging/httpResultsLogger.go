package httpResultslogging

type HttpResultsLogger interface {
	LogHeaderMessage()
	LogHttpResults(map[int]int)
	LogPostGroup()
	LoggerCleanUp()
}
