package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)
//читает файл, сплитим по \n, до последнего чтобы ентер, лупы, принимает флоты64 и пишем 64 чтобы он работал и апендим "f" в "а"
func main() {
	args := os.Args[1:]
	data, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Println("ERROR")
		return
	}
	s := strings.Split(string(data), "\n")
	s = s[:len(s)-1]
	var a []float64
	var f float64
	for i := 0; i < len(s); i++ {
		f, err = strconv.ParseFloat(s[i], 64)
		if err != nil {
			log.Println("ERROR")
			return

		}
		a = append(a, f)
	}
	//вызываем свои функции, чтобы округлить используем инт 64 мастраунд,
	average := int64(math.Round(average(a)))
	fmt.Println("Average:", average)
	median := int64(math.Round(median(a)))
	fmt.Println("Median:", median)
	variance := int64(math.Round(variance(a)))
	fmt.Println("Variance:", variance)
	stanDev := int64(math.Round(standardDeviation(a)))
	fmt.Println("Standard Deviation:", stanDev)
}
//сумма всех элементов делим на количество всех чисел
func average(f []float64) float64 {
	var ress float64
	var res float64
	for i := 0; i < len(f); i++ {
		ress += f[i]
	}
	res = ress / float64(len(f))
	return res
}
//все элементы минус средняарф. в квадрате
func variance(arr []float64) float64 {
	var res []float64
	for _, w := range arr {
		res = append(res, math.Pow(w-average(arr), 2))
	}
	return average(res)
}
//variance под корнем
func standardDeviation(arr []float64) float64 {
	return math.Sqrt(variance(arr))
}
//бабл сорт два цикла первый от 0 до предпоследнего, второй после нуля до последнего если первый желмент больше вторго они меняются местами
func median(f []float64) float64 {
	var res float64
	for i := 0; i < len(f)-1; i++ {
		for j := i + 1; j < len(f); j++ {
			if f[i] > f[j] {
				f[i], f[j] = f[j], f[i]
			}
		}
	}
	if len(f) == 0 {
		fmt.Println("0")
	} else if len(f)%2 == 0 {//если чётное 
		res = (f[len(f)/2-1] + f[len(f)/2]) / 2
	} else {
		res = f[len(f)/2]//если нечётное 
	}
	return res
}
