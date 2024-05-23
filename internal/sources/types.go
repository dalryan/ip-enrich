package sources

import "fmt"

type Result struct {
	Data interface{}
	Url  string
	Code int
	Done bool
	Name string
}

type APIQueryUnit struct {
	Name  string
	URL   string
	Model interface{}
}

type StatusMsg struct {
	URL  string
	KEY  string
	Code int
	DATA interface{}
}

type ErrMsg struct {
	URL string
	Err error
}

func (e ErrMsg) Error() string { return fmt.Sprintf("%s: %v", e.URL, e.Err) }
