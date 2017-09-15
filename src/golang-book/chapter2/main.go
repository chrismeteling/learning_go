package main
import "fmt"


func main() {
	var rx1, ry1 float64 = 0, 10
	fmt.Println(rx1)
	fmt.Println(ry1)
}
func olla(p int, s string) (int, string){
	return p, "12"
}
func zero(xPtr *int) {
	*xPtr = 0
	fmt.Println(xPtr)
}
func noPointer(x int) {
	fmt.Println(&x,"id")
}
func catch(){
	fmt.Println(recover())
}
func giveP(w string) (p string){
	p = "i am returned automatically"+w
	return
}

func findLowestNumber()int {
	x := []int{
		48, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 9, 17,
	}

	if (len(x) >2) {
		var safe int
		safe = x[0]
		for _, value := range x {
			if (safe < value) {
				safe = value
			}
		}
		return safe
	}
	return 0
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