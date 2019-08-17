package views

// basic response
type Response struct {
	Code    int
	Error   string
	Data    interface{}
}

// user profile
type Profile struct {
	Username         string
	Nickname         string
	Introduction     string
	TotalCollCount   uint16
	TotalFollowCount uint16
	CreateTime       int64
}
