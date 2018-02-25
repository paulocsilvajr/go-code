package person

import (
	"fmt"
	"time"
)

var currentId int

var ListPersons Persons

func init() {
	DaoCreatePerson(Person{Name: "Paulo C", Active: true, Create_at: time.Now()})
	DaoCreatePerson(Person{Name: "José da Silva"})
	DaoCreatePerson(Person{Name: "Paul"})
	DaoCreatePerson(Person{Name: "Joseph"})
}

func DaoFindPerson(id int) Person {
	for _, person := range ListPersons {
		if person.Id == id {
			return person
		}
	}

	return Person{}
}

func DaoCreatePerson(person Person) Person {
	currentId += 1
	person.Id = currentId
	ListPersons = append(ListPersons, person)

	return person
}

func DaoDestroyPerson(id int) error {
	for i, person := range ListPersons {
		if person.Id == id {
			ListPersons = append(ListPersons[:i], ListPersons[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Não foi encontrado uma pessoa com o id[%d] informado", id)
}
