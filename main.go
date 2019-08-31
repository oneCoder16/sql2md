package main

import (
	"flag"
	"io/ioutil"
	"sql2md/service"
	"sql2md/util"
)

var (
	sqlPath      string
	markdownPath string
)

func init() {
	flag.StringVar(&sqlPath, "sql", "", "sql file path")
	flag.StringVar(&markdownPath, "md", "", "markdown file path")
}

func main() {
	flag.Parse()

	if sqlPath == "" || markdownPath == "" {
		panic("sql or md is empty !")
	}

	content, err := util.GetContentByFile(sqlPath)
	if err != nil {
		panic(err)
	}

	parse := service.ParseService{}
	transfer := service.TransferService{}

	tables := parse.Str2Sql(content)
	str := transfer.Sql2md(tables)

	err = ioutil.WriteFile(markdownPath, []byte(str), 0666)
	if err != nil {
		panic(err)
	}
}
