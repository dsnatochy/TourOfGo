package main

import "golang.org/x/tour/pic"


func Pic(dx, dy int) [][] uint8{
	var pic [][]uint8
	for i:=0; i<dy; i++ {
		var line []uint8

		for j:=0; j< dx; j++{
			//rand.Seed(time.Now().Unix())
			line = append(line, uint8(i^j*j^i))
		}
		pic = append(pic, line)
	}
	return pic
}

func main() {
	//pic.Show(Pic)
}