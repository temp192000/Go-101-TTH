package clock

import (
	"fmt"
)

type Minutes int
// Disadvantage: Type Minute is still using an underlying Type int. So, it's possible to set an invalid value for Type Minute

func (m *Minutes) Increment(){
	*m = (*m + 1) % 60
}

func (m *Minutes) Display(){
	fmt.Println(*m)
}