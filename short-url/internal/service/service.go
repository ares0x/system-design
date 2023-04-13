package service

type Service interface {
	Encode(string, int) (string, error)
	Decode(string) (string, error)
}
