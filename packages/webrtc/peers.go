package webrtc

import (
	"Crunch-Garage/go-video-call/packages/chat"
	"sync"
)

type Room struct {
	Peers *Peers
	Hub   *chat .Hub
}

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP //learn what this does
}

func (p *Peers) DispatchKeyFrames() {

}
