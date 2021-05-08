package main

import (
	"github.com/heyjoakim/AARP/data"
	"github.com/heyjoakim/AARP/helpers"
	"github.com/heyjoakim/AARP/logic"
	"github.com/heyjoakim/AARP/set"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	root := "/home/hey/git/Zeeguu-API"

	extension := ".py"
	exclusion := "__init__.py"
	fileSlice := logic.WalkDirs(root, extension, exclusion)
	helpers.CountFiles(fileSlice)

	f, err := os.Create("data.wsd")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("@startuml\ntitle Automated Diagram\nskinparam nodesep 100\nskinparam ranksep 100\n")
	dependencies := make(map[string]bool)
	nodes := make(map[string]bool)

	for _, j := range *fileSlice {
		file, imports := helpers.ReadLines(j)
		dependencies, nodes = set.WriteToSet(dependencies, nodes, file, imports)
	}
	data.WriteFile(f, nodes)
	data.WriteFile(f, dependencies)
	f.WriteString("@enduml")
	defer f.Close()

}
