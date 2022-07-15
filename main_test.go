package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

func BenchmarkUnmarshal(b *testing.B) {
	p := &Person{
		Name:  "Gopher",
		Age:   30,
		Color: "black",
		PoB:   "China",
	}
	pJSON, _ := json.Marshal(p)
	for n := 0; n < b.N; n++ {
		stu := &Person{}
		json.Unmarshal(pJSON, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	var pool = sync.Pool{
		New: func() any {
			fmt.Println("Creating a new person...")
			return &Person{}
		},
	}
	p := &Person{
		Name:  "Gopher",
		Age:   30,
		Color: "black",
		PoB:   "China",
	}
	pJSON, _ := json.Marshal(p)

	for n := 0; n < b.N; n++ {
		stu := pool.Get().(*Person)
		json.Unmarshal(pJSON, stu)
		pool.Put(stu)
	}
}
