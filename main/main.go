package main

func main()  {
	//// declaring string variable
	//var hello = "Hello world"
	//fmt.Println(hello)
	//// declaring int variable
	//var number int = 20
	//fmt.Println(number)
	////declaring float variable
	//var float float64 = 26.9
	//fmt.Println(float)
	//// declaring constant variable
	//const gender string = "Male"
	//fmt.Println("Gender = "+gender)
	//
	////multiple declaration
	//var (
	//	a = 5
	//	b = 10
	//	c = 15
	//)
	//fmt.Println("printing the multiple declared variables ",a,b,c)

	/*
	Chapter 4
	Questions on variables
	1 What are two ways to create a new variable? explicitly inferring the type,
	and when the type is not declared.
	2. What is the value of x after running:
	x := 5; x += 1? 6
	3.What is scope and how do you determine the scope of a variable in Go?
	A scope determines how and when a variable can be used
	Scope of a variable in go is determined by the function it's being declared
	4. What is the difference between var and const?
	variables declared with var can be reassigned while const cannot be reassigned
	5. Using the example program as a starting point,
	write a program that converts from Fahrenheit
	into Celsius. (C = (F - 32) * 5/9)
	6. Write another program that converts from feet
	into meters. (1 ft = 0.3048 m)
	*/

	//	Solution to question 5
	//var celsius float32
	//var fahrenheit float32 = 2.0
	//celsius = (fahrenheit - 32) * 5/9
	//fmt.Println("Celsius = ",celsius)

	// Solution to question 6
	//fmt.Println("10ft = ",0.3048 *10)


	/*control statement //switch case
	i := 0
	for i <= 10 {
		//switch case
		switch i {
		case 0:fmt.Println("Zero")
		case 1:fmt.Println("One")
		case 2:fmt.Println("Two")
		case 3:fmt.Println("Three")
		case 4:fmt.Println("Four")
		case 5:fmt.Println("Five")
		case 6:fmt.Println("Six")
		case 7:fmt.Println("Seven")
		case 8:fmt.Println("Eight")
		case 9:fmt.Println("Nine")
		case 10:fmt.Println("Ten")
		default: fmt.Println("Unknown number")
		}
		i +=1
	}
	// for loop and if statement
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i,"even")

		} else {
			fmt.Println(i,"odd")
		}
	}


	Problems
	1.What does the following program print:
		i := 10
		if i > 10 {
		fmt.Println("Big")
		} else {
		fmt.Println("Small")
		}
	Small
	2.Write a program that prints out all the numbers
	evenly divisible by 3 between 1 and 100. (3, 6, 9,
	etc.)
	3.Write a program that prints the numbers from 1
	to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five
	print "Buzz". For numbers which are multiples
	of both three and five print "FizzBuzz"


	Solution question 2
	for i := 1; i <= 100; i++ {
		if i  % 3 == 0{
			fmt.Println(i)
		}

	}
	*/

	// Solution question 3
	//for i := 1; i <= 100; i++ {
	//	if i  % 3 == 0{
	//		fmt.Println("Fizz")
	//	} else if i % 3 != 0 {
	//		fmt.Println("Buzz")
	//	}else {fmt.Println("FizzBuzz")}
	//
	//}
	//declaring a slice
	//var array[] int
	// appending to a slice
	//array = append(array, 34,78)
	//fmt.Println("Array ",append(array, 6,7,10,78,67))
	// declaring and assigning a slice
	//var slice = []int{1,2,3,4,5,6}
	//fmt.Println("Slice ",slice)
	//declaring a map
	//var _map = make(map[string]string)
	// assigning values into a map
	//_map["name"]= "titobi"
	//_map["age"]= "12"
	//_map["height"]= "30m"
	//fmt.Println("Map ",_map)

	//problems
	//1. How do you access the 4th element of an array or
	//slice? array[4]

	//2. What is the length of a slice created using:
	//make([]int, 3, 9)?
	//var na  = make([]int,3,9)
	//fmt.Println(len(na))

	//3. Given the array:
	//x := [6]string{"a","b","c","d","e","f"} what would x[2:5] give you? [c d e]
	//x := [6]string{"a","b","c","d","e","f"}
	//fmt.Println(x[2:5])

	//4.Write a program that finds the smallest number
	//in this list:
	//x := []int{
	//	48,96,86,68,
	//	57,82,63,70,
	//	37,34,83,27,
	//	19,97, 9,17,
	//}
	//smallestNumber:= 100
	//largestNumber:= 0
	////for i := 0; i < len(x); i++ {
	//for i := range x{
	//	if x[i]<smallestNumber {
	//		smallestNumber = x[i]
	//	}
	//	if largestNumber < x[i]{
	//		largestNumber = x[i]
	//	}
	//}
	//fmt.Println("Smallest number ",smallestNumber)
	//fmt.Println("Largest number ",largestNumber)

	//Functions
	// when defining a function in a function you do not add the name of the function,
	//and it must be assigned to a variable like the example below Note: it's called a local function
	//add:= func (x,y int) int{
	//	return x+y
	//}
	//fmt.Println("Addition of numbers =",add(4,5))

	//nextEven := makeEvenGenerator()
	//fmt.Println(nextEven()) // 0
	//fmt.Println(nextEven()) // 2
	//fmt.Println(nextEven()) // 4

	//Panic differ and recover
	//defer func() {
	//	str := recover()
	//	fmt.Println(str)
	//}()
	//panic("PANIC")

	//Problems
	//1.sum is a function which takes a slice of numbers
	//and adds them together. What would its function signature look like in Go?
		//func sum(numbers ...int)int{
		//
		//}
	//2.Write a function which takes an integer and
	//halves it and returns true if it was even or false
	//if it was odd. For example half(1) should return
	//(0, false) and half(2) should return (1,
	//true).ans =  fmt.Println(halves(2))
	//3. Write a function with one variadic parameter
	//that finds the greatest number in a list of numbers. ans = fmt.Println("Greatest number",variadic(1,5,3,4,2))
	//4. Using makeEvenGenerator as an example, write a
	//makeOddGenerator function that generates odd
	//numbers.
	//oddNumbers := makeOddGenerator()
	//fmt.Println("Odd",oddNumbers())
	//5. The Fibonacci sequence is defined as: fib(0) =
	//0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
	//Write a recursive function which can find
	//fib(n) ans = fmt.Println(fibonacci(4))
	//6. What are defer, panic and recover? How do you
	//recover from a run-time panic? panic throws an error, recover helps to handle the panic
	//and defer is used to schedule a function call

	//Pointers
	//1. How do you get the memory address of a variable? memory address of a variable is gotten with the & sign
    //2. How do you assign a value to a pointer? *and the name of the variable
	//3. How do you create a new pointer? a new pointer is created with he new key word and the type
	//4. What is the value of x after running this program:
		//func square(x *float64) {
		// *x = *x * *x
		//}
		//func main() {
		// x := 1.5
		// square(&x)
		//}
	// answer = 2.25
	//5. Write a program that can swap two integers
	//(x := 1; y := 2; swap(&x, &y) should give you
	//x=2 and y=1)
	//x := 1
	//y := 2
	//newX , newY := swap(&x,&y)
	//fmt.Println("x =",newX,"\nY =",newY)
	//human:= Human{fName: "Titobi", lName: "Ligali", age: 12, cohort: "seven"}
	//living:= LivingThing{"Titobi"}

	//fmt.Println(human.getName())
	//fmt.Println("Human talking ",human.livingThing.talk("Titobiloluwa"))
	//fmt.Println("Living thing talking itself ",living.talk("Titobiloluwa"))
	//fmt.Println(human.move("I "))


}

// LivingThingInterface Declaration of an interface
type LivingThingInterface interface {
	move(movement string)string
}

// LivingThing inheritance and classes
type LivingThing struct {
	name string
}

//Method
func (livingThing *LivingThing) talk(name string) string  {
	return "My name is "+name
}

type Human struct {
	fName string
	lName string
	age   int64
	cohort string
	livingThing LivingThing
}
//Method
func (human *Human) getName() string  {
	return human.fName +" "+ human.lName
}

func (human *Human) move(movement string)string{
	return movement+human.getName() + " is Moving"
}
// End of inheritance and classes

// Functions
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func halves(number int) (int,bool) {
	if number % 2 ==0 {
		return number / 2 ,true
	}else {return number/ 2, false}
}

func variadic(numbers ...int)int  {
	largestNumber :=0
	for i:=range numbers{
		if largestNumber < numbers[i] {
			largestNumber = numbers[i]
		}
	}
	return largestNumber
}

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 3
		return i
	}
}

func fibonacci(number int) int {
	if number == 0 || number == 1{
		return number
	}else {
		return fibonacci(number - 1) + fibonacci(number - 2)
	}
}

func swap(x *int, y *int) (int,int) {
	j := new(int)
	*j = *y
	*y = *x
	*x = *j
	return *x,*y
}

