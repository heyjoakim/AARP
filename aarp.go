package main

import (
	"github.com/heyjoakim/AARP/helpers"
	"github.com/heyjoakim/AARP/logic"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	root := "/home/hey/git/Zeeguu-API"

	extension := ".py"
	exclusion := "__init__.py"
	fileSlice := logic.Walk(root, extension, exclusion)
	helpers.CountFiles(fileSlice)

	f, err := os.Create("data.wsd")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("@startuml\ntitle Automated Diagram\nskinparam nodesep 100\nskinparam ranksep 100\n")
	dependencies := make(map[string]bool)
	nodes := make(map[string]bool)

	for _, j := range *fileSlice {
		file, imports := helpers.Lines(j)
		dependencies, nodes = helpers.ToSet(dependencies, nodes, file, imports)
	}
	helpers.ToFile(f, nodes)
	helpers.ToFile(f, dependencies)
	f.WriteString("@enduml")
	defer f.Close()

}
