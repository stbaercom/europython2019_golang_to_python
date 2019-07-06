cmake_minimum_required(VERSION 3.14)


SET(LVL "2")


project(cgo_demo_${LVL})

set(CMAKE_CXX_STANDARD 17)


INCLUDE_DIRECTORIES(/Users/imhiro/AllFiles/103_prog/europython_2019/lvl${LVL}/cpp/cgo_lib)
LINK_DIRECTORIES(/Users/imhiro/AllFiles/103_prog/europython_2019/lvl${LVL}/cpp/cgo_lib)

add_executable(go_python_demo0 main.cpp cgo_lib/cgo_lib_lvl${LVL}.h)

TARGET_LINK_LIBRARIES(go_python_demo0 ${CMAKE_SOURCE_DIR}/cgo_lib/cgo_lib_lvl${LVL}.a)


add_executable(cgo_support cgo_support/lvl2_cgo_support.cpp cgo_support/lvl2_cgo_support.h cgo_support/main.cpp)


add_executable(cpp_personlib cpp_personlib/main.cpp cpp_personlib/person.cpp cpp_personlib/person.h)