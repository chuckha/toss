#!/usr/bin/env bash
# Copyright 2017 Heptio Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# These environment variables should be passed into this script
# ACCESS_KEY_ID
# SECRET_ACCESS_KEY
# READ_RESULTS_DIR
# WRITE_RESULTS_DIR
# BUCKET
# REGION

ACCESS_KEY_ID="${ACCESS_KEY_ID}" SECRET_ACCESS_KEY="${SECRET_ACCESS_KEY}" RESULTS_DIR="${READ_RESULTS_DIR}" /toss -bucket "${BUCKET}" -region "${REGION}"

mkdir -p "${WRITE_RESULTS_DIR}"
cat "${READ_RESULTS_DIR}" > "${WRITE_RESULTS_DIR}"/done
