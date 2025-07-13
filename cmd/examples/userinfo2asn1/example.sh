#!/bin/sh

der2fq(){
	cat /dev/stdin |
		fq -d asn1_ber
}

der2jer(){
	cat /dev/stdin |
		xxd -ps |
		tr -d '\n' |
		python3 -m asn1tools \
			convert \
			-i der \
			-o jer \
			./basic-user-info.asn \
			BasicUserList \
			- |
		dasel \
			--read=json \
			--write=yaml |
		bat --language=yaml
}

./userinfo2asn1 |
	cut -d: -f2- |
	cut -b2- |
	xxd -r -ps |
	der2jer
