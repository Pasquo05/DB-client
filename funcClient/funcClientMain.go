package funcclient

import (
	"fmt"
)

type PhoneNumber struct {
	Name        string
	Surname     string
	Phonenumber string
}

type ClientMatteo struct {
	c client
}

func NewClient(urlInput string) ClientMatteo {
	cl := ClientMatteo{
		c: New(urlInput),
	}
	return cl
}

func (c ClientMatteo) DeletePhoneNumber() {
	fmt.Println("ID ?")
	Id := ""
	fmt.Scanln(&Id)
	c.c.Delete("/phoneNumbers/delete/", Id)
}

func NewPhoneNumber() PhoneNumber {
	/*
		fmt.Println("Scrivere id , titolo , desc")
		idInput := ""
		titleInput := ""
		descInput := ""

		fmt.Scanln(&idInput)
		fmt.Scanln(&titleInput)
		fmt.Scanln(&idInput)
	*/

	value := PhoneNumber{
		Name:        "ciao",
		Surname:     "Bella",
		Phonenumber: "333333",
	}

	return value
}

func (a ClientMatteo) CreatePhoneNumber() {

	number := NewPhoneNumber()

	a.c.PostRequest("/phoneNumbers/create", number)

}

func (a ClientMatteo) GetPhoneNumber() {

	fmt.Println("ID ?")
	Id := ""
	fmt.Scanln(&Id)
	a.c.GetRequest(fmt.Sprintf("/phoneNumber/%s", Id))

}

func (a ClientMatteo) GetPhoneNumbers() {
	a.c.GetRequest("/phoneNumbers")
}
