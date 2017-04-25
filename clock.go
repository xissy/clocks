package main

import (
	"time"

	"github.com/getwe/figlet4go"
)

const dateFormat = "Monday January 2 MST -07:00"
const timeFormat = "MST 15:04:05 PM"

// Clock widget structure
type Clock struct {
	Time       time.Time
	Location   *time.Location
	DateFormat string
	TimeFormat string
	TimeFiglet string
	TimeString string
	DateString string
}

// NewClock creates a new Clock widget
func NewClock(locationName string) (*Clock, error) {
	location, err := time.LoadLocation(locationName)
	if err != nil {
		return nil, err
	}

	clock := &Clock{
		Time:       time.Now().In(location),
		Location:   location,
		DateFormat: dateFormat,
		TimeFormat: timeFormat,
	}
	clock.Update()

	return clock, nil
}

// Update updates the clock to now
func (c *Clock) Update() error {
	c.Time = time.Now().In(c.Location)
	c.TimeString = c.Time.Format(c.TimeFormat)
	c.DateString = c.Time.Format(c.DateFormat)

	ascii := figlet4go.NewAsciiRender()
	timeFiglet, err := ascii.Render(c.TimeString)
	if err != nil {
		return err
	}

	c.TimeFiglet = timeFiglet

	return nil
}
