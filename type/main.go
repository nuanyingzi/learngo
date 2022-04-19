package main
// 结构体
import "fmt"

// 类型定义
type NewInt int
// 类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a: %T\n", a)
	fmt.Printf("typr of b: %T\n", b)
	// type of a: main.NewInt
	// typr of b: int
}
