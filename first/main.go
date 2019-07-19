package main

import (
	"fmt"
	"os"
)

func main() {
	var name string

	name = os.Args[1]
	fmt.Println("Hello great", name, "!")

	name, age := "gandalf", 2019
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")
}

// func main() {
// 	fmt.Printf("%#v\n", os.Args)

// 	fmt.Println("Path:", os.Args[0])
// 	fmt.Println("1st argument:", os.Args[1])
// 	fmt.Println("2nd arguments:", os.Args[2])
// 	fmt.Println("3rd argument:", os.Args[3])

// 	fmt.Println("Number of items inside os.Args:", len(os.Args))
// }

// func main() {
// 	// var dir, file string
// 	_, file := path.Split("css/main.css")

// 	// fmt.Println("dir :", dir)
// 	fmt.Println("file:", file)
// }

// func main() {
// 	speed := 100
// 	force := 2.9

// 	speed = int(float64(speed) * force)
// 	fmt.Println(speed)
// 	fmt.Println(force, int(force))
// }

// func main() {
// 	var apple int
// 	var orange int32

// 	apple = int(orange)

// 	orange = 65
// 	color := string(orange)
// 	_ = apple
// 	fmt.Println(color)
// 	fmt.Println(string([]byte{104, 105}))
// }
