package replay

import (
	"fmt"
	"time"

	"go.formulabun.club/functional/strings"
)

func bytesToString(bs [16]byte) string {
	return strings.SafeNullTerminated(bs[:])
}

func (r *HeaderPreFileEntries) GetTitle() string {
	return strings.SafeNullTerminated(r.Title[:])
}

func (r *RecordAttackTimes) GetTime() time.Duration {
	return time.Millisecond * time.Duration(1000*r.Time/35)
}

func (r *RecordAttackTimes) GetBestLapTime() time.Duration {
	return time.Millisecond * time.Duration(1000*r.Lap/35)
}

func (r *HeaderPreFileEntries) GetGuestFileName() string {
	return fmt.Sprintf("MAP%02d-guest.lmp", r.GameMap)
}

func (r *HeaderPreFileEntries) IsRecordReplay() bool {
	return r.DemoFlags&0x2 != 0
}

func (p *PlayerEntryData) GetName() string {
	return bytesToString(p.Name)
}

func (p *PlayerEntryData) GetSkin() string {
	return bytesToString(p.Skin)
}

func (p *PlayerEntryData) GetColor() string {
	return bytesToString(p.Color)
}
