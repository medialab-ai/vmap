package vmap

import (
	"github.com/haxqer/vast"
)

// VMAP is used to express the structure of the ad inventory as a set of timed ad breaks
// within a publishers video content
type VMAP struct {
	Version  string    `xml:"version,attr"`
	AdBreaks []AdBreak `xml:"AdBreak"`
}

// AdBreak is point in time where one or more ads are scheduled for delivery
type AdBreak struct {
	TimeOffset     *Offset        `xml:"timeOffset,attr"`
	BreakType      string         `xml:"breakType,attr"`
	BreakID        string         `xml:"breakId,attr"`
	RepeatAfter    string         `xml:"repeatAfter,attr"`
	AdSource       AdSource       `xml:"AdSource"`
	TrackingEvents TrackingEvents `xml:"TrackingEvents"`
	Extensions     Extensions     `xml:"Extensions"`
}

// AdSource is used to describe the location for VAST ads to be retried from
type AdSource struct {
	ID               string   `xml:"id,attr"`
	AllowMultipleAds bool     `xml:"allowMultipleAds,attr"`
	FollowRedirects  bool     `xml:"followRedirects,attr"`
	AdTagURI         AdTagURI `xml:"AdTagURI"`
}

// AdTagURI is for specifiying a URI that will return VAST
type AdTagURI struct {
	TemplateType string `xml:"templateType,attr"`
	Text         string `xml:",cdata"`
}

// TrackingEvents is a list of tracking events
type TrackingEvents struct {
	Tracking []Tracking `xml:"Tracking"`
}

// Tracking is single tracking event
type Tracking struct {
	Event    vast.Tracking `xml:"event,attr"`
	Offset   vast.Offset   `xml:"offset,attr"`
	Tracking string        `xml:",chardata"`
}

// Extensions is a list of extension objects
type Extensions struct {
	Extension []Extension `xml:"Extension"`
}

// Extension is used to describe custom functionality contain in the object
type Extension struct {
	Type string `xml:"type,attr"`
}
