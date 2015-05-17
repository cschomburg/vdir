// Copyright (C) 2015 Constantin Schomburg <me@cschomburg.com>

package vdir

import "testing"

func TestMarshal(t *testing.T) {
	c := Card{
		Version:       "4.0",
		FormattedName: "Forrest Gump",
		Name: Name{
			FamilyName: []string{"Gump"},
			GivenName:  []string{"Forest"},
		},
		Org:   "Bubba Gump Shrimp Co.",
		Title: "Shrimp Man",
		Photo: Photo{
			Mediatype: "image/gif",
			Data:      "http://www.example.com/dir_photos/my_photo.gif",
		},
		Telephones: []TypedValue{
			{[]string{"work", "voice"}, "uri:tel:+1-111-555-1212"},
			{[]string{"home", "voice"}, "uri:tel:+1-111-404-1212"},
		},
		Addresses: []Address{
			{
				Type:        []string{"work"},
				Label:       "100 Waters Edge\nBaytown, LA 30314\nUnited States of America",
				Street:      "100 Waters Edge",
				Locality:    "Baytown",
				Region:      "LA",
				PostalCode:  "30314",
				CountryName: "United States of America",
			},
			{
				Type:        []string{"work"},
				Label:       "42 Plantation St.\nBaytown, LA 30314\nUnited States of America",
				Street:      "42 Plantation St.",
				Locality:    "Baytown",
				Region:      "LA",
				PostalCode:  "30314",
				CountryName: "United States of America",
			},
		},
		Emails: []TypedValue{
			{nil, "forrestgump@example.com"},
		},
	}
	b, err := Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n" + string(b))
}
