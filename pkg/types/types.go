package types

// Tidbit represents the database interface
type Tidbit interface {
	Set()
	Get()
	Delete()

	Query()
}

// IntervalsValue represents the intervals the value is present in
type IntervalsValue struct {
	Intervals []Interval
	Value     interface{}
}

// Interval represents a time interval
type Interval struct {
	From int64
	To   int64
}
