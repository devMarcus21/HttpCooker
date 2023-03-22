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
	url := flag.String("url", "", "")
	httpMethod := flag.String("method", "GET", "")
	bearerString := flag.String("bearer", "", "")
	groups := flag.Int("groups", 1, "")
	sets := flag.Int("sets", 1, "")
	calls := flag.Int("calls", 0, "")
	con := flag.Bool("con", false, "")

	flag.Parse()

	// Check for valid HTTP method
	method := strings.ToUpper(*httpMethod)
	if !isValidHttpMethod(method) {
		log.Fatal("Invalid HTTP method")
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
