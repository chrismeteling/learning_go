#Learn Go

##Variables
#####Variable Types

    bool                                    //true or false; use && || !
    String                                  //typical String
    uint int   //use this if possible, without number                      
    int8 byte; int16; int32 rune ; int64    //signed integer: +-
    uint8; uint16; uint32 ; uint64          //unsigned integer: +  
    float32; float64                        //float numbers : +- ,
    complex32; complex63                    //imaginary numbers
    NaN                                     //not a number, eg. 0/0; ∞; -∞
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
##fmt
Read input

    var input TYPE
    fmt.Scan