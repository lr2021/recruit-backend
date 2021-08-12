package model

type User struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	Sex      int32  `json:"sex"`
	Age      int32  `json:"age"`
	Address  string `json:"address"`
	ClassNum int32  `json:"classNum"`
	Img      string `json:"img"`
}

type AddUserRequest struct {
	User *User `json:"user"`
}

type AddUserResponse struct {
	Uid    string `json:"uid"`
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
}

type InspectUserRequest struct {
	Uid      string `json:"uid"`
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

type InspectUserResponse struct {
	User   *User  `json:"user"`
	Msg    string `json:"msg"`
	Status int32  `json:"status"`
}

type UpdateUserRequest struct {
	User *User `json:"user"`
}

type UpdateUserResponse struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
}

type DeleteUserRequest struct {
	Uid string `json:"uid"`
	Tel string `json:"tel"`
}

type DeleteUserResponse struct {
	Status int32  `json:"status"`
	Msg    string `json:"msg"`
}



