# Provisioner:
This is the part of klopac that is responsible for cloning the repos and copying the state files into the correct paths if a bundle already exists.

This code uses the [Black](https://github.com/psf/black) standard and should be formatted accordingly.

# Usage:
[!] If you want to test the code please run it with the **"--dev"** flag. This will set the data_path to the Parent folder.

```python3 setup.py install``` then change directories to the "provisioner-python" and simply run ```python3 -m provisioner```

or

```python -m venv venv```
```source venv/bin/activate```
```python -m pip install -r requirements.txt```

then after changing directories to the "provisioner-python" folder, simply run:

```python -m provisioner```
