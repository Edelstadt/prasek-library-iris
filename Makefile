#
#	You need to have an SSH key on the app
# 	TEST APP PORT 13644
# 	PRODUCTION APP PORT 13640
#

BASEDIR=$(shell pwd)
DESTINATIONMACHINE=alpha-node-6.rosti.cz
DESTINATIONPORT=13664
DESTINATIONDIR=/srv/app/
DESTINATIONUSER=app
NAME=knihovna
VERSION=0.1

.PHONY: all build upload clean

build:
	go build -i -o ${BASEDIR}/test ${BASEDIR}/main.go

upload: build
	rsync -avz --exclude=".*" -e "ssh -p ${DESTINATIONPORT}" ${BASEDIR}/ ${DESTINATIONUSER}@${DESTINATIONMACHINE}:${DESTINATIONDIR}

clean:
	rm -rf ${BASEDIR}/site

all: build upload
