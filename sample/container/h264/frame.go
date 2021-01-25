package h264

import (
	"github.com/kakami/pkg/bit"
)

// AVCPacket ...
type AVCPacket struct {
	PktType uint
	CTS     int64
}

// NewPacket ...
func NewPacket() *AVCPacket {
	return &AVCPacket{}
}

// ParseVideo ...
func (v *AVCPacket) ParseVideo(in []byte) (err error) {
	r := bit.NewReader(in)
	v.PktType = r.ReadUInt(8)
	v.CTS = int64(r.Read(24))

	return
}
