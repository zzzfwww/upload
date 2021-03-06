package upload

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
	"upload/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName + fmt.Sprintf("%v", time.Now().Second()))

	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return "storage/uploads"
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsNotExist(err)
}

func CheckContainExt(name string) bool {
	return true
}

func CheckMaxSize(f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	if size >= 100*1024*1024 {
		return true
	}

	return false
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
