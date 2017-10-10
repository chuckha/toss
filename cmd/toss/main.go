package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/heptio/toss/s3"
)

const (
	regionFlag = "region"
	bucketFlag = "bucket"
	keyEnv     = "ACCESS_KEY_ID"
	secretEnv  = "SECRET_ACCESS_KEY"
	resultsEnv = "RESULTS_DIR"
)

var (
	bucket, region          string
	key, secret, resultsDir string
)

func init() {
	key = os.Getenv(keyEnv)
	secret = os.Getenv(secretEnv)
	resultsDir = os.Getenv(resultsEnv)
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
	flag.Parse()
	requiredArg(bucketFlag, bucket)
	requiredEnv(keyEnv, key)
	requiredEnv(secretEnv, secret)
	requiredEnv(resultsEnv, resultsDir)

	doneFile := resultsDir + "/done"
	cfg := s3.Config(key, secret, region)
	contents := waitForFile(doneFile)

	for _, file := range bytes.Split(contents, []byte("\n")) {
		if len(file) == 0 {
			continue
		}
		err := s3.Upload(cfg, bucket, string(file))
		if err != nil {
			fmt.Printf("Error encountered in s3.Upload: %v", err)
		}
	}
}

func waitForFile(waitfile string) []byte {
	for {
		contents, err := ioutil.ReadFile(waitfile) // For read access.
		if err != nil {
			time.Sleep(1 * time.Second)
		} else {
			return contents
		}
	}
}
