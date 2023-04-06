package wad

type Wad struct {
	Header
	Directory []*FileLump
}

type Header struct {
	Identification  [4]byte
	NumWads         int32
	InfoTableOffset int32
}

type FileLump struct {
	FilePos int32
	Size    int32
	Name    [8]byte
}
