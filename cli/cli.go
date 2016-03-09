package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "bazil.org/fuse/fs/fstestutil"

	"../../cgroupfs-master"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s MOUNTPOINT CGROUP_DIR [VETH_NAME]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if flag.NArg() != 2 && flag.NArg() != 3 {
		Usage()
		os.Exit(2)
	}
	vethName := ""
	if flag.NArg() == 3 {
		vethName = flag.Arg(2)
	}
	if err := cgroupfs.Serve(flag.Arg(0), flag.Arg(1), vethName); err != nil {
		log.Fatal(err)
	}
}
