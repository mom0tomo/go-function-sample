package main

import (
	"fmt"
	"math"
)

// Goでは引数にデフォルト値は与えられない！！
func powerCharge(wat float64, h float64, rt float64) {

	kw := wat / 1000
	c := math.Floor(kw * h * rt)
	
	fmt.Println("e charge: ", c, "en")
	return
}

func main() {

	// 蛍光灯(北海道電力)
	powerCharge(60, 24, 25)
	// 蛍光灯(北陸電力)
	powerCharge(60, 24, 20)

	// ドライヤー
	powerCharge(1200, 5, 25)
}
