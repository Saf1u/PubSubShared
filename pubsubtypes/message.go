package pubsubtypes

type Message struct {
	Id      int
	Topic   string
	Type    string
	Content string
}
const (
	CLOSE_CONN    = "close"
	REGISTER_CONN = "register"
)