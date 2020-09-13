package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate sqlboiler --wipe mysql

func main() {
	fmt.Println("vim-go")
}
