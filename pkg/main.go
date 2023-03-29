package main

import (
	"flag"
	"log"
	"strings"

	httpclient "github.com/devMarcus21/HttpCooker/pkg/httpClientWrapper"
	"github.com/devMarcus21/HttpCooker/pkg/invokers"
	"github.com/devMarcus21/HttpCooker/pkg/orchestrator"
)

func main() {
	url := flag.String("url", "", "Url to make calls to")
	httpMethod := flag.String("method", "GET", "Http method to make call with")
	bearerString := flag.String("bearer", "", "Bearer authorization token to make call with")
	groups := flag.Int("groups", 1, "Number of groups")
	sets := flag.Int("sets", 1, "Number of sets per group")
	calls := flag.Int("calls", 0, "Number of HTTP calls to the url per set")
	con := flag.Bool("con", false, "Make calls concurrently or not")
	waitAfterGroup := flag.Int("wg", 0, "Wait in seconds to wait between each group")

	flag.Parse()

	// Check for valid HTTP method
	method := strings.ToUpper(*httpMethod)
	if !isValidHttpMethod(method) {
		log.Fatal("Invalid HTTP method")
	}

	// Check if url is empty
	if *url == "" {
		log.Fatal("Url is empty")
	}

	// Select invoker function
	invoker := invokers.InvokeSync
	if *con {
		invoker = invokers.InvokerAsync
	}

	httpClientFactory := httpclient.BuildHttpClientFunction

	orchestrator.GetOrchestratorBuilder().
		SetUrl(*url).
		SetMethod(method).
		SetBearerString(*bearerString).
		SetGroups(*groups).
		SetSets(*sets).
		SetCalls(*calls).
		AddHttpClientFunctionCallback(httpClientFactory).
		SetInvoker(invoker).
		AddHttpResultsLoggerBasic().
		SetWaitAfterGroup(*waitAfterGroup).
		Build().
		Run()
}

func isValidHttpMethod(httpMethod string) bool {
	validHttpMethods := map[string]bool{
		"GET":    true,
		"DELETE": true,
		"PUT":    true,
		"POST":   true,
	}

	return validHttpMethods[httpMethod]
}
