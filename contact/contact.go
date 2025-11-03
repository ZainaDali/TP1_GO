package contact

type Contact struct {
	ID    uint
	Name  string
	Email string
}

var (
	contacts      = make(map[uint]Contact)
	nextID   uint = 1
)

func AddContact(name, email string) {
	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
	}

	contacts[nextID] = newContact

	nextID++
}
