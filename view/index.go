package view

import (
	"github.com/seventy-two/tweeter-go/model"
)

type Index struct {
	Page
	Tweets []model.Tweet
	IsEmpty bool
}
