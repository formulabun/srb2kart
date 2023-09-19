package replay

type ReplayRaw struct {
	HeaderPreFileEntries
	WadEntries
<<<<<<< Updated upstream
=======
	RecordAttackTimes
>>>>>>> Stashed changes
	HeaderPostFileEntries
	CVarEntries
	PlayerEntries
	PlayerListingEnd byte
}

type HeaderPreFileEntries struct {
	DemoHeader  [12]byte
	Version     uint8
	SubVersion  uint8
	DemoVersion uint16

	Title    [64]byte
	Checksum [16]byte

	Play    [4]byte
	GameMap uint16
	MapMd5  [16]byte

	DemoFlags uint8
	GameType  uint8

	FileCount byte
}

type WadEntries []WadEntry

type WadEntry struct {
	FileName string
	WadMd5   [16]byte
}

type HeaderPostFileEntries struct {
	Time uint32
	Lap  uint32

	Seed     uint32
	Reserved uint32

	CVarCount uint16
}

type CVarEntries []CVarEntry

type CVarEntry struct {
	CVarId uint16
	Value  string
	False  uint8
}

type PlayerEntries []PlayerEntry

type PlayerEntry struct {
	Spec uint8
	PlayerEntryData
}

type PlayerEntryData struct {
	Name   [16]byte
	Skin   [16]byte
	Color  [16]byte
	Score  uint32
	Speed  byte
	Weight byte
}
