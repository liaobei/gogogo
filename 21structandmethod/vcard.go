package main

type Address struct {
	addressname string
}

type Vcard struct {
	name     string
	address  *Address
	img      string
	birthday string
}
