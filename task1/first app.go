package main
import (
	"fmt"
	"math"
)
func L (r float64) float64 {
	return r * math.Pi * 2
}
func S (r float64) float64 {
	return math.Pow(r,2.0) * math.Pi
}
func V (r float64) float64 {
	return math.Pow(r,3.0) * math.Pi * 0.75
}

func main() {
	var funcs[3] string = [3]string{"L","S","V"}
	r := 10.0
	for i := 0; i < len(funcs); i++ {
		fmt.Print(funcs[i]," = ")
		if funcs[i] == "L" {
			fmt.Println(L(r))
		}
		if funcs[i] == "S" {
			fmt.Println(S(r))
		}
		if funcs[i] == "V" {
			fmt.Println(V(r))
		}
	}
}

