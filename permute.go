package main

import (
	"fmt"
)

func main() {
	var in string
	fmt.Scanf("%s", &in)
	out := make([]byte, 0)
	l := len(in)
	used := make([]bool, l)
	permute(in, out, used, l, 0)
}

func permute(in string, out []byte, used []bool, l int, lvl int) {
	if lvl == l {
		fmt.Println(string(out))
		return
	}
	for i := 0; i < l; i++ {
		if used[i] == true {
			continue
		}
		out = append(out, in[i])
		used[i] = true
		permute(in, out, used, l, lvl+1)
		used[i] = false
		out = out[:lvl]
	}

}
