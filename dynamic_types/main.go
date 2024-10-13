package main

import "fmt"

func main()  {
	fmt.Println(add(5, 10))        
	fmt.Println(add(5.5, 4.5))     
	fmt.Println(add(5, "string"))
}

func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}
	return nil
}