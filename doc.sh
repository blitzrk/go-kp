#!/bin/bash

pkill godoc
godoc -html ./ga > ../../../../pkg/github.com/blitzrk/go-kp/ga/doc.html 2>/dev/null
godoc -html ./dp > ../../../../pkg/github.com/blitzrk/go-kp/dp/doc.html 2>/dev/null
godoc -http=:6060 -goroot="." > /dev/null 2>&1 &
xdg-open "http://localhost:6060/pkg/github.com/blitzrk/go-kp/" > /dev/null 2>&1

