package files

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

// With this saveData function we are gonna pass our structs such as Task or Note and call their save function
func SaveData(data Saver) error {
	err := data.Save()

	if err != nil {
		return err
	}
	return nil

}

func OutputData(data Outputtable) error {
	data.Display()
	return data.Save()
}
