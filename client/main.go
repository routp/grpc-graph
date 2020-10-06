package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var address string
	var graph string
	var graphId string
	var src string
	var dest string

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createCmd.StringVar(&graph, "graph", "", "graph edges example: A-B,B-C or 'null' for null graph")
	createCmd.StringVar(&address,"address", ":9007", "server listener address")

	findCmd := flag.NewFlagSet("find", flag.ExitOnError)
	findCmd.StringVar(&graphId, "id", "", "graph id")
	findCmd.StringVar(&src, "src", "", "source node")
	findCmd.StringVar(&dest, "dest", "", "destination node")
	findCmd.StringVar(&address,"address", ":9007", "server listener address")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteCmd.StringVar(&graphId, "id", "", "graph id")
	deleteCmd.StringVar(&address,"address", ":9007", "server listener address")


	if len(os.Args) < 2 {
		fmt.Println("expected operations 'create' or 'find' or 'delete'")
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "create":
		createCmd.Parse(os.Args[2:])
		if createCmd.Parsed() {
			if graph == ""{
				createCmd.PrintDefaults()
				os.Exit(1)
			}
			createGraph(address, graph)
		}
	case "find":
		findCmd.Parse(os.Args[2:])
		if findCmd.Parsed() {
			if graphId == ""|| src == "" || dest == ""{
				findCmd.PrintDefaults()
				os.Exit(1)
			}
			findPath(address, graphId, src, dest)
		}
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if deleteCmd.Parsed() {
			if graphId == ""{
				deleteCmd.PrintDefaults()
				os.Exit(1)
			}
			deleteGraph(address, graphId)
		}

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}