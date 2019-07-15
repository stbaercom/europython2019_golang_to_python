from distutils.core import setup
from Cython.Distutils import build_ext, Extension

libname = "cgo_lib/cgo_lib.a"
extention_name = "cgo_cython"
setup(
    name=extention_name,
    ext_modules=[
        Extension(extention_name,
                  sources=[extention_name + ".pyx"],
                  include_dirs=["cgo_lib","../go/src/com.stbaer/demo_cgo/"],                  
                  libraries=[libname],
                  language="c++",
                  extra_objects=[libname],
                  )
    ], cmdclass={'build_ext': build_ext},
)

