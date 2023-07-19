package main

import (
	"crypto/md5"
	"encoding/hex"
)

type Codec struct {
	hashmap map[string]string
}

func Constructor() Codec {
	return Codec{
		hashmap: map[string]string{},
	}
}

func (this *Codec) encode(longUrl string) string {
	hash := md5.Sum([]byte(longUrl))
	shortUrl := hex.EncodeToString(hash[:])
	this.hashmap[shortUrl] = longUrl
	return shortUrl
}

func (this *Codec) decode(shortUrl string) string {
	return this.hashmap[shortUrl]
}
