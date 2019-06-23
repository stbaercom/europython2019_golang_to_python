go build  -buildmode=c-archive src/com.stbaer/demo_cgo/main.go

cp -a main.* /Users/imhiro/AllFiles/103_prog/cpp/go_python_demo0/cgo_lib
cp -a main.* /Users/imhiro/AllFiles/103_prog/python/go_python_demo0/cgo_lib

rm main.{h,a}