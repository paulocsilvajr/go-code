package person

import "time"

type Person struct {
	Name      string
	Active    bool
	Create_at time.Time
}

type Persons []Person
