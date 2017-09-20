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
   
    var rx1, ry1 float64 = 0, 10 //xy1=0.0  ry1=10.0
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
You cannot randomly fuse/add different types, always Typecast
  
##Control Structures
For Loop

    for{}// never ending for
    i:=1
    for i<=12 {i++}  //uses previous i
    for i:=1; i<=10; i++ {} //new intern i
    for i=1; i<=10; i++ {} //uses previous i 
    
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
Example, special For-loop 1

    for i, value := range x {}
    Goes threw the array x, value is x[i]; i is the iterator | you need to use i
Example, special For-loop 2

    for _, value := range x {}
    Goes threw the array x, value is x[i]; the iterator itself cant be used  
#####Slices
Arrays wihtout length bounding, they are based on finite arrays

    var x []float64             //slize and used array is not set
    y := make([]float64, 5)     //current-slice-length5; array-length 5
    z := make([]float64, 5, 10) //current-slice-length5; array-length 10 
    arr := [5]float64{1,2,3,4,5}
    a := arr[0:5]               //is arr, from arr[0] to arr[5]
    b := arr [:]               //is arr from arr[0] to arr[len(arr)]    
Slice functions

    slice1:= []int{1,2,3}       //{1 2 3}
    slice2:= append(slice1,4,5) //{1 2 3 4 5}
    append uses slices and other arguments
<span><span>

    slice3:= make([]int,2)      //{0 0}
    copy (slice3,slice1)        //slice2 {1 2}
    Copyies 2nd argument in 1st argument, attention if space is limited
#####Maps
Maps are a Hash of Values   
The length of maps is infinite

        var x =make(map[string]int) //maps Strings to ints
        y := make(map[int]bool)     //map for ints to booleans
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
    
    func olla(p int, s String) (int, String){
    return p, "12"
    }
    func hello(thisAInt, thisAnotherInt int){
    fmt.Println(thisAInt, thisAnotherInt)
    }
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
##Structs and Interfaces
###Structs    
Structs can be used to create a new customized type
#####Creation

    type Random struct {
    d int
    a float64
    b String
    }
    type Circle struct {
    x, y, r float64 
    }
#### Initialization
You can initialize a struct like any other type, but to enter values is a bit different:

    var c Circle //x=0.0 y=0.0 r =0.0
    d:= new(Circle) //returns a Pointer
    e:= Circle{x: 0, y:0, r:5}  
    f:= Circle{0,0,5}
####Access the fields

    c.x     //=0
    c.r     //5
####Use in Functions
One thing to remember is that arguments are always copied in Go. If we attempted to modify one of the fields inside of another function, it would not modify the original variable.    
If you want to modify the value use Pointers, otherwise it will be copied

    c := Circle{0, 0, 5}
    fmt.Println(circleArea(&c)) //give the function circleArea a Pointer
<span></span>
    
    func circleArea(c *Circle) float64 { //take a Pointer (*Circle)
      return math.Pi * c.r*c.r  //works with the original Values
    }
### Methods
Methods are functions wich are defined for a special type. So you can only call a method by using a variable of the specified type.
If Using a Pointer it might be useful to use Methods instead of functions       
Example for creating a method for the type Circle

    func (c *Circle) area() float64 { //this is a method by Circle. func (Calling Object) Methodname(Input) Output
      return math.Pi * c.r*c.r
    }
You can call the Method by this: This will work with the calling Circle


    c.area()    //c is an instance of Circle
    //The Pointer is automatically created by Go
####Embedded Types
Using customized types in customized types, this is our inner Type:

    type Person struct {
    Name string
    }
Has Relationship: Android has a Person

    type Android struct {
      PersonIdentifier Person //call Person by .PersonIdentifier
      Model string     //cal Person vars .PersonIdentifier.Name
    }
    
Is Relationship: Android is a Person

    type Android struct {
      Person        //call Person by .Person 
      Model string  //call Person vars by .Person.Name
     }              //or directly .Name  since Android is a Person
Use the "Is Relationship" to create OO Programms, calll methhouds my Android.methodOfPerson()
###Interfaces
Create an Interface, interfaces contain a number of methods, not vars (like structs)

    type Shape interface {
      area() float64        //all objects of shape have to contain the method area()
    }
Interfaces can be used like a placeholder for any fitting object, therefore it can even be used as a field 

          type MultiShape struct { 
            shapes []Shape      //the interface is used as a field in a struct
          }
##fmt
Read input

    var input TYPE
    fmt.Scan("%ABBREVIATION", input)