package main
import "fmt"

func main()  {
	var w int=2
	fmt.Println(string(w))
	a:=fahrenheitToCelsius(41)
	fmt.Println(a)
	//dividibleThrew3()

}

func fahrenheitToCelsius( grad int) int{
	grad=(grad-32)*5/9
	return grad
}

func countTo10(){
	fmt.Println(`1 2 3 4 5 6 7 8 9 10
11`)
}

func forLoop1(){
	i := 1
	for i <= 11 {
		fmt.Println(i)
		i++
	}
	if i%2 ==0 {
	}else if i==2 {
	}else if true&&true{}
}

func foorLoop2()  {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func dividibleThrew3()  {
	for i := 3; i <= 100; i=i+3 {
		fmt.Println(i)
	}

}