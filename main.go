package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vaibhav/practice/projectpharmamanagementsystem/Config"
	"github.com/vaibhav/practice/projectpharmamanagementsystem/Routes"
)

func main() {
	Config.DbSetup()
	Routes.MakeRouter()
}
