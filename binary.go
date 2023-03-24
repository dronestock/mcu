package main

type binary struct {
	Go string `default:"${BINARY_GO=go}" json:"go,omitempty"`
	Yq string `default:"${BINARY_YQ=yq}" json:"yq,omitempty"`
}
