package entities

type User struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password []byte `json:"password,omitempty"` //for sha256
}

//type User struct {
//	Id       int64  `json:"id,omitempty"`
//	Name     string `json:"firstname,omitempty"`
//	LastName string `json:"lastname,omitempty"`
//}

//
//type Users struct {
//	Users []User `json:"Users,omitempty"`
//}
// bunu models kisminda deneyecegim
