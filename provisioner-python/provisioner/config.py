from pathlib import Path

uid = gid = 1000
data_path = Path.cwd().parent
bundle_path = Path(data_path, Path("bundle"))
vars_path = Path(data_path, Path("vars"))
defaults_path = Path(vars_path, Path("defaults"))
repo_path = Path(data_path, Path("repo"))
manifests_path = Path(data_path, Path("manifests"))
