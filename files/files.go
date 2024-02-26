package files

import "fmt"

// Our saver interface is gonna accept any structure wich contains a Save method attached  therefore we're gonna pass it task and note
type Saver interface {
	Save() error
}
type Displayer interface {
	Display()
}
type Outputtable interface {
	Saver
	Displayer
}

func Print(value interface{}) {
	vlAsInt, isInt := value.(int)
	vlAsStr, isStr := value.(string)

	if isStr {
		fmt.Print("Str: ", vlAsStr)
	}
	if isInt {
		fmt.Print("Int:", vlAsInt)
	}
	fmt.Print(value)
}
func OutputData(data Outputtable) error {
	data.Display()
	return data.Save()
}
