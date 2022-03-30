package typesmap

import (
	"testing"
)

type T_T1 struct {
	A1 string
	A2 int
	A3 []byte
	A4 []string
	A5 string `xx:"-"`
}

type T_T2 struct {
	A1 map[string]string
	A2 T_T1
	A3 map[string][]string
	A4 []*T_T1
}

func Test_kvExtractor(t *testing.T) {
	var xx = NewSimpleKvExtractor("xx")
	err := xx.Put("t1", &T_T1{
		A1: "hello",
		A2: 3,
		A3: []byte("world"),
		A4: []string{"xxx", "yyy"},
	})
	if err != nil {
		t.Errorf("Put T_T1 fail")
	}
	t2 := &T_T2{
		A1: map[string]string{
			"a1": "b1", "a2": "b2",
		},
		A2: T_T1{
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
	if err != nil {
		t.Errorf("Put T_T2 fail")
	}
	bytes, err := xx.GetBytes("t1.A3")
	if err != nil {
		t.Errorf("get t1.A3 fail")
	}
	if string(bytes) != "world" {
		t.Errorf("t1.A3 not equal")
	}
	bytes, err = xx.GetBytes("t2.A2.A3")
	if err != nil {
		t.Errorf("get t1.A2.A3 fail")
	}
	if string(bytes) != "world" {
		t.Errorf("t1.A2.A3 not equal")
	}
	strings, err := xx.GetStrings("t2.A2.A4")
	if err != nil {
		t.Errorf("get t2.A2.A4 fail")
	}
	if !compareStrings(strings, []string{"xxx", "yyy"}) {
		t.Errorf("check t2.A2.A4 not equal")
	}
	strings, err = xx.GetStrings("t2.A3.a1")
	if err != nil {
		t.Errorf("get t2.A3.a1 fail")
	}
	if !compareStrings(strings, []string{"x", "y", "z"}) {
		t.Errorf("check t2.A3.a1 not equal")
	}
	aa, err := xx.GetKvExtractors("t2.A4")
	if err != nil {
		t.Errorf("get t2.A4 fail")
	}
	ss, err := aa[0].GetString("A1")
	if err != nil {
		t.Errorf("get A1 fail")
	}
	if ss != "hello" {
		t.Errorf("A1 not equal")
	}
	strings, err = xx.GetStrings("t2.A2.A4")
	if err != nil {
		t.Errorf("get A4 fail")
	}
	if !compareStrings(strings, []string{"xxx", "yyy"}) {
		t.Errorf("check A4 not equal")
	}
	if _, err = xx.GetString("t1.A5"); err == nil {
		t.Errorf("t1.A5 should not get")
	}
}

func compareStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}
