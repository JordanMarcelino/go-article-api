package model

type WebResponse[T any] struct {
	Info map[string]any `json:"info"`
	Data T              `json:"data"`
}
