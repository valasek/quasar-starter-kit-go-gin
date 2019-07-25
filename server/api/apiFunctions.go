// Copyright © 2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"archive/zip"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	// "io/ioutil"
	"os"
	"path/filepath"
	// "strconv"

	"github.com/valasek/quasar-starter-kit-go-gin/server/logger"

	"github.com/spf13/viper"
)

// FileList returs map tables and input file for initial seeding
func FileList() map[string]string {
	list := map[string]string{
		"rates":            filepath.Join(".", "data", viper.GetString("data.rates")),
		"consultants":      filepath.Join(".", "data", viper.GetString("data.consultants")),
		"projects":         filepath.Join(".", "data", viper.GetString("data.projects")),
		"reported_records": filepath.Join(".", "data", viper.GetString("data.reportedRecords")),
		"holidays":         filepath.Join(".", "data", viper.GetString("data.holidays")),
	}
	return list
}

// func uploadedFileList() (list map[string]string, err error) {
// 	list = make(map[string]string)
// 	files, err := ioutil.ReadDir(filepath.Clean(viper.GetString("uploadFolderTemp")))
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(files) != 5 {
// 		return nil, errors.New("archive should contain 5 files")
// 	}

// 	for _, file := range files {
// 		table := tableFromFilename(file.Name())
// 		if len(table) > 0 {
// 			if _, ok := list[table]; ok {
// 				return nil, errors.New("archive contains same data: " + table)
// 			}
// 			list[table] = filepath.Join(".", viper.GetString("uploadFolderTemp"), file.Name())
// 		} else {
// 			logger.Log.Warn(file.Name(), " - ignored")
// 		}
// 	}

// 	if len(list) != 5 {
// 		logger.Log.Error("expected 5 files, got: ", list)
// 		return nil, errors.New("expected 5 files, got: " + strconv.Itoa(len(list)))
// 	}

// 	fmt.Println(list)
// 	return list, nil
// }

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			logger.Log.Error(err)
			panic(err)
		}
	}()

	extractAndWrite := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				logger.Log.Error(err)
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			return errors.New("got folder, zip and upload only csv files")
		}
		os.MkdirAll(filepath.Dir(path), f.Mode())
		ff, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer func() {
			if err := ff.Close(); err != nil {
				logger.Log.Error(err)
				panic(err)
			}
		}()

		_, err = io.Copy(ff, rc)
		if err != nil {
			return err
		}

		return nil
	}

	for _, f := range r.File {
		err := extractAndWrite(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func appendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filepath.Base(filename))
	if err != nil {
		msg := "failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("failed to write %s to zip: %s", filename, err)
	}

	return nil
}