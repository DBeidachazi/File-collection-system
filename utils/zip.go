package utils

import (
	"FengfengStudy/global"
	"context"
	"github.com/mholt/archiver/v4"
	"os"
)

// Zip zipDir recursively walks a directory and adds all files to a zip file
func Zip(dir, dst string) error {
	files, err := archiver.FilesFromDisk(nil, map[string]string{
		dir: "", // 相对路径
	})
	if err != nil {
		global.Log.Warnln("压缩文件失败", err)
		return err
	}

	// create output
	out, err := os.Create(dst)

	if err != nil {
		global.Log.Warnln("压缩文件失败", err)
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
		}
	}(out)

	format := archiver.CompressedArchive{
		Archival: archiver.Zip{},
	}

	err = format.Archive(context.Background(), out, files)
	if err != nil {
		global.Log.Warnln("压缩文件失败", err)
		return err
	}
	return nil
}
