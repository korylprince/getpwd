package getpwd

//Passwd contains user information in a format similar to pwd.h's struct passwd
type Passwd struct {
	Name    string
	Passwd  string
	UID     uint
	GID     uint
	GECOS   string
	HomeDir string
	Shell   string
}
