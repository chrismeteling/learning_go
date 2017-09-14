#Learn Go

##Variables
#####Variable Types
Always Typecast if using different types! 

    bool        //true or false; use && || !
    String      //typical String
    uint        //positive, as many bits as the system
    int         //positive negative, as many bits as the system
    float32     //float numbers : +- ,32bit
    float64     //float numbers : +- ,64bit
    complex32   //imaginary numbers, 32bit
    complex64   //imaginary numbers, 64bit
    NaN          //not a number, eg. 0/0; ∞; -∞
    int8; int16; int32=rune ; int64    //signed integer: +- 
    uint8=byte; uint16; uint32; uint64          //unsigned integer: +  
#####Create a Variable (String)

    var x string 
    x="yoo"
    var y string="joo"
    var z = "jojojoguhrt"
    a := "boi"
Create multiple Variables at once
   
    var(                //multiple vars at once
        a="hello world" //automatic typecast possible
        b=-12           //multiple types possible too
        c uint16=12     //signed type possible too
    )

#####Create a constant
    
    const x float32 = 3.14
    const y = "Hello World"
    
#####Type Cast

    Type(value)
    string(65)
#####Danger on Types
You cannot randomly fuse/add different types:

    These do not combine:
    float ! int ! string ! bol    
##Control Structures
For Loop

    i:=1
    for i<=12 {i++}
    for i:=1; i<=10; i++ {}
    
If Case
    
        if i%2 ==0 {
        }else if i==2 {
        }else if true||false {}
Switch

    switch i {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
##Arrays

    var x [25]int
    y := [5]float64{ 98, 93, 77, 82, 83, } // the , after 83 is for easy adding other values
    y[3}=10          // [98 93 77 10 83] 
    len(x)          
Example, special for loop

    for i, value := range x {}
    Goes threw it, value is x[i]; i is the iterator | you need to use i
Example, special for loop

    for _, value := range x {}
    Goes threw it, value is x[i]; the iterator itself cant be used  
#####Slices
Arrays wihtout length bounding, they are based on a finite array

    var x []float64             //current-slice-length; array-length 0
    y := make([]float64, 5)     //current-slice-length5; array-length 5
    z := make([]float64, 5, 10) //current-slice-length5; array-length 10 
    arr := [5]float64{1,2,3,4,5}
    a := arr[0:5]               //is arr, from arr[0] to arr[5]
    b := arr [:]               //is arr from arr[0] to arr[len(arr)]    
Slice functions

    slice1:= []int{1,2,3}       //{1 2 3}
    slice2:= append(slice1,4,5) //{1 2 3 4 5}
    take slice and other arguments
<span><span>

    slice3:= make([]int,2)      //{0 0}
    copy (slice3,slice1)        //slice2 {1 2}
    Copyies 2nd argument in 1st argument, attention if space is limited
#####Maps
Maps are a Hash of Values, x is a map of `string` to  `ints`
"key"] = 10    
The length is not static

        var x =make(map[string]int)
        y := make(map[int]bool)
        x["key"]=10
        fmt.Println(x["key"])   //10
        fmt.Println(x["noneExample"]) //returns the zero value (0)
        delete(x,1)
        elements := map[string]string{
          "H":  "Hydrogen",
          "He": "Helium",
        }
Map a Map with a Map

        eElements := map[string]map[string]string{
            "H": map[string]string{
              "name":"Hydrogen",
              "state":"gas",
            },
             "He": map[string]string{
              "name":"Helium",
              "state":"gas",
             },
         }
Test if none

    name, ok := elements["H"]  //ok: true; name: "Hydrogen"
    name, ok := elements["Un"] //ok: false; name: ""
    //oder elements["Un"]=="") ist true
###Test for Null

    stringcontainer :="w"
    var emptycontainer string
    variable1, variable2 := stringcontainer //w true
    variable1, variable2 := emptycontainer  //  false
##Functions
General:    
    
    func (p int, s String) (int, String){
    return p "12" }
Define a Variable to return

    func giveP() (p string){
        p = "i already exist"
        return
     }
Multiple arguments elemental type, array or slice  
    
    func test(args ...int) int{
Closure: Increment and x form the closure

    func test(){
        x:=0
        inside := func()int{
             x=x+1
            return x}
        fmt.Println(inside) //1
        fmt.Println(inside) //2
        
Objekt Orientiert (i guess)

    func callerObject() func() uint {
      i := uint(0)
      return func() uint {
        i += 2
        return i
    }
    func main(){
        myFunction:= callerObject()
        fmt.Println(myFunction) //2
        fmt.Println(myFunction) //4
        fmt.Println(myFunction) //6
        }
##Defer, Panic & Recover
###Defer
Use `dever` to execute a command at the end of the function.   
Defer is executed before the return is done   
And even if there is a panic

    func example(x int) (p int) {
    defer fml.Println(still no return)
    p=x+2
    return 
###Panic
Creates an Exception with a message, the programm is exited if not recover
    
    func main(){
        panic("Mean Messaeg")
        fmt.Println(recover()) //this is never reached
    }
###Recover
Recover stops the panic and catches the error, use defer to make use of it.   
The Recover-Defer has to be called before the panic starts.

    func main(){
        defer func(){
            fmt.Println(recover())
        }()
        panic("PAAAANIK")
    }
##Pointer
Pointers refer to memory location, use them to change the same location within different functions
create a Pointer (3 Ways, actually 2)

    zPtr := new(int)    //create a random Pointer to a memory space for a type
    var yPtr *string  //poits to a string 
    var x int   //create a random variable
    &x          //access the memory location vor the varialbe   
    xPtr := &x  //xPtr refers to the int memory location of x
<span><span>
    
    z := 12
    *xPtr=z    //the value of the memory location of x is set to 12
    --> x =12 
     zPtr= &z     //pointer points to memory of z
Use Pointers in functions:

    func zero(xPtr *int) {
      *xPtr = 0
    }
    func main() {
      x := 5
      zero(&x)
      fmt.Println(x) // x is 0
    }
##fmt
Read input

    var input TYPE
    fmt.Scan("%ABBREVIATION", input)