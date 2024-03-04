package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Address struct {
	City   string
	Street string
	No     int
}

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string
	Email     string `json:"-"`
	Age       int
	Address   Address
}

/*
json:
{
	"FirstName":"Mostafa",
	"LastName" :"Solati",
	"Email":"mostafa.solati@gmail.com",
	"Age":38,
	"Address": {
		"City":"Tehran",
		"Street":"Some Street name",
		"No":54
	}
}


xml:
<Customer>
	<FirstName>Mostafa</FirstName>
	<LastName>Solati</LastName>
	<Email>mostafa.solati@gmail.com</Email>
	<Age>38</Age>
	<Address>
		<City>Tehran</City>
		<Street>Some Street Name</Street>
		<No>54</No>
	</Address>
</Customer>

*/

func main() {
	user := Customer{
		FirstName: "Mostafa",
		LastName:  "Solati",
		Email:     "mostafa.solati@gmail.com",
		Age:       38,
		Address: Address{
			City:   "Tehran",
			Street: "Some Street Name",
			No:     54,
		},
	}
	buf := new(bytes.Buffer)
	file, _ := os.Create("customer.json")
	defer file.Close()
	err := json.NewEncoder(file).Encode(user)
	// fmt.Println(err)
	// fmt.Println(buf.String())
	bs, err := json.Marshal(user)
	fmt.Println(string(bs), err)

	user2 := new(Customer)
	fmt.Println(user2)
	json.NewDecoder(buf).Decode(user2)
	fmt.Println(user2)

	user3 := new(Customer)
	fmt.Println(user3)
	json.Unmarshal(bs, user3)
	fmt.Println(user3)

	xmlFile, _ := os.Create("customer.xml")
	xml.NewEncoder(xmlFile).Encode(user)

	gobFile, _ := os.Create("customer.gob")
	gob.NewEncoder(gobFile).Encode(user)
	//xml.Marshal()
	//xml.Unmarshal()

}
