package dto

// type personFullName func(FirstName string, MiddleName string, LastName string) string
type Contact struct {
	Type   string `json:"type"`
	Detail string `json:"detail"`
}
type PAN struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}
type Person struct {
	ID         int
	FirstName  string
	MiddleName string
	LastName   string
	Salary     any
	IsMarried  bool
	Contacts   []Contact `json:"contacts"` // Array of contacts
	PAN        PAN
	Hobbies    []string `json:"hobbies"`

	// FullName   personFullName
}

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}
