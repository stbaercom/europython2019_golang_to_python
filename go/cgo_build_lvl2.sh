#!/usr/bin/env bash

LVL=2
NAME="cgo_lib_lvl${LVL}"

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )"/.. >/dev/null 2>&1 && pwd )"

CGO="${DIR}/go/cgo_lib"

DSTPY="${DIR}/py/cgo_lib"

DSTCPP="${DIR}/cpp/cgo_lib"


#echo ${LVL} ${NAME} ${DIR} $'\n' ${SRCDIR} $'\n' ${DSTPY} $'\n' ${DSTCPP}

rm ${CGO}/*.{a,h} >/dev/null 2>&1

go build  -o "${CGO}/${NAME}.a" -buildmode=c-archive src/com.stbaer/demo_cgo/main.go

rm ${DSTPY}/*.{a,h} >/dev/null 2>&1
cp ${CGO}/*.{a,h} ${DSTPY}

rm ${DSTCPP}/*.{a,h} >/dev/null 2>&1
cp ${CGO}/*.{a,h} ${DSTCPP}

