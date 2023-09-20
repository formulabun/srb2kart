package graphic

type header struct {
	width      int16
	height     int16
	leftOffset int16
	topOffset  int16
	posts      []*post
}

type post struct {
	topDelta int8
	length   uint8
	_        byte
	data     []uint8
	_        byte
	end      uint8
}
