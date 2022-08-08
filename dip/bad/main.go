// Bad pattern

package main

import "fmt"

type MySQL struct {
}

func (ms MySQL) QuerySomeData() []string {
	return []string{"info1", "info2", "info3"}
}

type MyRepository struct {
	db MySQL
}

func (mr MyRepository) GetData() []string {
	return mr.db.QuerySomeData()
}

func main() {
	mysqlDB := MySQL{}
	repo := MyRepository{db: mysqlDB}
	fmt.Println(repo.GetData())
}