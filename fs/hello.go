package fs

import (
	"bazil.org/fuse"
	_ "bazil.org/fuse/fs/fstestutil"
	"os/user"
	"strconv"

	"golang.org/x/net/context"
)

// File implements both Node and Handle for the hello file.
type File struct{}

const greeting = "hello there, this is cgroupfs\n"

func (File) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = INODE_HELLO
	a.Mode = 0777
	user, err := user.Current()
        uid, err := strconv.ParseInt(user.Uid, 10, 32)
	if err != nil {
		panic(err)
	}
	gid, err := strconv.ParseInt(user.Gid, 10, 32)
	if err != nil {
		panic(err)
	}
	a.Uid = uint32(uid)
	a.Gid = uint32(gid)
	a.Size = uint64(len(greeting))
	return nil
}

func (File) ReadAll(ctx context.Context) ([]byte, error) {
	return []byte(greeting), nil
}
