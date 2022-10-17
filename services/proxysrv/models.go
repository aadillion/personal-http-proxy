package proxysrv

type ProcessedRequest struct {
	Id      string
	Status  int
	Headers map[string]string
	Length  int
}
