package domain

import (
	"errors"
	"net/http"
	"time"
)

type Video struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Date         time.Time `json:"date"`
	DateUploaded string    `json:"dateUploaded"`
	Description  string    `json:"description"`
	IsPrivate    bool      `json:"isPrivate"`
	Price        int8      `json:"price"`
	URL          string    `json:"url"`
}

func (v Video) Bind(r *http.Request) error {
	if v.Title == "" && v.Description == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type VideoCategory struct {
	Id         string `json:"id"`
	VideoId    string `json:"videoId"`
	CategoryId string `json:"categoryId"`
}

func (v VideoCategory) Bind(r *http.Request) error {
	if v.VideoId == "" && v.CategoryId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type VideoComment struct {
	Id        string `json:"id"`
	VideoId   string `json:"videoId"`
	CommentId string `json:"categoryId"`
}

func (v VideoComment) Bind(r *http.Request) error {
	if v.VideoId == "" && v.CommentId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type Category struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c Category) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type VideoRelated struct {
	Id             string `json:"id"`
	CurrentVideoId string `json:"currentVideo"`
	RelatedVideoId string `json:"relatedVideoId"`
}
