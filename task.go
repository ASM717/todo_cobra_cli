package main

// Task описывает нашу задачу
type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}
