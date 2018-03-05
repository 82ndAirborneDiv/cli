package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const awsProfileName = "default"
const awsRegionName = "ca-central-1"
const awsBucketName = "cli-update"
const awsBucketPrefix = "update/state"

var sourcePath = filepath.Join(getRootPath(), "public", "update")

var sess *session.Session

func main() {
	fmt.Printf("Uploading files from %s\n", sourcePath)

	// Enable loading shared config file
	os.Setenv("aws_SDK_LOAD_CONFIG", "1")
	// Specify profile to load for the session's config
	var err error
	sess, err = session.NewSessionWithOptions(session.Options{
		Profile: awsProfileName,
		Config:  aws.Config{Region: aws.String(awsRegionName)},
	})
	if err != nil {
		fmt.Println("failed to create session,", err)
		fmt.Println(err)
		os.Exit(1)
	}

	// Get list of files to upload
	fmt.Printf("Getting list of files\n")
	fileList := []string{}
	filepath.Walk(sourcePath, func(path string, f os.FileInfo, err error) error {
		if isDirectory(path) {
			// Do nothing
			return nil
		}
		fileList = append(fileList, path)
		return nil
	})

	// Upload the files
	fmt.Printf("Uploading %d files\n", len(fileList))
	for _, path := range fileList {
		fmt.Printf("Uploading %s\n", path)
		// An s3 service
		s3Svc := s3.New(sess)

		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Failed to open file", file, err)
			os.Exit(1)
		}
		fileInfo, _ := file.Stat()
		size := fileInfo.Size()
		buffer := make([]byte, size)
		file.Read(buffer)

		defer file.Close()
		var key string
		key = awsBucketPrefix + path
		key = strings.Replace(key, sourcePath, "", 1)
		fmt.Printf(" \\- Destination: %s\n", key)

		params := &s3.PutObjectInput{
			Bucket:             aws.String(awsBucketName),
			Key:                aws.String(key),
			Body:               bytes.NewReader(buffer),
			ContentLength:      aws.Int64(size),
			ContentType:        aws.String(http.DetectContentType(buffer)),
			ContentDisposition: aws.String("attachment"),
			ACL:                aws.String("public-read"),
		}
		_, err = s3Svc.PutObject(params)
		if err != nil {
			fmt.Printf("Failed to upload data to %s/%s, %s\n",
				awsBucketName, key, err.Error())
			return
		}
	}
}

func getRootPath() string {
	pathsep := string(os.PathSeparator)

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("Could not call Caller(0)")
	}

	abs := filepath.Dir(file)

	// When tests are ran with coverage the location of this file is changed to a temp file, and we have to
	// adjust accordingly
	if strings.HasSuffix(abs, "_obj_test") {
		abs = ""
	}

	var err error
	abs, err = filepath.Abs(filepath.Join(abs, "..", ".."))

	if err != nil {
		return ""
	}

	return abs + pathsep
}

func isDirectory(path string) bool {
	fd, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	switch mode := fd.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}
	return false
}
