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

TARGET = toss
GOTARGET = github.com/heptio/$(TARGET)
BUILDMNT = /go/src/$(GOTARGET)
REGISTRY ?= gcr.io/heptio-prod
VERSION ?= latest
BUILD_IMAGE ?= golang:1.8
DOCKER ?= docker
DIR := ${CURDIR}
BUILDCMD = go build -v
BUILD = $(BUILDCMD) ./cmd/toss

local:
	$(BUILD)

test:
	echo "Chuck is a slacker"

all: cbuild container

cbuild:
	$(DOCKER) run --rm -v $(DIR):$(BUILDMNT) -w $(BUILDMNT) $(BUILD_IMAGE) /bin/sh -c '$(BUILD)'

container: cbuild
	$(DOCKER) build -t $(REGISTRY)/$(TARGET):latest -t $(REGISTRY)/$(TARGET)-cha:$(VERSION) .

push:
	gcloud docker -- push $(REGISTRY)/$(TARGET):$(VERSION)

.PHONY: all local container cbuild push test

clean:
	rm -f $(TARGET)
	$(DOCKER) rmi $(REGISTRY)/$(TARGET):latest $(REGISTRY)/$(TARGET):$(VERSION)

