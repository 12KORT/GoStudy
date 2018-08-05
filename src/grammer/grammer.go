package grammer

import (
	"fmt"
	"strconv"
)

func Slice_use() {
	s0 := make([]string, 3)
	fmt.Println("empty s0: ", s0)

	s0[0] = "string0"
	s0[1] = "string1"
	s0[2] = "string2"
	fmt.Println("init s0", s0)

	s0 = append(s0, "string3")
	s0 = append(s0, "string4")
	fmt.Println("after append", s0)

	s1 := make([]string, len(s0))
	copy(s1, s0)
	fmt.Println("s1: ", s1)

	s2 := s1[2:4]
	fmt.Println("s2: ", s2)

	s3 := make([][]string, 3)
	for i := 0; i < 3; i++ {
		len := i + 1
		s3[i] = make([]string, len)
		for j := 0; j < len; j++ {
			s3[i][j] = strconv.Itoa(i) + strconv.Itoa(j)
		}
	}

	fmt.Println("s3:", s3)
}

func Map_Use() {
	m := make(map[string]int)
	m["v1"] = 1
	m["v2"] = 2
	fmt.Println("map m:", m)

	v1 := m["v1"]
	fmt.Println("map v1: ", v1)

	delete(m, "v1")
	fmt.Println("map m:", m)

	_, prs1 := m["v1"] //第一个值为是否在m里，返回false
	fmt.Println("prs1", prs1)

	v2, prs2 := m["v2"] //v2存在，返回true
	fmt.Println("v2, prs2:", v2, prs2)
}

func VarArgs(str string, nums ...int) {
	fmt.Println("nums:", nums)

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println("total :", total)
}
