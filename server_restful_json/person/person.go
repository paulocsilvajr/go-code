package person

import "time"

type Person struct {
	Name      string
	active    bool
	create_at time.Time
}

type Persons []Person
