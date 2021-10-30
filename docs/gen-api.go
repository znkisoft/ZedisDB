package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type Field struct {
	FieldName string    `json:"fieldName"`
	Commands  []Command `json:"commands"`
}

type Command struct {
	Name          string `json:"name"`
	IsImplemented bool   `json:"isImplemented"`
}

var jsonFilePath = "./docs/commands.json"
var templateFilePath = "./docs/template.go.tpl"

func main() {
	source, _ := os.ReadFile("./README.md")

	root := goldmark.DefaultParser().Parse(text.NewReader(source))

	fields := deserialize()

	err := ast.Walk(root, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			// print n type
			switch n.(type) {
			case *ast.HTMLBlock:

				// find the anchor block
				if n.(*ast.HTMLBlock).HTMLBlockType == 2 {
					// append(n.(*ast.HTMLBlock).Text())pend the docs to the anchor block

					block := generateDocTable(fields)

					if n.HasChildren() {
						n.ReplaceChild(n, n.LastChild(), block)
					} else {
						n.AppendChild(n, block)
					}

				}
			default:
				// fmt.Printf("%+v\n", n)
			}

		}
		return ast.WalkContinue, nil
	})

	if err != nil {
		log.Fatal("render md file error: ", err)
	}

	// output, err := os.Create("./test.md")
	// if err != nil {
	// 	log.Fatal("create file error:", err)
	// }
	// defer output.Close()

	fmt.Print(string(root.Text(source)))

	// goldmark.DefaultRenderer().Render(os.Stdout, source, root)
}

func deserialize() []Field {
	buf, err := os.ReadFile(jsonFilePath)
	if err != nil {
		panic(err)
	}

	fields := []Field{}
	err = json.Unmarshal(buf, &fields)

	if err != nil {
		panic(err)
	}

	return fields
}
