package main

import (
	"fmt"
	"os"
)

type massiveVershine struct {
	i     int
	versh int
}

func main() {
	var n, m int
	var masM []massiveVershine
	var M massiveVershine
	fmt.Fscan(os.Stdin, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &m)
		M.i = i
		M.versh = m
		masM = append(masM, M)
	}

	fmt.Println(lvlDreva(masM))
}

func lvlDreva(masM []massiveVershine) (lvl int) {
	var lvlVse []int
	var lvlTek []int

	n := len(masM)
	lvlVse = append(lvlVse, -1)
	for l := 0; l < n; l++ {
		for j := 0; j < len(lvlVse); j++ {

			for i := 0; i < len(masM); i++ {

				if masM[i].versh == lvlVse[j] {
					lvlTek = append(lvlTek, masM[i].i)
					masM = append(masM[:i], masM[i+1:]...)
					i--
				}
			}
		}
		if lvlTek != nil {
			lvl++
		}
		lvlVse = append(lvlVse[len(lvlVse):], lvlTek...)
		lvlTek = nil
	}

	return
}
