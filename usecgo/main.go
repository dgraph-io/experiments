package main

// #cgo LDFLAGS: -lm
// #include <math.h>
//
// double ps(double a, double b) {
// 	double result = pow(a, b);
//  result = sqrt(result);
//  return result;
// }
import "C"
import "fmt"

func Pow(b, e float64) float64 {
	return float64(C.pow(C.double(b), C.double(e)))
}

func Sqrt(b float64) float64 {
	return float64(C.sqrt(C.double(b)))
}

func main() {
	b, e := 5.0, 2.0
	fmt.Println("5 ^ 2:", Pow(b, e))
	fmt.Println("Sqrt of 5:", Sqrt(b))
	fmt.Println("sq:", C.ps(C.double(b), C.double(e)))
}
