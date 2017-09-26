package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/go-brocade-vtm/api/monitor"
	"github.com/sky-uk/go-rest-api"
	"os"
)

func monitorGetAll(client *rest.Client, flagSet *flag.FlagSet) {

	if apiVersion != "" {
		monitor.MonitorEndpoint = "/api/tm/" + apiVersion + "/config/active/monitors/"
	}

	getAllMonitorsAPI := monitor.NewGetAll()
	err := client.Do(getAllMonitorsAPI)
	if err != nil {
		fmt.Printf("\nError retreiving the list of monitors. Error: %+v\n", err)
		os.Exit(1)
	}
	monitorList := getAllMonitorsAPI.ResponseObject().(*monitor.MonitorsList)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "HRef"}

	for _, monitorItem := range monitorList.Children {
		row := map[string]interface{}{}
		row["Name"] = monitorItem.Name
		row["HRef"] = monitorItem.HRef
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	getAllMonitorFlags := flag.NewFlagSet("monitor-show-all", flag.ExitOnError)
	getAllMonitorFlags.StringVar(&apiVersion, "apiversion", "", "usage: -apiversion 3.8")
	RegisterCliCommand("monitor-show-all", getAllMonitorFlags, monitorGetAll)
}
