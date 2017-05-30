package main

import (
	"fmt"
	"github.com/sky-uk/go-brocade-vtm"
	"github.com/sky-uk/go-brocade-vtm/api/pool"
)

// RunPoolExample  - runs examples
func RunPoolExample(vtmAddress, vtmUser, vtmPassword string, debug bool) {
	vtmClient := brocadevtm.NewVTMClient(vtmAddress, vtmUser, vtmPassword, true, debug)

	/*
	//Example to get all the pools
	getAllAPI := pool.NewGetAll()

	// make api call.
	err := vtmClient.Do(getAllAPI)

	// check if there were any errors
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// check the status code and proceed accordingly.
	if getAllAPI.StatusCode() == 200 {
		AllPools := getAllAPI.GetResponse().ChildPools
		for _, pool := range AllPools {
			fmt.Printf("Name: %-20s HRef: %-20s\n", pool.Name, pool.Href)
		}
	} else {
		fmt.Println("Status code:", getAllAPI.StatusCode())
		fmt.Println("Response: ", getAllAPI.ResponseObject())
	}

	//Example to get a single pool
	getSingleAPI := pool.NewGetSingle("pool_test_rui_2")
	// make api call.
	err2 := vtmClient.Do(getSingleAPI)
	if err2 != nil {
		fmt.Println("Error: ", err2)
	}
	if getSingleAPI.StatusCode() == 200 {
		MyPool := getSingleAPI.GetResponse().Properties
		fmt.Println("Response: ", getSingleAPI.ResponseObject())
		fmt.Println(MyPool)
	} else {
		fmt.Println("Status code:", getSingleAPI.StatusCode())
		fmt.Println("Response: ", getSingleAPI.ResponseObject())
	}


*/

        pool_nodes := []pool.MemberNodes{}

	pool_nodes = append(pool_nodes,pool.NewMemberNodes("127.0.0.1:80",1,"active",1 ))
	pool_nodes = append(pool_nodes,pool.NewMemberNodes("127.0.0.1:82",2,"active",1 ))

	CreateAPI := pool.NewCreate("pool_test_rui_5", pool_nodes, []string{"ping"})
	errCreate := vtmClient.Do(CreateAPI)
	if errCreate != nil {
		fmt.Println("Error Creating:", errCreate)
	}
	if CreateAPI.StatusCode() == 200 {
		fmt.Println("Created")
	} else {
		fmt.Println("Status code:", CreateAPI.StatusCode())
		fmt.Println("Response: ", CreateAPI.ResponseObject())
	}

	/*DeleteAPI := pool.NewDelete("pool_test_rui_4")
	// make api call.
	errDelete := vtmClient.Do(DeleteAPI)
	if errDelete != nil {
		fmt.Println("Error: ", errDelete)
	}
	if DeleteAPI.StatusCode() == 200 {
		fmt.Println("Deleted ")
	} else {
		fmt.Println("Status code:", DeleteAPI.StatusCode())
		fmt.Println("Response: ", DeleteAPI.ResponseObject())
	}*/

}
