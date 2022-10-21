package typ

type User struct {
	UserName     string `json:"user_name"`
	PassWord     string `json:"pass_word"`
	AuthVerified bool   `json:"auth_verified"`
}

type Auth interface {
	log()
	Verified(code any) bool
}

type AuthVerified struct {
}

func (a AuthVerified) log() {

}

func (a AuthVerified) Verified(code any) bool {
	switch code.(type) {
	case string:

	case bool:

	}
	return true
}
