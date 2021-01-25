package container

// Packet ...
type Packet interface {
	ParseVideo(in []byte) (err error)
}
