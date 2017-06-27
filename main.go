package main

import "fmt"


func powerCharge(wat float64, h float64) {

	var rt float64 = 25
	
	kw := wat / 1000
	c := kw * h * rt
	
	fmt.Println("e charge: ", c, "en")
	return
}

func main() {

	// 蛍光灯
	powerCharge(60, 24)

	// ドライヤー
	powerCharge(1200, 5)
}
