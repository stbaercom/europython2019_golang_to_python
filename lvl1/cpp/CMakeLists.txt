cmake_minimum_required(VERSION 3.14)
project(go_python_demo0)

set(CMAKE_CXX_STANDARD 17)


INCLUDE_DIRECTORIES(/Users/imhiro/AllFiles/103_prog/cpp/go_python_demo0/cgo_lib)
LINK_DIRECTORIES(/Users/imhiro/AllFiles/103_prog/cpp/go_python_demo0/cgo_lib)

add_executable(go_python_demo0 main.cpp cgo_lib/main.h)

TARGET_LINK_LIBRARIES(go_python_demo0 ${CMAKE_SOURCE_DIR}/cgo_lib/main.a)
