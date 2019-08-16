package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sql2md/pkg/sql"
)

func main() {
	sql_file := "./tb_channel.sql"

	file, err := os.Open(sql_file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

}
