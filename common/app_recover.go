package common

import "fmt"

func AppRecover() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}