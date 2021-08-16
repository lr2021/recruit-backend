package model

type Profile struct {
	Username  string `json:"username"`
	Name      string `json:"name"`
	StuNumber string `json:"stuNumber"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
	Password  string `json:"password"`
	Sex       int32  `json:"sex"`
	Age       int32  `json:"age"`
	Grade     int32  `json:"grade"`
	QQNumber  string `json:"qqNumber"`
	Img       string `json:"img"`
}

type Solves struct {
	SolvedNumber  int32 `json:"solvedNumber"`
	Rank          int32 `json:"rank"`
	WebSolved     int32 `json:"webSolved"`
	AndroidSolved int32 `json:"androidSolved"`
	MLSolved      int32 `json:"MLSolved"`
	UISolved      int32 `json:"UISolved"`
	SecSolved     int32 `json:"SecSolved"`
}

type SolvedRecords struct {
	Username      string  `json:"username"`
	ProblemID     string  `json:"problemID"`
	ProblemTitle  string  `json:"problemTitle"`
	ProblemField  string  `json:"problemField"`
	ProblemScore  int32   `json:"problemScore"`
	Magnification float32 `json:"magnification"`
	SolvedTime    string  `json:"solvedTime"`
}

type LoginRequest struct {
	Tel            string `json:"tel"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	ValidationCode string `json:"validationCode"`
	Token		   string `json:"token"`
}

type LoginResponse struct {
	Status   int32  `json:"status"`
	Username string `json:"uid"`
	Msg      string `json:"msg"`
	Token    string `json:"token"`
}

type RegisterRequest struct {
	Tel              string `json:"tel"`
	StuNumber        string `json:"stuNumber"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeatedPassword"`
	Grade            int    `json:"grade"`
	ValidationCode   string `json:"validationCode"`
	Token			 string `json:"token"`
}

type RegisterResponse struct {
	Status   int32  `json:"status"`
	Msg      string `json:"msg"`
	Username string `json:"username"`
}

type LogoutRequest struct {
	Username string `json:"username"`
}

type LogoutResponse struct {
	Status   int32  `json:"status"`
	Msg      string `json:"msg"`
	Username string `json:"username"`
}

type GetUserProfileRequest struct {
	Username string `json:"username"`
	Tel      string `json:"tel"`
	Type     int32  `json:"type"` // type=1: show profile; type=2: update profile
}

type GetUserProfileResponse struct {
	Status  int32   `json:"status"`
	Msg     string  `json:"msg"`
	Profile Profile `json:"profile"`
	Solves  Solves  `json:"solves"`
}

type GetUserSolvedRequest struct {
	Username string `json:"username"`
}

type GetUserSolvedResponse struct {
	Status int32           `json:"status"`
	Msg    string          `json:"msg"`
	Solved []SolvedRecords `json:"solvedRecords"`
}

type UpdateUserProfileRequest struct {
	Profile          Profile `json:"profile"`
	NewPassword      string  `json:"newPassword"`
	RepeatedPassword string  `json:"repeatedPassword"`
}

type UpdateUserProfileResponse struct {
	Status   int32  `json:"status"`
	Msg      string `json:"msg"`
	Username string `json:"username"`
}

type GetUserRankRequest struct {
	Username string `json:"status"`
	Field    string `json:"field"` // all, web, ui, ml, android, security
}

type GetUserRankResponse struct {
	Status   int32  `json:"status"`
	Msg      string `json:"msg"`
	Username string `json:"username"`
	Field    string `json:"field"`
	Rank     int32  `json:"rank"`
}