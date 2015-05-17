// Copyright (C) 2015 Constantin Schomburg <me@cschomburg.com>

package vdir

// Calendar is a iCalendar as defined in RFC 5545.
type Calendar struct {
	Profile  string `vdir:"vcalendar,profile"`
	Version  string
	ProdId   string
	Events   []Event    `vdir:",object"`
	ToDos    []Todo     `vdir:",object"`
	Journals []Journal  `vdir:",object"`
	FreeBusy []FreeBusy `vdir:",object"`
}

type Event struct {
	Profile     string `vdir:"vevent,profile"`
	UID         string
	DTStamp     string
	Organizer   Person
	DTStart     string
	DTEnd       string
	Summary     string
	Categories  []string
	Description string
	Method      string
	Status      string
	Class       string
	Sequence    string
	Alarms      []Alarm `vdir:",object"`
}

type Person struct {
	CommonName string `vdir:"cn"`
	Url        string
}

type Alarm struct {
	Profile     string `vdir:"valarm,profile"`
	Trigger     string
	Action      string
	Description string
	Repeat      string
	Duration    string
}

type Todo struct {
	Profile  string `vdir:"vtodo,profile"`
	DTStamp  string
	Sequence string
	UID      string
	Due      string
	Status   string
	Summary  string
	Alarms   []Alarm `vdir:",object"`
}

type Journal struct {
	Profile     string `vdir:"vjournal,profile"`
	DTStamp     string
	UID         string
	Organizer   Person
	Status      string
	Class       string
	Categories  []string
	Description string
}

type FreeBusy struct {
	Profile   string `vdir:"vfreebusy,profile"`
	Organizer Person
	DTStart   string
	DTEnd     string
	FreeBusy  []string `vdir:",multiple"`
	Url       string
}
