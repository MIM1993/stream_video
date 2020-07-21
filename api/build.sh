#!/bin/bash

go install

rm -rf ../bin/api

mv $GOROOT/bin/api  ../bin/