// Package main
// Time : 2023/8/6 12:41
// Author : PTJ
// File : main
// Software: GoLand
package main

import (
	_ "user/conf"
	"user/model"
)

func main() {
	model.InitDB()
}
