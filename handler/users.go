package handler

type UsersReq struct {
	Firstname         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	FilesAttachements []Files `json:"file_attachements,omitempty"`
}

type Files struct {
	FileID    uint
	FileName  string
	FileType  string
	FielBytes []byte
}
