#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

export GOPATH=${DIR}/go

# Build C Library
mkdir "${DIR}/py/cgo_lib" >/dev/null 2>&1
rm "${DIR}/py/cgo_lib/cgo_lib.a" "${DIR}/py/cgo_lib/cgo_lib.h" >/dev/null 2>&1
go build -o "${DIR}/py/cgo_lib/cgo_lib.a"   -buildmode=c-archive "${DIR}/go/src/com.stbaer/demo_cgo/"


# Build Python Extension
cd "${DIR}/py"
rm -rf *.so  build >/dev/null 2>&1
python setup.py build_ext --inplace