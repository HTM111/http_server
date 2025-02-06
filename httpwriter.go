package main

type httpWriter interface {
	Write([]byte)
	AddHeader(key string, value string)
	WriteHeader(i int)
}
