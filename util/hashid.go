package util

import (
    "github.com/speps/go-hashids"
)

const (
	minLength = 24
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-"
)

func HashIdDecode(salt string, i string) (int64, error) {
    r := int64(0)
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	hd.Alphabet = alphabet
	h, err := hashids.NewWithData(hd)
	if err != nil{
	    return r, err
    }

    o, err := h.DecodeInt64WithError(i)
	if err != nil{
		return r, err
	}else {
	    r = o[0]
	}

    return r, nil
}

func HashIdEncode(salt string, i uint64) (string, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	hd.Alphabet = alphabet

    h, err := hashids.NewWithData(hd)
	if err != nil{
		return "", err
	}

    v := make([]int64, 0)
    v = append(v, int64(i))

    return h.EncodeInt64(v)
}
