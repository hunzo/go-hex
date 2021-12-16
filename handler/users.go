package handler

type UsersReq struct {
	Firstname string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UsersResponse struct {
	Firstname string `json:"first_name"`
	LastName  string `json:"last_name"`
}
