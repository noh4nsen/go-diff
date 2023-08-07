package main

import (
	"fmt"
	"go-diff/helper"
	"go-diff/service"
	"go-diff/validator"
	"os"
)

func main() {
	validator.Validate(os.Args)
	fmt.Println(helper.EncodeToJson(service.ExecDiff(os.Args)))
}
