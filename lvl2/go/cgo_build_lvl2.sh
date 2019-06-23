#!/usr/bin/env bash

LVL=2
NAME="cgo_lib/cgo_lib_lvl${LVL}"

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )"/.. >/dev/null 2>&1 && pwd )"

SRC="${DIR}/go/${NAME}.{a,h}"

DSTPY="${DIR}/py/cgo_lib"

DSTCPP="${DIR}/cpp/cgo_lib"


#echo ${LVL} ${NAME} ${DIR} $'\n' ${SRCDIR} $'\n' ${DSTPY} $'\n' ${DSTCPP}

rm "${DIR}/go/cgo_lib/*.{a,h}"

go build  -o ${NAME}.a -buildmode=c-archive src/com.stbaer/demo_cgo/main.go

rm ${DSTPY}/*.{a,h}
cp ${SRC} ${DSTPY}

rm ${DSTCPP}/*.{a,h}
cp ${SRC} ${DSTCPP}


rm ${SRC}