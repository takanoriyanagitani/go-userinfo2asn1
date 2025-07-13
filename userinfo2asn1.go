package userinfo2asn1

import (
	"encoding/asn1"
)

type NamedUser struct {
	Name string `asn1:"tag:0,utf8"`
	ID   int    `asn1:"tag:1"`
}

type UserKind struct {
	UserNumber int       `asn1:"tag:0,optional"`
	UserName   string    `asn1:"tag:1,optional,utf8"`
	NamedUser  NamedUser `asn1:"tag:2,optional"`
}

type UserInfo struct {
	Kind UserKind `asn1:"tag:0"`
}

type NamedGroup struct {
	Name string `asn1:"tag:0,utf8"`
	ID   int    `asn1:"tag:1"`
}

type GroupKind struct {
	GroupNumber int        `asn1:"tag:0,optional"`
	GroupName   string     `asn1:"tag:1,optional,utf8"`
	NamedGroup  NamedGroup `asn1:"tag:2,optional"`
}

type GroupInfo struct {
	Kind GroupKind `asn1:"tag:0"`
}

type BasicUserInfo struct {
	UserID       UserKind    `asn1:"tag:0"`
	PrimaryGroup GroupKind   `asn1:"tag:1"`
	SubGroups    []GroupKind `asn1:"tag:2"`
}

type BasicUserList []BasicUserInfo

// ToDER converts the BasicUserInfo struct to DER encoded bytes.
func (b *BasicUserInfo) ToDER() ([]byte, error) {
	return asn1.Marshal(*b)
}
