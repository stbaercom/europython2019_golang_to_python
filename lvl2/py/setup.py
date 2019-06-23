from distutils.core import setup
from Cython.Distutils import build_ext, Extension
import os

lvl = 2
libname = "cgo_lib/cgo_lib_lvl%s.a" % lvl
extention_name = "cgo_cython_lvl%s" % lvl


setup(
    name=extention_name,
    ext_modules=[
        Extension(extention_name,
                  sources=[extention_name + ".pyx"],
                  include_dirs=["cgo_lib"],
                  libraries=[libname ],
                  language="c++",
                  extra_objects=[libname],
                  )
    ],
    cmdclass={'build_ext': build_ext},

)
