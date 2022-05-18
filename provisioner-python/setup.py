from setuptools import setup, find_packages

setup(
    name='klopac-provisioner',
    version='0.1-alpha',
    packages=find_packages(include=['provisioner']),
    install_requires=[
        'PyYAML',
    ],
    extras_require={
    }
)
