package model

type User struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Name string `json:"name"`
	Password string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	Img      string `json:"img"`
	StuNumber string `json:"stuNumber"`
	ProblemSolvedNumber int `json:"problemSolvedNumber"`
	Grade               int    `json:"grade"`
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
	Tel      string `json:"tel"`
	Username string  `json:"username"`
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

type ServiceHealthCheckRequest struct {
}

type ServiceHealthCheckResponse struct {
	Health bool `json:"health"`
}