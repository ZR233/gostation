/*
@Time : 2019-07-12 16:24
@Author : zr
*/
package model

import (
	model2 "github.com/ZR233/session/model"
	"strconv"
	"time"
)

type Session struct {
	Token    string
	UserId   int
	Channel  string
	ExpireAt time.Time
}

func NewSession(s *model2.Session) *Session {
	userId, err := strconv.Atoi(s.UserId)
	if err != nil {
		panic(err)
	}
	s_ := &Session{
		s.Token,
		userId,
		s.Channel,
		s.ExpireAt,
	}
	return s_
}
