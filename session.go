// Copyright (c) 2018. Flying Gopher Authors
// license that can be found in the LICENSE file.

package sessions

import (
	"math/rand"
	"time"
)

type Session struct {
	ID                  uint64
	UserID              int
	CreatedAt, DeadLine time.Time
}

func RegisterSession(userid int) *Session {
	id := rand.Uint64()

	for id == 0 {
		id = rand.Uint64()
	}
	createdat := time.Now().UTC()
	deadline := time.Now().Add(7 * 24 * time.Hour).UTC()

	return &Session{id, userid, createdat, deadline}
}
