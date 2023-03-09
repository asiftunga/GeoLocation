package entities

type Auth struct {
	AuthId  int64  `json:"auth_id"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Email   string `json:"email,omitempty"`
}

type Login struct {
	Email    string `json:"user_email"`
	UserPass string `json:"user_pass"`
	//UserPass string `json:"user_pass" validate:"required,min=5,max=20"`
	//above code is simple server side validation.
	/*
		we can check validations like this :
			err := Validate.Struct(u)
			if err != nil {
				fmt.Println(err)
			}
	*/
}
