package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/heptio/tos3/s3"
)

const (
	regionFlag   = "region"
	bucketFlag   = "bucket"
	filepathFlag = "filepath"
	keyEnv       = "ACCESS_KEY_ID"
	secretEnv    = "SECRET_ACCESS_KEY"
)

var (
	bucket, filepath, region string
	key, secret              string
)

func init() {
	key = os.Getenv(keyEnv)
	secret = os.Getenv(secretEnv)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func requiredArg(arg, value string) {
	if value == "" {
		fmt.Printf(" * %v is required\n", arg)
		usage()
	}
}

func requiredEnv(env, value string) {
	if value == "" {
		fmt.Printf(" * %v is a required environment variable\n", env)
		usage()
	}
}

func main() {
	flag.StringVar(&region, regionFlag, "us-west-1", "the region the bucket lives in")
	flag.StringVar(&bucket, bucketFlag, "", "the bucket to upload the file to")
	flag.StringVar(&filepath, filepathFlag, "", "the file to upload")
	flag.Parse()
	requiredArg(bucketFlag, bucket)
	requiredArg(filepathFlag, filepath)
	requiredEnv(keyEnv, key)
	requiredEnv(secretEnv, secret)
	cfg := s3.Config(key, secret, region)
	s3.Upload(cfg, bucket, filepath)
}
