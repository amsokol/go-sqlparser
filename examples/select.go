package main

import (
	"log"

	"github.com/amsokol/go-sqlparser"
)

func main() {
	tree, err := sqlparser.Parse("select * from db1.table1")
	fatalOnError(err)

	log.Println("Tree:", sqlparser.String(tree))
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
