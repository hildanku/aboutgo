package main

import "fmt"

func main() {
	CheckScore(50)
}

func CheckScore(score int) {
	if score > 80 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Gagal")
	}
}

func CheckScoreWithSwitch() {

}
