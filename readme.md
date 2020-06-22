# LiveWedge control library and samples

// Copyright 2015, Cerevo Inc. All rights reserved.  
// Use of this source code is governed by a BSD-style  
// license that can be found in the LICENSE file.  

This provides basic operations for the video switcher, "LiveWedge".

(See http://livewedge.cerevo.com/en/ about LiveWedge.)

This library is still alpha version. Compatiblity may break in future update.

## Supported operations
* Screen transfer: Cut, Mix, Dip, Wipe
* Sub screen control: PinP, Chroma-key
* Start and stop recording and broadcasting
* Upload a still picture and use it as ch.4 input source
* Find LiveWedge within the same network

## Not yet supported
Getting status from LiveWedge is still under construction. func (vsw Vsw) Request* are not yet fully documented.

## Contents
### libvsw
Common library for manupulating LiveWedge by network.

### autotrans
Sample program to make video transition automatically. See samples/autotrans/00Readme.txt

### samples/trans
Very simple program just to repeat cut and mix.

### samples/wipe
Very simple program just to repeat wipe all pattern. 

### samples/pinp
Very simple program of Picture in Picture. 

### samples/chromakey
Very simple program of chroma key. 

### samples/rec
Very simple program just to send recording start/stop command with web UI.

### samples/status
A sample program for getting status via UDP.

### samples/find
A sample program for finding a LiveWedge within the same network.

## How to build

0. Install go language.

Tested in linux/amd64,windows/amd64. Go version 1.14. I hope Mac works, too.

    go get github.com/cerevo/LiveWedge_API/v2

Or 

Write the code that loads the package and execute it.

    import (
        "github.com/cerevo/LiveWedge_API/v2/libvsw"
    )

    go run .
    go build .
    go test .


## API document

https://godoc.org/github.com/cerevo/LiveWedge_API/src/libvsw

## How to use
At first, create vsw instance.

	ip_address := "192.168.21.5" //ip address of LiveWedge
	vsw, err := libvsw.NewVsw(ip_address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open LiveWedge: %s\n", err)
		os.Exit(1)
	}
	defer vsw.Close()

Then, invoke method of vsw. For example, cut every 3 seconds.

	i := 0
	for {
		vsw.Cut(i + 1)  // src is {1..4}
		i = (i + 1) % 4
		time.Sleep(3 * time.Second)
	}

See samples/trans/sample_trans.go for complete example.
