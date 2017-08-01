#!/usr/bin/env bash

# These environment variables should be passed into this script
# ACCESS_KEY_ID
# SECRET_ACCESS_KEY
# RESULTS_DIR
# BUCKET
# REGION

# TODO pass results into this garbage
ACCESS_KEY_ID="${ACCESS_KEY_ID}" SECRET_ACCESS_KEY="${SECRET_ACCESS_KEY}" RESULTS_DIR="${RESULTS_DIR}" /toss -bucket "${BUCKET}" -region "${REGION}"
echo "ALL DONE!"