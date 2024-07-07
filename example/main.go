package main

import (
	"fmt"

	"github.com/nghtf/utils"
)

func main() {

	size, _ := utils.GetFileSize("./main.go")
	fmt.Println(utils.SizeHumanized(size))

}
