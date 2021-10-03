package clock

import "fmt"

type Minutes struct {
	value int
}

func (m *Minutes) Increment() {
	m.value = (m.value + 1) % 60
}

func (m *Minutes) Set(newValue int) {
	m.value = newValue % 60
}

// No Receiver Pointer as it doesn't need to affect the field in anyway. But, for Consistency Reasons we'll keep it as Pointer Receiver
func (m *Minutes) Display() {
	fmt.Println(m.value)
}