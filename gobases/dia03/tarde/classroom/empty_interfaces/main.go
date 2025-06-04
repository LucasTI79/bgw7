package main

import "fmt"

type SliceAny struct {
	Data []any
}

func main() {
	l := SliceAny{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "hello")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)
}

// Data [1,2,3]
// Data [1,"hello",true]

// var myVariable int
// var anyType any
// anyType = myVariable

// map[string]any
// {
// 	status: 200,
// 	timestamp: 12312321312,
// 	data: {
// 	}
// }
