package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Hello Golang")
	values()
	variables()
	constants()
	forLoops()
	ifElseLoop()
	switchLoop()
	arrays()
	slices()
}

func slices() {
	//An important data type in go, giving a more powerful interface sequences than arrays.

	// Unlike array's, Slices are types only by the elements they contain (not the number of elements)
	// An uninitialized slice equals to nil and has length 0

	fmt.Println("Slices")

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create an empty slice with non-zero length, use the builtin `make`. Here we make a slice of `string`
	// of length 3,
	//By default a new slices capacity is equal to it's length; if we know the slice is going to grow ahead of time
	// its possible to pass a capacity explicitly

	s = make([]string, 3)
	fmt.Println("emp :", s, "len:", len(s), "cap: ", cap(s))

	//We can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set :", s)
	fmt.Println("get :", s[2])

	//len returns the length of the slice as expected

	fmt.Println("len:", len(s))

	// In  addition, Slices support more which makes it richer than array
	// Builtin append, which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we can get new value

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd::", s)

	//slices can also be copy'd
	// We create an empty slice 'C' ot the same length as 's' and copy C into s

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy::", c)

	//Slice's support a "Slice" operator with the syntax
	// slice[low:high] ->
	//
	l := s[2:5]
	fmt.Println("sl1 :: ", l)

	//This slices up to (but excluding) `s[5]`
	l = s[:5]
	fmt.Println("sl2", l)

	//And this slices up from (and including) `s[2]`
	l = s[2:]
	fmt.Println("sl3", l)

	//We can declare and initialize a variable for slice
	//in a single line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//Slices can be composed into multi-dimensional  data structure
	//The length of the inner slices can vary, unlike with multi-dimensional arrays

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func arrays() {
	// In Go Array is a numbered sequence of elements of a specific length, In typical Go code, [Slices](slices) are
	//much more common; Arrays are useful in some special scenarios

	fmt.Println("an array 'a' that hold exactly 5 int's, And bdefalut an array is zero-valued.")
	var a [5]int
	fmt.Println("emp :", a)

	//We can set a value at an index using the array[index] = value, and get value
	fmt.Println("We can set a value at an index using the array[index] = value, and get value")
	a[4] = 100
	fmt.Println("Set ::", a)
	fmt.Println("Get ::", a[4])

	//Length
	fmt.Println("len ::", len(a))

	//The builtIn 'len' returns the length of an array
	fmt.Println("len :", len(a))

	//Use this syntax to declare and initialize an array in one line
	fmt.Println("Use this syntax to declare and initialize an array in one line")
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl :", b)

	fmt.Println("2d array")
	//array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2 d :", twoD)
}

func switchLoop() {
	// Switch statements express conditionals across many branches
	fmt.Println("Here is basic switch")
	i := 2
	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("Multiple expression in switch")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println("Switch without expression and case expressionas can be non-constants.")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("it's after noon")
	}

	//A type switch compares types instead of values. You can use this to discover the types of an interface value
	// In this ex, the variable `t` will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Println("Don't know type %T \n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func ifElseLoop() {
	//Branching with `if` and `else` in Go is straight-forward

	fmt.Println("Basic Example")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// A Statement can precede conditionals; any variables
	// declared in this statement are available in the current and all subsequent branches.

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// Note that you don't need parentheses around conditions
	// in Go, but that the braces are required.
}

func forLoops() {
	//most basic type with single condition
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	// A classic initial/condition/after `for` loop
	fmt.Println("A classic initial/condition/after `for` loop")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//'for' without a condition will loop repeatedly until you `break` out of loop or `return`
	// from enclosing function

	fmt.Println("'for' without a condition will loop repeatedly until you `break` out of loop or `return`")
	for {
		fmt.Println("loop")
		break
	}

	//you can also continue to the next iteration of the loop
	fmt.Println("you can also continue to the next iteration of the loop")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

const s string = "constant"

func constants() {
	// Go supports _constants_ of character, string, boolean,
	// and numeric values.
	fmt.Println(s)

	// A `const` statement can appear anywhere a `var`
	// statement can.
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	fmt.Println(math.Sin(n))
}

func variables() {
	// `var` declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	// This syntax is only available inside functions.
	f := "apple"
	fmt.Println(f)
}

func values() {
	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
