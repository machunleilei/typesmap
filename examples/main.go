package main

import (
	"fmt"

	"github.com/machunleilei/typesmap"
)

type T1 struct {
	A1 string
	A2 int
	A3 []byte
	A4 []string
}

type T2 struct {
	A1 map[string]string
	A2 T1
	A3 map[string][]string `xx:"-"`
	A4 []*T1
}

func main() {
	var xx = typesmap.NewSimpleKvExtractor("xx")
	err := xx.Put("t1", &T1{
		A1: "hello",
		A2: 3,
		A3: []byte("world"),
		A4: []string{"xxx", "yyy"},
	})
	fmt.Println(err)
	t2 := &T2{
		A1: map[string]string{
			"a1": "b1", "a2": "b2",
		},
		A2: T1{
			A1: "hello",
			A2: 3,
			A3: []byte("world"),
			A4: []string{"xxx", "yyy"},
		},
		A3: map[string][]string{
			"a1": []string{"x", "y", "z"},
		},
	}
	t2.A4 = append(t2.A4, &t2.A2)
	t2.A4 = append(t2.A4, &t2.A2)
	err = xx.Put("t2", t2)

	fmt.Println(err)
	fmt.Println(xx.GetBytes("t1.A3"))
	fmt.Println(xx.GetBytes("t2.A2.A3"))
	fmt.Println(xx.GetStrings("t2.A2.A4"))
	fmt.Println(xx.GetStrings("t2.A3.a1"))
	aa, _ := xx.GetKvExtractors("t2.A4")
	fmt.Println(aa[0].GetString("A1"))
	fmt.Println(aa[0].GetStrings("A4"))
}
