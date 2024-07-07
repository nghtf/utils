package main

import (
	"fmt"
	"utils"
)

func main() {

	size, _ := utils.GetFileSize("./main.go")
	fmt.Println(utils.SizeHumanized(size))

}
