package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api"
	"github.com/sky-uk/go-brocade-vtm/api/rule"
	"github.com/sky-uk/go-rest-api"
	"io/ioutil"
	"os"
)

var ruleName string
var trafficScriptFile string

func createRule(client *rest.Client, flagSet *flag.FlagSet) {

	headers := make(map[string]string)
	headers["Content-Type"] = "application/octet-stream"
	headers["Content-Transfer-Encoding"] = "BINARY"
	client.Headers = headers
	client.Debug = true

	if ruleName == "" {
		fmt.Printf("\nName argument is required. Usage: -name vtm-rule-name\n")
		os.Exit(1)
	}

	if trafficScriptFile == "" {
		fmt.Printf("\nRule argument is required. Usage: -script trafficScript\n")
		os.Exit(1)
	}

	trafficScript, fileErr := ioutil.ReadFile(trafficScriptFile)
	if fileErr != nil {
		fmt.Printf("\nError reading file %s\n", trafficScriptFile)
		os.Exit(2)
	}
	createRuleAPI := rule.NewCreate(ruleName, trafficScript)

	err := client.Do(createRuleAPI)
	if err != nil {
		errObj := *createRuleAPI.ErrorObject().(*api.VTMError)
		PrettyPrintErrorObj(errObj)
		os.Exit(3)
	}
	fmt.Printf("\nSuccessfully created rule %s\n", ruleName)
}

func init() {
	createRuleFlags := flag.NewFlagSet("rule-create", flag.ExitOnError)
	createRuleFlags.StringVar(&ruleName, "name", "", "usage: -name vtm-rule-name")
	createRuleFlags.StringVar(&trafficScriptFile, "script", "", "usage: -script location-of-traffic-script-file")
	RegisterCliCommand("rule-create", createRuleFlags, createRule)
}
