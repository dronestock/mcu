package main

type replace struct {
	// 从
	From dependency `json:"from"`
	// 到
	To dependency `json:"to"`
}
