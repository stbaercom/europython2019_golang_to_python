#!/usr/bin/env bash


export GOPATH="/Users/imhiro/AllFiles/103_prog/europython_2019/lvl2/go"

cd /Users/imhiro/AllFiles/103_prog/europython_2019/lvl2/go/src/com.stbaer/demo_cgo
go build -o ../../../cgo_lib/cgo_lib.a -buildmode=c-archive


cp /Users/imhiro/AllFiles/103_prog/europython_2019/lvl2/go/cgo_lib/* /Users/imhiro/AllFiles/103_prog/europython_2019/lvl2/py/cgo_lib
cp /Users/imhiro/AllFiles/103_prog/europython_2019/lvl2/go/cgo_lib/* /Users/imhiro/AllFiles/103_prog/europython_2019/lvl2/cpp/cgo_lib