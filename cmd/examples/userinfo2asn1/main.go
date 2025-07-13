package main

import (
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"log"

	userinfo "github.com/takanoriyanagitani/go-userinfo2asn1"
)

var dummyInfo userinfo.BasicUserList = userinfo.BasicUserList{
	{
		UserID: userinfo.UserKind{
			UserName: "johndoe",
		},
		PrimaryGroup: userinfo.GroupKind{
			GroupName: "users",
		},
		SubGroups: []userinfo.GroupKind{
			{
				GroupName: "admin",
			},
			{
				GroupNumber: 1001,
			},
		},
	},
	{
		UserID: userinfo.UserKind{
			UserNumber: 1000,
		},
		PrimaryGroup: userinfo.GroupKind{
			GroupName: "wheel",
		},
		SubGroups: []userinfo.GroupKind{},
	},
	{
		UserID: userinfo.UserKind{
			NamedUser: userinfo.NamedUser{
				Name: "alice",
				ID:   1001,
			},
		},
		PrimaryGroup: userinfo.GroupKind{
			GroupNumber: 20,
		},
		SubGroups: []userinfo.GroupKind{
			{
				NamedGroup: userinfo.NamedGroup{
					Name: "devs",
					ID:   5000,
				},
			},
		},
	},
}

// der2stdout prints DER bytes to standard output in hex format.
func der2stdout(derBytes []byte) {
	fmt.Printf("DER Bytes (hex): %s\n", hex.EncodeToString(derBytes))
}

func main() {
	derBytes, err := asn1.Marshal(dummyInfo)
	if err != nil {
		log.Fatalf("Failed to marshal dummyInfo to DER: %v", err)
	}

	der2stdout(derBytes)
}
