package fixedlocationtime

import (
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Location interface {
	LocationUTC | LocationJST
	GetLocation() *time.Location
}

type Time[L Location] struct {
	time time.Time
}

func New[L Location]() Time[L] {
	return Time[L]{}
}

func FromTime[L Location](t time.Time) Time[L] {
	return Time[L]{
		time: t.In(L{}.GetLocation()),
	}
}

func FromTimestamppb[L Location](t *timestamppb.Timestamp) (Time[L], error) {
	err := t.CheckValid()
	if err != nil {
		return Time[L]{}, err
	}
	return Time[L]{time: t.AsTime().In(L{}.GetLocation())}, nil
}

func FromSpannerNullString[L Location](t spanner.NullTime) Time[L] {
	return Time[L]{time: t.Time.In(L{}.GetLocation())}
}

func (t *Time[L]) AsTime() time.Time {
	return t.time
}

func (t *Time[L]) AsTimestamppb() *timestamppb.Timestamp {
	return timestamppb.New(t.time)
}

func (t *Time[L]) AsSpannerNullString() spanner.NullTime {
	return spanner.NullTime{
		Time:  t.time,
		Valid: !t.time.IsZero(),
	}
}

func (t Time[L]) ZoneBounds() (start, end time.Time) { return t.time.ZoneBounds() }
func (t Time[L]) GobEncode() ([]byte, error)         { return t.time.GobEncode() }
func (t Time[L]) Hour() int                          { return t.time.Hour() }
func (t Time[L]) ISOWeek() (year, week int)          { return t.time.ISOWeek() }
func (t Time[L]) IsDST() bool                        { return t.time.IsDST() }
func (t Time[L]) IsZero() bool                       { return t.time.IsZero() }
func (t Time[L]) Location() *time.Location           { return t.time.Location() }
func (t Time[L]) MarshalBinary() ([]byte, error)     { return t.time.MarshalBinary() }
func (t Time[L]) MarshalJSON() ([]byte, error)       { return t.time.MarshalJSON() }
func (t Time[L]) MarshalText() ([]byte, error)       { return t.time.MarshalText() }
func (t Time[L]) Minute() int                        { return t.time.Minute() }
func (t Time[L]) Month() time.Month                  { return t.time.Month() }
func (t Time[L]) Nanosecond() int                    { return t.time.Nanosecond() }
func (t Time[L]) Second() int                        { return t.time.Second() }
func (t Time[L]) Unix() int64                        { return t.time.Unix() }
func (t Time[L]) UnixMicro() int64                   { return t.time.UnixMicro() }
func (t Time[L]) UnixMilli() int64                   { return t.time.UnixMilli() }
func (t Time[L]) UnixNano() int64                    { return t.time.UnixNano() }
func (t Time[L]) Weekday() time.Weekday              { return t.time.Weekday() }
func (t Time[L]) Year() int                          { return t.time.Year() }
func (t Time[L]) YearDay() int                       { return t.time.YearDay() }
func (t Time[L]) Zone() (name string, offset int)    { return t.time.Zone() }

type JST = Time[LocationJST]
type UTC = Time[LocationUTC]

type LocationUTC struct{}
type LocationJST struct{}

var (
	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
	utc = time.FixedZone("UTC", 0)
)

func (t LocationUTC) GetLocation() *time.Location { return utc }
func (t LocationJST) GetLocation() *time.Location { return jst }

// Compilation error
// jst := &Time[LocationJST]{}
// utc := &Time[LocationUTC]{}
// if jst == utc { }
