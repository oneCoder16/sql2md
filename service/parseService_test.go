package service

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test2Sql(t *testing.T) {
	sql_file := "../tb_channel.sql"

	file, err := os.Open(sql_file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	srv := ParseService{}

	for _, table := range srv.Str2Sql(data) {
		t.Logf("%+v\n", table)
	}
}
