package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Leon")
	data.PushBack("Siahaan")
	data.PushBack("Hinalang")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) //Leon

	next := head.Next() //Siahaan
	fmt.Println(next.Value)

	next = next.Next() //Siahaan
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
