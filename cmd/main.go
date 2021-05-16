package main

import (
	"os"

	"github.com/heyjoakim/AARP/data"
	"github.com/heyjoakim/AARP/helpers"
	"github.com/heyjoakim/AARP/logic"
	"github.com/heyjoakim/AARP/set"
	log "github.com/sirupsen/logrus"
)

func main() {
	root := "/home/hey/git/Zeeguu-API"
	extension := ".py"
	exclusion := "__init__.py"

	files := logic.WalkDirs(root, extension, exclusion)
	helpers.CountFiles(files)

	f, err := os.Create("out/data.wsd")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("@startuml\ntitle Automated Diagram\nskinparam nodesep 100\nskinparam ranksep 100\n")

	dependencies := make(map[string]bool)
	nodes := make(map[string]bool)

	for _, j := range *files {
		f, imports := logic.ReadLines(j)
		dependencies, nodes = set.WriteToSet(dependencies, nodes, f, imports)
	}
	data.WriteFile(f, nodes)
	data.WriteFile(f, dependencies)
	f.WriteString("@enduml")
	defer f.Close()

	// Requires plantuml binary in cmd folder
	// cmd := exec.Command("./plantuml", "out/data.wsd")
	// cmd.Run()

	// if err != nil {
	// 	log.Fatal(err)
	// }
}
