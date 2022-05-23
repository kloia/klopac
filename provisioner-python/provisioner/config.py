from pathlib import Path

uid = gid = 1000
DATA_PATH = Path.cwd().parent
BUNDLE_PATH = Path(DATA_PATH, Path("bundle"))
VARS_PATH = Path(DATA_PATH, Path("vars"))
DEFAULTS_PATH = Path(VARS_PATH, Path("defaults"))
REPO_PATH = Path(DATA_PATH, Path("repo"))
MANIFESTS_PATH = Path(DATA_PATH, Path("manifests"))
LAYERS = ["engine", "image", "instance"]
SHORTHANDS = ["engine", "img", "ins"]
