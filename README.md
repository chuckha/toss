# github.com/heptio/tos3

A simple script to upload a file to s3

# Goals

Be able to upload one (and maybe more if necessary) file to s3 with a configurable region, bucket, credentials and file.

# Why does this exist?

This will be run as a container as part of the sonobuoy master pod. It will take data written by the master and send it off to s3 for safe keeping.

# Build lifted entirely from https://github.com/thockin/go-build-template