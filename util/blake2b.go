// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util

import (
//	"golang.org/x/crypto/blake2b"
	"github.com/zeebo/blake3"
	"fmt"
)

// HashBlake2b calculates the hash blake2b(b).
func HashBlake2b(buf []byte) []byte {
//	hashedBuf := blake2b.Sum256(buf)
	hasher := blake3.New()
	hasher.Write(buf)
	hashedBuf := hasher.Sum(nil)
	fmt.Println("=hashedBuf={}",hashedBuf)
	return hashedBuf[:]
}
