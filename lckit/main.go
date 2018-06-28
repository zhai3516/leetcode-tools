package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	TreeTemplate = []string{
		"package main",
		"\n",
		"\nimport (",
		"\n\t\"fmt\"",
		"\n",
		"\n\t\"github.com/zhai3516/leetcode-tools/model\"",
		"\n\t\"github.com/zhai3516/leetcode-tools/tools\"",
		"\n)",
		"\n",
		"\nfunc main() {",
		"\n\tvar treeList []interface{}",
		"\n\tvar root *model.TreeNode",
		"\n\tvar result []int",
		"\n",
		"\n\ttreeList = []interface{}{5, 4, 7, 3, nil, 2, nil, -1, nil, 9}",
		"\n\troot = tools.Slice2TreeNodes(treeList)",
		"\n}",
	}

	TmpMap = map[string][]string{
		"tree": TreeTemplate,
	}
)

func CreateTemplate(name, tmpType string) {
	if TmpMap[tmpType] == nil {
		panic(errors.New("Unsopport template type"))
	}

	dir := strings.Replace(name, " ", "", -1)

	var err error
	err = os.Mkdir(dir, os.ModePerm)
	if err != nil && os.IsNotExist(err) {
		panic(err)
	}

	var fd *os.File
	fd, err = os.Create(dir + "/main.go")
	if err != nil && os.IsNotExist(err) {
		panic(err)
	}

	for _, line := range TmpMap[tmpType] {
		fd.Write([]byte(line))
	}

	fd.Close()
	fmt.Println("Success !")
}

func main() {
	var flagSet = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var kit = flagSet.String("kit", "tmp", "Chose a kit, support `tmp`")
	var fileName = flagSet.String("name", "lckit-create", "Template dir name")
	var fileTyep = flagSet.String("type", "tree", "Template type, support `tree` only")
	flagSet.Parse(os.Args[1:])
	if *kit == "tmp" {
		CreateTemplate(*fileName, *fileTyep)
	} else {
		fmt.Println("Unsupport kit")
	}
}
