package testdata

import (
	"FengfengStudy/global"
	"context"
	"github.com/mholt/archiver/v4"
	"os"
	"testing"
)

func _Test(t *testing.T) {
	files, err := archiver.FilesFromDisk(nil, map[string]string{
		"../files/1234/26/": "", // 相对路径
	})
	if err != nil {
		global.Log.Warnln(err)
		return
	}

	// create output
	out, err := os.Create("done.zip")
	if err != nil {
		global.Log.Warnln(err)
		return
	}
	defer out.Close()

	format := archiver.CompressedArchive{
		Archival: archiver.Zip{},
	}

	err = format.Archive(context.Background(), out, files)
	if err != nil {
		global.Log.Warnln(err)
		return
	}
}
