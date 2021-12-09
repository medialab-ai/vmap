package vmap

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type (
	// Duration is a time.Duration alias
	Duration time.Duration

	// Offset is a VAMP duration expressed as  hh:mm:ss
	Offset struct {
		TimeDur  Duration
		Position OffsetPosition
		Percent  int
	}

	// OffsetPosition is for either passing start/end as the position of the ad
	OffsetPosition string
)

const (
	// OffsetPositionStart "start"
	OffsetPositionStart OffsetPosition = "start"
	// OffsetPositionEnd "end"
	OffsetPositionEnd OffsetPosition = "end"
)

// MarshalText implements the encoding.TextMarshaler interface.
func (off *Offset) MarshalText() ([]byte, error) {
	if off == nil {
		return nil, nil
	}

	if off.Position == OffsetPositionStart || off.Position == OffsetPositionEnd {
		return []byte(off.Position), nil
	}

	if off.Percent != 0 {
		return []byte(fmt.Sprintf("%d%%", off.Percent)), nil
	}

	dur := off.TimeDur
	h := dur / Duration(time.Hour)
	m := dur % Duration(time.Hour) / Duration(time.Minute)
	s := dur % Duration(time.Minute) / Duration(time.Second)
	ms := dur % Duration(time.Second) / Duration(time.Millisecond)
	if ms == 0 {
		return []byte(fmt.Sprintf("%02d:%02d:%02d", h, m, s)), nil
	}
	return []byte(fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, ms)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (off *Offset) UnmarshalText(data []byte) (err error) {
	s := string(data)
	switch OffsetPosition(s) {
	case OffsetPositionStart, OffsetPositionEnd:
		off.Position = OffsetPosition(s)
		return nil
	}

	if strings.Contains(s, "%") {
		s = strings.Replace(s, "%", "", -1)
		off.Percent, err = strconv.Atoi(s)
		return err
	}

	dur := &off.TimeDur
	s = strings.TrimSpace(s)
	if s == "" || strings.ToLower(s) == "undefined" {
		*dur = 0
		return nil
	}
	parts := strings.SplitN(s, ":", 3)
	if len(parts) != 3 {
		return fmt.Errorf("invalid duration: %s", data)
	}
	if i := strings.IndexByte(parts[2], '.'); i > 0 {
		ms, err := strconv.ParseInt(parts[2][i+1:], 10, 32)
		if err != nil || ms < 0 || ms > 999 {
			return fmt.Errorf("invalid duration: %s", data)
		}
		parts[2] = parts[2][:i]
		*dur += Duration(ms) * Duration(time.Millisecond)
	}
	f := Duration(time.Second)
	for i := 2; i >= 0; i-- {
		n, err := strconv.ParseInt(parts[i], 10, 32)
		if err != nil || n < 0 || n > 59 {
			return fmt.Errorf("invalid duration: %s", data)
		}
		*dur += Duration(n) * f
		f *= 60
	}
	return nil
}
