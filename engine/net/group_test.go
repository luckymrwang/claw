package net

import (
	"testing"
)

type PeerTmp struct {
	uid string
}

func (u *PeerTmp) Send(msg []byte) {
}


var (
	g *Group
)

func init() {
	g = NewGroup()
}


func TestGroup(t *testing.T) {
	g.AddPeer("user1", &PeerTmp{"user1"})
	if g.AddPeer("user1", &PeerTmp{"user1"}) {
		t.Error("Repeated add peer should fail")
	}
	if g.AddPeer("user2", &PeerTmp{"user2"}) == false {
		t.Error("Add a new peer should success")
	}
}
