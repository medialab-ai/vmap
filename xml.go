package vmap

import (
	"encoding/xml"

	"github.com/haxqer/vast"
)

// VMAP is used to express the structure of the ad inventory as a set of timed ad breaks
// within a publishers video content
type VMAP struct {
	XMLName  xml.Name  `xml:"http://www.iab.net/videosuite/vmap vmap:VMAP"`
	Version  string    `xml:"version,attr"`
	AdBreaks []AdBreak `xml:"AdBreak"`
}

// AdBreak is point in time where one or more ads are scheduled for delivery
type AdBreak struct {
	XMLName        xml.Name       `xml:"vmap:AdBreak"`
	TimeOffset     *Offset        `xml:"timeOffset,attr"`
	BreakType      string         `xml:"breakType,attr"`
	BreakID        string         `xml:"breakId,attr"`
	RepeatAfter    string         `xml:"repeatAfter,attr"`
	AdSource       AdSource       `xml:"vmap:AdSource"`
	TrackingEvents TrackingEvents `xml:"vmap:TrackingEvents"`
	Extensions     Extensions     `xml:"vmap:Extensions"`
}

// AdSource is used to describe the location for VAST ads to be retried from
type AdSource struct {
	XMLName          xml.Name `xml:"vmap:AdSource"`
	ID               string   `xml:"id,attr"`
	AllowMultipleAds bool     `xml:"allowMultipleAds,attr"`
	FollowRedirects  bool     `xml:"followRedirects,attr"`
	AdTagURI         AdTagURI `xml:"vmap:AdTagURI"`
}

// AdTagURI is for specifiying a URI that will return VAST
type AdTagURI struct {
	XMLName      xml.Name `xml:"vmap:AdTagURI"`
	TemplateType string   `xml:"templateType,attr"`
	Text         string   `xml:",cdata"`
}

// TrackingEvents is a list of tracking events
type TrackingEvents struct {
	XMLName  xml.Name   `xml:"vmap:TrackingEvents"`
	Tracking []Tracking `xml:"vmap:Tracking"`
}

// Tracking is single tracking event
type Tracking struct {
	XMLName  xml.Name      `xml:"vmap:Tracking"`
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
	XMLName xml.Name `xml:"vmap:vmap:Extension"`
	Type    string   `xml:"type,attr"`
}
