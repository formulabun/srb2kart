package conversion

import "time"

const Tic = time.Second / 35

func TimeToFrames(time time.Duration) uint {
	return uint(time.Abs() / Tic)
}

func FramesToTime(frames uint) time.Duration {
	return time.Duration(frames * uint(Tic))
}
