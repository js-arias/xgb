// Copyright 2009 The XGB Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/js-arias/xgb"
)

func main() {
	c, err := xgb.Dial(os.Getenv("DISPLAY"))
	if err != nil {
		fmt.Printf("cannot connect: %v\n", err)
		os.Exit(1)
	}
	defer c.Close()

	fmt.Printf("vendor = %q\n", string(c.Setup.Vendor))

	win := c.NewId()
	gc := c.NewId()
	s := c.DefaultScreen()

	c.CreateWindow(0, win, s.Root, 150, 150, 200, 200, 0, 0, 0, 0, nil)
	c.ChangeWindowAttributes(win, xgb.CWBackPixel|xgb.CWEventMask,
		[]uint32{
			s.BlackPixel,
			xgb.EventMaskExposure | xgb.EventMaskKeyRelease,
		})
	c.CreateGC(gc, win, xgb.GCForeground, []uint32{s.WhitePixel})
	c.MapWindow(win)

	atom, _ := c.InternAtom(false, "HELLO")
	fmt.Printf("atom = %d\n", atom.Atom)

	// Set WmDelete event
	protName := "WM_PROTOCOLS"
	wmProt, _ := c.InternAtom(false, protName)
	wmProtocols := wmProt.Atom
	fmt.Printf("atom wmProtocols = %d\n", wmProtocols)
	atomName := "ATOM"
	atomTp, _ := c.InternAtom(false, atomName)
	atomType := atomTp.Atom
	fmt.Printf("atom atom = %d\n", atomType)
	wmDel := "WM_DELETE_WINDOW"
	atmDel, _ := c.InternAtom(false, wmDel)
	atomDel := atmDel.Atom
	fmt.Printf("atom wmDelete = %d\n", atomDel)
	wmDelete := make([]byte, 4)
	wmDelete[0] = byte(atomDel)
	wmDelete[1] = byte(atomDel >> 8)
	wmDelete[2] = byte(atomDel >> 16)
	wmDelete[3] = byte(atomDel >> 32)

	points := make([]xgb.Point, 2)
	points[0] = xgb.Point{5, 5}
	points[1] = xgb.Point{100, 120}

	c.ChangeProperty(xgb.PropModeReplace, win, xgb.AtomWmName, xgb.AtomString,
		8, []byte("demo window"))
	c.ChangeProperty(xgb.PropModeReplace, win, wmProtocols, atomType, 32, wmDelete)

	fontpaths, _ := c.GetFontPath()
	for _, fontpath := range fontpaths.Path {
		fmt.Printf("fontpath = %q\n", fontpath.Name)
	}

	hosts, _ := c.ListHosts()
	for _, host := range hosts.Hosts {
		fmt.Printf("hosts = %q\n", host.Address)
	}

	ecookie := c.ListExtensionsRequest()
	exts, _ := c.ListExtensionsReply(ecookie)
	for _, name := range exts.Names {
		fmt.Printf("exts = %q\n", name.Name)
	}

	for {
		reply, err := c.WaitForEvent()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("event %T\n", reply)
		switch event := reply.(type) {
		case xgb.ClientMessageEvent:
			if event.Type != wmProtocols {
				continue
			}
			if xgb.Id(event.Data.Data32[0]) != atomDel {
				continue
			}
			return
		case xgb.ExposeEvent:
			c.PolyLine(xgb.CoordModeOrigin, win, gc, points)
		case xgb.KeyReleaseEvent:
			fmt.Printf("key release!\n")
			points[0].X = event.EventX
			points[0].Y = event.EventY
			c.PolyLine(xgb.CoordModeOrigin, win, gc, points)
			c.Bell(75)
		}
	}
}
