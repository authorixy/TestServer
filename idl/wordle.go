package idl

import "github.com/TestServer/db"

type UpdateScoreRequest struct {
	Name string
	Score db.Score
}