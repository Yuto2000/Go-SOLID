// Bad pattern

// package main

// import "fmt"

// type MySQL struct {
// }

// func (ms MySQL) QuerySomeData() []string {
// 	return []string{"info1", "info2", "info3"}
// }

// type MyRepository struct {
// 	db MySQL
// }

// func (mr MyRepository) GetData() []string {
// 	return mr.db.QuerySomeData()
// }

// func main() {
// 	mysqlDB := MySQL{}
// 	repo := MyRepository{db: mysqlDB}
// 	fmt.Println(repo.GetData())
// }

// Good pattern

package main

import "fmt"

type DBconn interface {
	QuerySomeData() []string
}

type MySQL struct {
}

func (ms MySQL) QuerySomeData() []string {
	return []string{"info1", "info2", "info3"}
}

type MyRepository struct {
	db DBconn
}

func (mr MyRepository) GetData() []string {
	return mr.db.QuerySomeData()
}

func main() {
	mysqlDB := MySQL{}
	repo := MyRepository{db: mysqlDB}
	fmt.Println(repo.GetData())
}