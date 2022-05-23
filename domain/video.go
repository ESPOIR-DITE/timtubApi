package domain

import "time"

type Video struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	DateUploaded string    `json:"dateUploaded"`
	Description  string    `json:"description"`
	IsPrivate    bool      `json:"isPrivate"`
	Price        int8      `json:"price"`
}

type VideoCategory struct {
	Id         string `json:"id"`
	VideoId    string `json:"videoId"`
	CategoryId string `json:"categoryId"`
}

type VideoComment struct {
	Id         string `json:"id"`
	VideoId    string `json:"videoId"`
	CategoryId string `json:"categoryId"`
}
type Category struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
