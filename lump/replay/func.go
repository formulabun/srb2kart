package replay

import (
	"fmt"
	"time"

	"go.formulabun.club/functional/strings"
)

func (r *HeaderPreFileEntries) GetTitle() string {
	return strings.SafeNullTerminated(r.Title[:])
}

func (r *HeaderPostFileEntries) GetTime() time.Duration {
	return time.Millisecond * time.Duration(1000*r.Time/35)
}

func (r *HeaderPostFileEntries) GetBestLapTime() time.Duration {
	return time.Millisecond * time.Duration(1000*r.Lap/35)
}

func (r *HeaderPreFileEntries) GetGuestFileName() string {
	return fmt.Sprintf("MAP%02d-guest.lmp", r.GameMap)
}
