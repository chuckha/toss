package s3

import (
	"fmt"
	"os"

	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Upload is a nice wrapper for uploading a file
func Upload(key, secret, bucket, filename string) error {
	// build credentials
	creds := credentials.NewStaticCredentials(key, secret, "")
	cfg := aws.NewConfig().WithRegion("us-west-2").WithCredentials(creds)

	// Startup a new session with our config
	svc := s3.New(session.New(), cfg)

	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	size := fileInfo.Size()
	buff := make([]byte, size)
	// Read in the file contents

	// get the filetype, we know it's tar.gz so replace this maybe
	fileType := http.DetectContentType(buff)
	path := "/sonobuoy/" + file.Name()

	params := &s3.PutObjectInput{
		Bucket:        &bucket,
		Key:           &path,
		Body:          file,
		ContentLength: &size,
		ContentType:   &fileType,
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		fmt.Printf("bad response: %s", err)
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))

	return nil
}