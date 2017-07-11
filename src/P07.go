package exercise

import "fmt"

func P07(x interface{}, y interface{}, z interface{}) string {
	return fmt.Sprintf("%v時の%vは%v", x, y, z)
}
