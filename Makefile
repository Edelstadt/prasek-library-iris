#
#	You need to have an SSH key on the app. Or password
# 	TEST APP PORT 13644
# 	PRODUCTION APP PORT 13664
#

BASEDIR=$(shell pwd)
DESTINATIONMACHINE=node-13.rosti.cz
DESTINATIONPORT=13664
DESTINATIONDIR=/srv/app/
DESTINATIONUSER=app
NAME=knihovna
VERSION=0.2

.PHONY: all build upload clean

build:
	stylus ${BASEDIR}/templates/style.styl -o ${BASEDIR}/templates/style.css
	python ${BASEDIR}/minify.py ${BASEDIR}/templates/style.css
	go build -ldflags="-s -w" -i -o ${BASEDIR}/test ${BASEDIR}/main.go
	#upx --brute ${BASEDIR}/test

upload: build

	rsync -avz --exclude=".*" -e "ssh -p ${DESTINATIONPORT}" ${BASEDIR}/ ${DESTINATIONUSER}@${DESTINATIONMACHINE}:${DESTINATIONDIR}


clean:
	rm -rf ${BASEDIR}/site

all: build upload
