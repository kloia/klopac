from setuptools import setup, find_packages

setup(
    name="klopac-validator",
    version="0.1-alpha",
    packages=find_packages(include=["validator"]),
    install_requires=["PyYAML", "rich"],
    extras_require={},
)
