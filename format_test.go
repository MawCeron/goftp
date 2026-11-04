// Copyright 2015 Muir Manders.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package goftp

import (
	"testing"
	"time"
)

// A path is data, not a format string: whatever the caller put in it has to
// reach the server as it is.
func TestAPathWithPercentSignsReachesTheServerUnchanged(t *testing.T) {
	s := startPASVServer(t)
	c, err := DialConfig(Config{
		User:        "u",
		Password:    "p",
		DisableEPSV: true,
		Timeout:     3 * time.Second,
	}, s.ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	const dir = "/100% done %s %d"
	c.ReadDir(dir)

	for _, line := range s.sent() {
		if line == "LIST "+dir {
			return
		}
	}
	t.Errorf("the server never received %q; got %q", "LIST "+dir, s.sent())
}
