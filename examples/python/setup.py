#!/usr/bin/env python

import os
import sys

from setuptools import setup, find_packages

os.chdir(os.path.dirname(sys.argv[0]) or ".")

setup(
    name="gohttp",
    version="0.3.3",
    description="Bindings for gohttplib, exposing Go's http.Server",
    long_description=open("README.rst", "rt").read(),
    url="https://github.com/shazow/gohttplib",
    author="Andrey Petrov",
    author_email="andrey.petrov@shazow.net",
    classifiers=[
        "Programming Language :: Python :: 2",
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
    ],
    packages=find_packages(),
    install_requires=["cffi>=1.0.0"],
    setup_requires=["cffi>=1.0.0"],
    cffi_modules=[
        "./gohttp/build_gohttplib.py:ffi",
    ],
    package_data={
        "gohttp": ["libgohttp.so"],
    },
)
