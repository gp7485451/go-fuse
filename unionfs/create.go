package unionfs

import (
	"github.com/gp7485451/go-fuse/fuse"
	"os"
)

func NewUnionFsFromRoots(roots []string, opts *UnionFsOptions, roCaching bool) (*UnionFs, error) {
	fses := make([]fuse.FileSystem, 0)
	for i, r := range roots {
		var fs fuse.FileSystem
		fi, err := os.Stat(r)
		if err != nil {
			return nil, err
		}
		if fi.IsDir() {
			fs = fuse.NewLoopbackFileSystem(r)
		}
		if fs == nil {
			return nil, err

		}
		if i > 0 && roCaching {
			fs = NewCachingFileSystem(fs, 0)
		}

		fses = append(fses, fs)
	}

	return NewUnionFs(fses, *opts), nil
}
