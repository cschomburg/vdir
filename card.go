// Copyright (C) 2015 Constantin Schomburg <me@cschomburg.com>

package vdir

// Card is a vCard as defined in RFC 6350.
type Card struct {
	Profile       string `vdir:"vcard,profile"`
	Version       string
	FormattedName string `vdir:"fn"`
	Name          Name   `vdir:"n"`
	NickName      []string
	Birthday      string       `vdir:"bday"`
	Addresses     []Address    `vdir:"adr"`
	Telephones    []TypedValue `vdir:"tel"`
	Email         []TypedValue
	Url           []TypedValue
	Title         string
	Role          string
	Org           string
	Categories    []string
	Note          string
	URL           string
	Photo         Photo

	Rev    string
	ProdId string
	Uid    string

	XICQ   string `vdir:"x-icq"`
	XSkype string `vdir:"x-skype"`
	IMPP   []IMPP
}

type IMPP struct {
	Type         []string `vdir:",param"`
	XServiceType string   `vdir:"x-service-type,param"`
	Value        string
}

type Name struct {
	FamilyName        []string
	GivenName         []string
	AdditionalNames   []string
	HonorificNames    []string
	honorificSuffixes []string
}

type Photo struct {
	Encoding  string `vdir:",param"`
	MediaType string `vdir:",param"`
	Type      string `vdir:",param"`
	Value     string `vdir:",param"`
	Data      string
}

type Address struct {
	Type            []string `vdir:",param"`
	Label           string   `vdir:",param"`
	PostOfficeBox   string
	ExtendedAddress string
	Street          string
	Locality        string
	Region          string
	PostalCode      string
	CountryName     string
}

type TypedValue struct {
	Type  []string `vdir:",param"`
	Value string
}
