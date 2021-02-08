package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "n sei"
	}
}

func main() {
	fmt.Println(tipo(time.Now()))
}