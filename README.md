# aqua-go

A fantastic GO SDK for Aqua's Enterprise Platform
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

aqua, err := aqua.NewClient("192.168.1.52", 443, "user id", "password")