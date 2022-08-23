package main

import (
	"fmt"
	"github.com/veerakumarak/go_validator"
)

type Test struct{
	Var8 int8
	Var16 int16
	Var32 int32
}

func main() {
	errs := go_validator.Int8(49).Gt(49).Lt(49).Validate()
	if errs != nil {
		fmt.Println(errs)
	} else {
		fmt.Println("valid")
	}

	errs = go_validator.Bool(true).False().Validate()
	if errs != nil {
		fmt.Println(errs)
	} else {
		fmt.Println("valid")
	}

	errs = go_validator.String("veer").NotEmpty().Max(3).Min(2).Validate()
	if errs != nil {
		fmt.Println(errs)
	} else {
		fmt.Println("valid")
	}

	t := Test{Var8: 10, Var16: 200}

	errs = go_validator.Struct(&t,
		go_validator.Int8(t.Var8).Gt(10).Lte(127),
		).Validate()
	if errs != nil {
		fmt.Println(errs)
	} else {
		fmt.Println("valid")
	}

}