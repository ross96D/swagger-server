package shared

type EmptyInput struct {
}

type ListOutput[T any] struct {
	List []T `json:"list"`
}

type DefResp[T any] struct {
	Body T
}
