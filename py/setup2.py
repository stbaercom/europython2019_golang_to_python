from distutils.core import setup
from Cython.Distutils import build_ext, Extension


extention_name = "cpp_cython"
setup(
    name=extention_name,
    ext_modules=[
        Extension(extention_name,
                  sources=[extention_name + ".pyx"],
                  include_dirs=["cgo_lib"],

                  language="c++",

                  )
    ], cmdclass={'build_ext': build_ext},
)

