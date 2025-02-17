package utils

import (
	xerror "device-manager/utils/error"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const CHAR_SET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const (
	ErrorCodeDoNotPermissionCreateUploadFolder = "do_not_permission_create_upload_folder"
	ErrorCodeDoNotPermissionCreateNewFile      = "do_not_permission_create_new_file"
	ErrorCanNotReadFileContent                 = "can_not_read_file_content"
	ErrorDoNotPermissionToUploadFile           = "do_not_permission_to_upload_file"
)

func GenerateFilename(prefix string, ext string, length int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = CHAR_SET[seededRand.Intn(len(CHAR_SET))]
	}
	return prefix + "_" + string(b) + "." + ext
}

func UploadAndReturnFilePath(rootPath string, prefixDir string, storageDir string, prefixPath string, userID int, file multipart.File, headers *multipart.FileHeader) string {
	fileNameArr := strings.Split(headers.Filename, ".")
	if len(fileNameArr) < 2 {
		return "fail"
	}
	extension := fileNameArr[len(fileNameArr)-1]

	prefixPathList := strings.Split(prefixPath, ".")
	fmt.Println(prefixPathList)
	if len(prefixPathList) > 1 {
		prefixPath = prefixPathList[0]
	}

	rootAndStoragePath := rootPath + "/" + storageDir
	if _, err := os.Stat(rootAndStoragePath + "/" + prefixDir); os.IsNotExist(err) {
		if err := os.Mkdir(rootAndStoragePath+"/"+prefixDir, 0755); err != nil {
			return "fail"
		}
	}

	fileDir := rootAndStoragePath + "/" + prefixDir + "/" + strconv.Itoa(userID)
	fileName, err := UploadFile(prefixPath, fileDir, file, extension)
	if err != nil {
		return "fail"
	}
	filePath := storageDir + "/" + prefixDir + "/" + strconv.Itoa(userID) + "/" + fileName
	return filePath
}

func UploadFile(prefix string, dirPath string, file multipart.File, extension string) (string, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.Mkdir(dirPath, 0755); err != nil {
			return "", xerror.NewError(http.StatusForbidden, ErrorCodeDoNotPermissionCreateUploadFolder, "Do not permission create upload folder")
		}
	}
	fileName := GenerateFilename(prefix, extension, 16)
	filePath := dirPath + "/" + fileName
	destFile, err := os.Create(filePath)
	if err != nil {
		return "", xerror.NewError(http.StatusForbidden, ErrorCodeDoNotPermissionCreateNewFile, "Do not permission to create new file")
	}
	defer destFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", xerror.NewError(http.StatusInternalServerError, ErrorCanNotReadFileContent, "Can't read file content")
	}

	if _, err := destFile.Write(fileBytes); err != nil {
		return "", xerror.NewError(http.StatusForbidden, ErrorDoNotPermissionToUploadFile, "Do not permission to  upload file")
	}

	return fileName, nil
}
func UploadAndReturnFilePathWithFileByte(rootPath string, prefixDir string, storageDir string, prefixPath string, userID int, fileBytes []byte, fileName string) string {
	fileNameArr := strings.Split(fileName, ".")
	if len(fileNameArr) < 2 {
		return "fail"
	}
	extension := fileNameArr[len(fileNameArr)-1]

	prefixPathList := strings.Split(prefixPath, ".")
	if len(prefixPathList) > 1 {
		prefixPath = prefixPathList[0]
	}

	rootAndStoragePath := rootPath + "/" + storageDir
	if _, err := os.Stat(rootAndStoragePath + "/" + prefixDir); os.IsNotExist(err) {
		if err := os.Mkdir(rootAndStoragePath+"/"+prefixDir, 0755); err != nil {
			return "fail"
		}
	}

	fileDir := rootAndStoragePath + "/" + prefixDir + "/" + strconv.Itoa(userID)
	fileName, err := UploadFileWithFileByte(prefixPath, fileDir, fileBytes, extension)
	if err != nil {
		return "fail"
	}
	filePath := storageDir + "/" + prefixDir + "/" + strconv.Itoa(userID) + "/" + fileName
	return filePath
}
func UploadFileWithFileByte(prefix string, dirPath string, fileBytes []byte, extension string) (string, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.Mkdir(dirPath, 0755); err != nil {
			return "", xerror.NewError(http.StatusForbidden, ErrorCodeDoNotPermissionCreateUploadFolder, "Do not permission create upload folder")
		}
	}
	fileName := GenerateFilename(prefix, extension, 16)
	filePath := dirPath + "/" + fileName
	destFile, err := os.Create(filePath)
	if err != nil {
		return "", xerror.NewError(http.StatusForbidden, ErrorCodeDoNotPermissionCreateNewFile, "Do not permission to create new file")
	}
	defer destFile.Close()

	if _, err := destFile.Write(fileBytes); err != nil {
		return "", xerror.NewError(http.StatusForbidden, ErrorDoNotPermissionToUploadFile, "Do not permission to  upload file")
	}

	return fileName, nil
}

func GetPermissionFile() os.FileMode {
	return 0755
}

func CreateFolder(folderPath string) error {
	//Create Folder if not exists
	if _, err := os.Stat(folderPath); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(folderPath, GetPermissionFile())
		} else {
			return err
		}
	}

	return nil
}

func UploadBase64Image(ext string, parent string, content string) (storedFile string, err error) {
	wd, _ := os.Getwd()
	parentDir := fmt.Sprintf("%s/storage/%v", wd, parent)
	if err := CreateFolder(parentDir); err != nil {
		return "", err
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	fileName := fmt.Sprintf("%v_%v.%v", time.Now().Unix(), r.Int63(), ext)
	fileDir := fmt.Sprintf("%v/%v", parentDir, fileName)

	decode, err := base64.StdEncoding.DecodeString(content)
	file, err := os.Create(fileDir)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err = file.Write(decode); err != nil {
		return "", err
	}
	return fmt.Sprintf("/storage/%v/%v", parent, fileName), nil
}

func CopyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dst, input, 0644); err != nil {
		return err
	}

	return nil
}
