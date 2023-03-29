package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	var peringkat string
	nilai := (math + science + english + indonesia) / 4
	if nilai == 100 {
		peringkat = "Sempurna"
	} else if nilai >= 90 {
		peringkat = "Sangat Baik"
	} else if nilai >= 80 {
		peringkat = "Baik"
	} else if nilai >= 70 {
		peringkat = "Cukup"
	} else if nilai >= 60 {
		peringkat = "Kurang"
	} else if nilai < 60 {
		peringkat = "Sangat kurang"
	} else {
		peringkat = "Nilai tidak boleh lebih dari 100 dan kurang dari 0"
	}
	return peringkat // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
