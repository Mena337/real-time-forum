package handler

type Reg struct {
	// Id         int    `json:"id"`
	Nickname   string `json:"nickname"`
	First_Name string `json:"firstname"`
	Last_Name  string `json:"lastname"`
	Gender     string `json:"gender"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Log struct {
	// Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
