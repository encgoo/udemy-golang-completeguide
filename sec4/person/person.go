package person

import (
	// module name = udemy/sec4, subfolder person, package name contactinfo
	"udemy/sec4/person/contactinfo"
)

type Person struct {
	FirstName string
	LastName  string
	Contact   contactinfo.ContactInfo // composition, embedding struct
}

// factory design pattern

func MakePerson(fName string, lName string, email string, zip int) Person {
	person := Person{
		FirstName: fName,
		LastName:  lName,
		Contact: contactinfo.ContactInfo{
			Email:   email,
			ZipCode: zip,
		},
	}
	return person
}

// receiver functions
func (pP *Person) SetEmail(email string) {
	// pointer shortcut. No need to get the pointer of Contact. Use the value directly.
	(*pP).Contact.SetEmail(email)
}
