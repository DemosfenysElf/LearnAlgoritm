package main

import (
	"bufio"
	"fmt"
	"os"
)

type skobki struct {
	i      int
	skobka string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	s1 := split1(text)
	//

	var mFullSkobok []skobki
	var mDropSkobok []skobki
	for i, s := range s1 {
		if (s == "(") || (s == ")") || (s == "[") || (s == "]") || (s == "{") || (s == "}") {
			var sk1 skobki
			sk1.skobka = s
			sk1.i = i
			mDropSkobok = append(mDropSkobok, sk1)
			mFullSkobok = append(mFullSkobok, sk1)
		}
	}

	//
	if len(mDropSkobok) != 0 {
		ok := trueSk(mFullSkobok, mDropSkobok)
		if ok == 0 {
			fmt.Print("Success")
		} else {
			fmt.Print(ok)
		}
	} else {
		fmt.Print("Success")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func split1(s string) []string {
	chunks := make([]string, len(s))
	for i := range chunks {
		to := min(len(s), i+1)
		chunks[i] = s[i:to]
	}
	return chunks
}

func trueSk(sFull, sDrop []skobki) int {

	var iS int
	for i, i2 := range sFull {
		switch i2.skobka {
		case "(":
			iS++
		case "{":
			iS++
		case "[":
			iS++
		case ")":
			if iS == 0 {
				return sFull[i].i + 1
			}
			if sDrop[iS-1].skobka == "(" {
				if i != len(sFull) {
					sDrop = append(sDrop[:iS-1], sDrop[iS+1:]...)
					iS--
				} else {
					sDrop = sDrop[:iS-1]
				}
			} else {
				return sFull[i].i + 1
			}

		case "}":
			if iS == 0 {
				return sFull[i].i + 1
			}
			if sDrop[iS-1].skobka == "{" {

				if i != len(sFull) {
					sDrop = append(sDrop[:iS-1], sDrop[iS+1:]...)
					iS--
				} else {
					sDrop = sDrop[:iS-1]
				}
			} else {
				return sFull[i].i + 1
			}
		case "]":
			if iS == 0 {
				return sFull[i].i + 1
			}
			if sDrop[iS-1].skobka == "[" {

				if i != len(sFull) {
					sDrop = append(sDrop[:iS-1], sDrop[iS+1:]...)
					iS--
				} else {
					sDrop = sDrop[:iS-1]
				}
			} else {
				return sFull[i].i + 1
			}
		default:

		}
	}
	if len(sDrop) != 0 {
		return sDrop[len(sDrop)-1].i + 1
	}

	return 0
}
