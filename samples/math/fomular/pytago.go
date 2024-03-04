package fomular

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/math"
)

func SampleMathPytago() {
	println("Pytago of number (integer): ", fmt.Sprintf("%f", math.Pytago(3, 4)))
	println("Pytago of number (float): ", fmt.Sprintf("%f", math.PytagoF(3.0, 4.0)))
}
