package service

import (
	"errors"
	"mime/multipart"
	"os"
	"upload/pkg/upload"
)

type FileInfo struct {
	Name string
}

func (svc *Service) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckMaxSize(file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	return &FileInfo{Name: fileName}, nil
}
