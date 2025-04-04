package main
import ("fmt"
		"math"	
			)

func aritmatikaModular(a, b float64)int{
	var hasil, i float64
	hasil = math.Pow(a, b)
	i = hasil
	for int(i)%100 != 0{
		i-=1
	}
	return int(hasil)%int(i)
}
func main(){
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Print(aritmatikaModular(a, b))
}
