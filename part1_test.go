//	Copyright (c) Miloš Rackov 2020
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	cardPub int
	doorPub int
	modulo  int
	subject int
}

type test struct {
	name              string
	args              args
	wantCardLoop      int
	wantDoorLoop      int
	wantEncryptionKey int
}

var tests = []test{
	{
		name: "example",
		args: args{
			cardPub: 5764801,
			doorPub: 17807724,
			modulo:  20201227,
			subject: 7,
		},
		wantCardLoop:      8,
		wantDoorLoop:      11,
		wantEncryptionKey: 14897079,
	},
	{
		name: "inverted example",
		args: args{
			cardPub: 17807724,
			doorPub: 5764801,
			modulo:  20201227,
			subject: 7,
		},
		wantCardLoop:      11,
		wantDoorLoop:      8,
		wantEncryptionKey: 14897079,
	},
	{
		name: "real task",
		args: args{
			cardPub: 15335876,
			doorPub: 15086442,
			modulo:  20201227,
			subject: 7,
		},
		wantCardLoop:      250288,
		wantDoorLoop:      14519824,
		wantEncryptionKey: 11707042,
	},
}

func Test_part1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCardLoop, gotDoorLoop, gotEncryptionKey := part1(tt.args.cardPub, tt.args.doorPub, tt.args.modulo, tt.args.subject)
			assert.Equal(t, tt.wantCardLoop, gotCardLoop)
			assert.Equal(t, tt.wantDoorLoop, gotDoorLoop)
			assert.Equal(t, tt.wantEncryptionKey, gotEncryptionKey)
		})
	}
}

func Benchmark_part1(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				part1(test.args.cardPub, test.args.doorPub, test.args.modulo, test.args.subject)
			}
		})
	}
}
