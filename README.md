vdir
====

[![API Documentation](http://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](http://godoc.org/github.com/xconstruct/vdir)
[![BSD License](http://img.shields.io/badge/license-BSD-blue.svg?style=flat-square)](http://opensource.org/licenses/BSD-2-Clause)

Go Package vdir implements RFC 2425 directory encoding with support for vCard
and iCalendar profiles.

The directory blocks can be mapped to an arbitrary Go Value which is
described in the Marshal and Unmarshal functions. In addition, the package
provides conversions to/from a generalized Object representation which
allows for detailed access and easier modification.

Install
-------
```
go get "github.com/xconstruct/vdir"
```

Example
-------
```go
vcf := []byte("BEGIN:VCARD ...")

var c vdir.Card
err := vdir.Unmarshal(vcf, &c)
fmt.Println(c.FormattedName)
fmt.Println(c.Addresses[0].Street)
```
