package document

// Type ...
const (
	CPF   = "CPF"
	CNPJ  = "CNPJ"
	EMPTY = "EMPTY"
)

// Email ...
type Email struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// IsValid ...
func (e *Email) IsValid() bool {
	return len(e.Address) > 0
}

// Document ...
type Document struct {
	ID     string  `json:"id"`
	Type   string  `json:"type"`
	Domain string  `json:"domain"`
	Owner  string  `json:"owner"`
	Emails []Email `json:"emails"`
}

// New ...
func New() *Document {
	return &Document{
		Emails: []Email{},
		Type:   EMPTY,
	}
}

// SetOwner ...
func (d *Document) SetOwner(owner string) {
	if len(owner) > 0 {
		d.Owner = owner
	}
}

// AddEmail ...
func (d *Document) AddEmail(email Email) {
	d.Emails = append(d.Emails, email)
}

// SetID ...
func (d *Document) SetID(handle string) {
	if len(handle) == 14 {
		d.ID = handle
		d.Type = CNPJ
	} else if len(handle) == 11 {
		d.ID = handle
		d.Type = CPF
	}
}
