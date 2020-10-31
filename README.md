# aqua-go

A fantastic GO SDK for both Aqua's Enterprise and Wave Platforms<br/>
Status: Experimental

Documentation
----
https://godoc.org/github.com/jeffthorne/aqua-go


Installation
----
Install: go get -u github.com/jeffthorne/aqua-go


Usage
----
Default usage against a secure endpoint with InsecureSkipVerify set to true 

aqua, err := aqua.NewCSP("192.168.1.52", 443, "user id", "password")