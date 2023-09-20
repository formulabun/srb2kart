package graphic

import (
	"encoding/binary"
	"errors"
	"io"
)

func readHeader(r io.Reader, head *header) (err error) {
	err = binary.Read(r, binary.LittleEndian, &head.width)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &head.height)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &head.leftOffset)
	if err != nil {
		return err
	}
	err = binary.Read(r, binary.LittleEndian, &head.topOffset)
	if err != nil {
		return err
	}

	return nil
}

// reader must be after the TopOffset value
func readImage(r io.Reader, head *header) (err error) {
	rs, ok := r.(io.ReadSeeker)
	if !ok {
		return errors.New("Can only decode doom image lumps from a ReadSeeker")
	}

	var pointers = make([]uint32, head.width)
	head.posts = make([]*post, head.width)
	err = binary.Read(r, binary.LittleEndian, &pointers)
	for i, ptr := range pointers {
		_, err = rs.Seek(int64(ptr), io.SeekStart)
		if err != nil {
			return err
		}

		head.posts[i] = &post{}
		err = binary.Read(rs, binary.LittleEndian, &head.posts[i].topDelta)
		if err != nil {
			return err
		}
		err = binary.Read(rs, binary.LittleEndian, &head.posts[i].length)
		if err != nil {
			return err
		}

		padding := make([]byte, 1)
		rs.Read(padding)

		data := make([]uint8, head.posts[i].length)
		err = binary.Read(rs, binary.LittleEndian, &data)
		if err != nil {
			return err
		}
		head.posts[i].data = data[:]

		rs.Read(padding)

		err = binary.Read(rs, binary.LittleEndian, &head.posts[i].end)
		if err != nil {
			return err
		}
	}

	return nil
}
