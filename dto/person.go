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

// package sizeof function, it gives 24 as the output.
type myStruct struct {
	myBool  bool    // 1 byte
	myFloat float64 // 8 bytes
	myInt   int32   // 4 bytes
}

// To optimise the previous example, we have to rearrange the struct’s
// fields in order to minimise the amount of padding bytes that are required
// for allocating the struct. Since the myInt and myBool fields have a combined size less
// than a word size, they can be located right next to each other(order doesn’t matter),
// reducing the padding bytes required from 11 to 3; so it results in something like this:
// package sizeof function, it gives 16 as the output.
type Mystruct struct {
	myFloat float64 // 8 bytes
	myBool  bool    // 1 byte
	myInt   int32   // 4 bytes
}
