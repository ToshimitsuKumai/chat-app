package chatgpt

type Service interface {
	Ask(message string) (string, error)
}
