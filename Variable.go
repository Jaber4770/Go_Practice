package main
import ("fmt")


var z int
var y float32 = 3.3
var x bool = false
var zx string = "joe"



func main() {
	fmt.Println("Hello World! from Go")

	//Variable
	var studentName string = "Jaber Ahmed"
	var student1 string = "John"

	fmt.Println(studentName)
	fmt.Println(student1)


	// another way to write variable
	var nameOfDeveloper = "jack sargey"
	fmt.Println(nameOfDeveloper)

	// another way of varialbe

	developerNumber := 4
	fmt.Println(developerNumber);


	// another way or boolean
	var jack bool = true
	fmt.Println(jack)
	jackSargey := true
	fmt.Println(jackSargey)

	// another way of var declration

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


	var money int
	money = 5000
	fmt.Println(money)

	var dollar float32
	dollar = 44.34
	fmt.Println(dollar)


	//

	
	z = 1
	fmt.Println(z)
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(zx)


	// multiple variable

	var aa,bb,cc, dd int = 1, 2, 3, 4
	fmt.Println(aa,bb,cc,dd)
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)


	//

	var aaa, bbb = 4.4, "hello"
	ccc, ddd := 3.4, "world"

	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)
	fmt.Println(ddd)

// Variable Declaration in a Block
var (
	a1 int
	b1 float32 = 44.44
	c1 string = "Hello Go!"
)

fmt.Println(a1)
fmt.Println(b1)
fmt.Println(c1)


// Const
// constant variable rule same as var rule. only change here that const is not changeable of it's value;

const VALUE_OF_PI float32 = 3.1416
fmt.Println(VALUE_OF_PI)

const A = "Hello World!"
fmt.Println(A)


// Go Output Functions

/*

Print() , [do not add white space between arguments]
Println() , [add white space between arguments]
Printf() , it have to part: %v and %T
	%v is used to print the value of the arguments
	%T is used to print the type of the arguments


Print() if we se it, it will print normally. did not add white space between arguments.
if we print any integer then they will be stay separately but in string they will be one word.


*/






}

// this is single line comment system of Go
/* this is multi line comment system */

