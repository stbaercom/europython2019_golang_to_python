from distutils.core import setup

from Cython.Distutils import build_ext, Extension

setup(
    name='cgo_person',
    ext_modules=[
        Extension("cgo_person",
                  sources=["person.pyx"],
                  include_dirs=["cgo_lib"],
                  libraries=["cgo_lib/main.a", ],
                  language="c++",
                  extra_objects=["cgo_lib/main.a"],
                  )
    ],
    cmdclass={'build_ext': build_ext},

)
