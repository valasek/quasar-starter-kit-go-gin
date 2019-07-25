// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/valasek/quasar-starter-kit-go-gin/server/logger"
)

// API - structure to support DB persistence
type API struct {
	// entityName     *models.EntityManager
}

// NewAPI -
func NewAPI() *API {

	return &API{
		// entity:  		entitymgr,
	}
}

// Download -
func (api *API) Download(c *gin.Context) {
	fileName := "./data/README.md"

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		c.String(http.StatusNotFound, "file not found")
		return
	}

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	file.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := file.Stat()  //Get info from file
	FileSize := FileStat.Size() //Get file size

	//Send the headers
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="README.md"`,
	}

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	file.Seek(0, 0)
	c.DataFromReader(http.StatusOK, FileSize, FileContentType, file, extraHeaders)
	return
}

// Upload -
func (api *API) Upload(c *gin.Context) {

	// parse and validate file and post parameters
	form, err := c.MultipartForm()
	if err != nil {
		logger.Log.Error("unable to upload file, INVALID_FILE: ", err)
		c.String(http.StatusBadRequest, "INVALID_FILE: "+err.Error())
		return
	}
	uploadFileName := ""
	for k := range form.File {
		uploadFileName = k
		break
	}
	if uploadFileName == "" {
		logger.Log.Error("unable to upload file, INVALID_FILE: empty filename")
		c.String(http.StatusBadRequest, "INVALID_FILE: empty filename")
		return
	}
	ffile, err := c.FormFile(uploadFileName)
	if err != nil {
		logger.Log.Error("unable to upload file, INVALID_FILE: ", err)
		c.String(http.StatusBadRequest, "INVALID_FILE: "+err.Error())
		return
	}
	file, err := ffile.Open()
	if err != nil {
		logger.Log.Error("unable open file, INVALID_FILE: ", err)
		c.String(http.StatusBadRequest, "INVALID_FILE: "+err.Error())
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Log.Error("unable to upload file: INVALID_FILE: ", err)
		c.String(http.StatusBadRequest, "INVALID_FILE: "+err.Error())
		return
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	if (filetype != "application/zip") && (filetype != "application/x-zip-compressed") {
		logger.Log.Error("unable to upload file, INVALID_FILE_TYPE (supported: application/zip, application/x-zip-compressed): ", filetype)
		c.String(http.StatusBadRequest, "INVALID_FILE_TYPE: "+filetype)
		return
	}
	fileName := randToken(12) + ".zip"
	uploadPath := viper.GetString("uploadFolder")
	newPath := filepath.Join(uploadPath, fileName)
	// FIXME check file type
	// fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		logger.Log.Error("unable to upload file, CANT_WRITE_FILE: ", err)
		c.String(http.StatusInternalServerError, "CANT_WRITE_FILE: "+err.Error())
		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		logger.Log.Error("unable to upload file. CANT_WRITE_FILE: ", err)
		c.String(http.StatusInternalServerError, "CANT_WRITE_FILE: "+err.Error())
		return
	}
	// implement other functionality and handling of the file before it is deleted
	err = os.RemoveAll(viper.GetString("uploadFolder"))
	if err != nil {
		logger.Log.Error("unable to delete: ", err)
		c.String(http.StatusInternalServerError, "CANT_DELETE: "+err.Error())
		return
	}
	err = os.Mkdir(viper.GetString("uploadFolder"), os.ModeDir)
	if err != nil {
		logger.Log.Error("unable to delete: ", err)
		c.String(http.StatusInternalServerError, "CANT_DELETE: "+err.Error())
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", ffile.Filename))
}

// DownloadDocs -
func (api *API) DownloadDocs(c *gin.Context) {

	fileName := filepath.Join("documentation", "documentation.md")
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		c.String(http.StatusOK, fileName+" does not exist")
		return
	}
	fi, err := f.Stat()
	if err != nil {
		c.String(http.StatusOK, fileName+" cannot get file size")
		return
	}
	if fi.Size() == 0 {
		c.String(http.StatusOK, fileName+" does not exist")
		return
	}
	c.File(fileName)
}

// DownloadLogs -
func (api *API) DownloadLogs(c *gin.Context) {

	logLevel := c.Param("logLevel")
	if len(logLevel) < 1 {
		logger.Log.Error("unable to download log files, param 'logLevel' is missing")
		c.String(http.StatusInternalServerError, "unable to download log files, param 'logLevel' is missing")
		return
	}

	file := ""
	switch logLevel {
	case "0":
		file = "info.log"
	case "1":
		file = "error.log"
	default:
		logger.Log.Error("unable to download log files, unknown logLevel: ", logLevel)
	}
	fileName := filepath.Join(viper.GetString("logFolder"), file)
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		c.String(http.StatusOK, file+" contains no log entries")
		return
	}
	fi, err := f.Stat()
	if err != nil {
		c.String(http.StatusOK, file+" cannot get file size")
		return
	}
	if fi.Size() == 0 {
		c.String(http.StatusOK, file+" contains no log entries")
		return
	}
	c.File(fileName)
}

// // exports all data from DB into file timesheet-backup.zip
// func export() (fileName string, err error) {
// 	fileName = "timesheet-backup.zip"
// 	exportFolder := viper.GetString("export.location")

// 	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
// 	file, err := os.OpenFile("timesheet-backup.zip", flags, 0644)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	files, err := ioutil.ReadDir(exportFolder)
// 	if err != nil {
// 		return "", err
// 	}

// 	zipw := zip.NewWriter(file)
// 	defer zipw.Close()

// 	for _, file := range files {
// 		err := appendFiles(filepath.Join(exportFolder, file.Name()), zipw)
// 		if err != nil {
// 			return "", err
// 		}
// 	}

// 	return fileName, nil
// }
