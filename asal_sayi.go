package main

import (
	"fmt"
)

func asal_sayi_mi(sayi int ) bool{
	if ( sayi < 2){ return false; }
	var max_kontrol_sayi int = sayi/2
	for index :=2; index <= max_kontrol_sayi; index++{
		if (sayi%index == 0){
			return false
		}
	}
	return true
}
func main(){
	fmt.Println("Starting..")
	for index := 1; index < 1000000; index++{
		sonuc := asal_sayi_mi(index)
		if sonuc{
			fmt.Println(index, sonuc)
		}
	}
	
}