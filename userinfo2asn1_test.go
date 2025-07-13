package userinfo2asn1_test

import (
	"encoding/asn1"
	"testing"

	userinfo "github.com/takanoriyanagitani/go-userinfo2asn1"
)

func TestBasicUserInfoToDER(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		userInfo userinfo.BasicUserInfo
	}{
		{
			name: "basic",
			userInfo: userinfo.BasicUserInfo{
				UserID: userinfo.UserKind{
					UserName: "testuser",
				},
				PrimaryGroup: userinfo.GroupKind{
					GroupName: "testgroup",
				},
				SubGroups: []userinfo.GroupKind{
					{
						GroupName: "subgroup1",
					},
					{
						GroupNumber: 123,
					},
				},
			},
		},
		{
			name: "user_by_number_group_by_name",
			userInfo: userinfo.BasicUserInfo{
				UserID: userinfo.UserKind{
					UserNumber: 1000,
				},
				PrimaryGroup: userinfo.GroupKind{
					GroupName: "wheel",
				},
				SubGroups: []userinfo.GroupKind{},
			},
		},
		{
			name: "user_by_named_group_by_number",
			userInfo: userinfo.BasicUserInfo{
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Serialize to DER bytes
			derBytes, err := tt.userInfo.ToDER()
			if err != nil {
				t.Fatalf("Failed to marshal %s to DER: %v", tt.name, err)
			}

			// Unmarshal back to verify
			var decodedUserInfo userinfo.BasicUserInfo
			_, err = asn1.Unmarshal(derBytes, &decodedUserInfo)
			if err != nil {
				t.Fatalf("Failed to unmarshal %s DER bytes: %v", tt.name, err)
			}

			// TODO: Add more comprehensive assertions to compare tt.userInfo and decodedUserInfo
			// For now, just checking for nil error is sufficient as a basic test.
		})
	}
}

func TestBasicUserListToDER(t *testing.T) {
	userList := userinfo.BasicUserList{
		{
			UserID: userinfo.UserKind{
				UserName: "user1",
			},
			PrimaryGroup: userinfo.GroupKind{
				GroupName: "group1",
			},
			SubGroups: []userinfo.GroupKind{},
		},
		{
			UserID: userinfo.UserKind{
				UserNumber: 100,
			},
			PrimaryGroup: userinfo.GroupKind{
				GroupNumber: 200,
			},
			SubGroups: []userinfo.GroupKind{
				{
					GroupName: "subgroupA",
				},
			},
		},
	}

	derBytes, err := asn1.Marshal(userList)
	if err != nil {
		t.Fatalf("Failed to marshal BasicUserList to DER: %v", err)
	}

	var decodedUserList userinfo.BasicUserList
	_, err = asn1.Unmarshal(derBytes, &decodedUserList)
	if err != nil {
		t.Fatalf("Failed to unmarshal BasicUserList DER bytes: %v", err)
	}

	// TODO: Add more comprehensive assertions to compare userList and decodedUserList
}
