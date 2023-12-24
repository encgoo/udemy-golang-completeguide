package person

import (
	"testing"
)

func TestSetEmail(t *testing.T) {
	person := MakePerson("Alex", "Anderson", "aa@example.com", 91210)
	person.SetEmail("aa@hotmail.com")
	if person.Contact.Email != "aa@hotmail.com" {
		t.Errorf("Got wrong email %s", person.Contact.Email)
	}
}
