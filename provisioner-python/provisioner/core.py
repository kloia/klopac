import yaml
import pwd
import grp
import os
from pathlib import Path
from collections import Mapping
from provisioner import logger

# should we use the following package
# https://github.com/zerwes/hiyapyco
# https://gist.github.com/angstwad/bf22d1822c38a92ec0a9
# a: {a:1,b:2}
# b: {b:3, c:4}
# merged_dict = {**a, **b}
def dict_merge(dct, merge_dct):
    """Recursive dict merge. Inspired by :meth:``dict.update()``, instead of
    updating only top-level keys, dict_merge recurses down into dicts nested
    to an arbitrary depth, updating keys. The ``merge_dct`` is merged into
    ``dct``.
    :param dct: dict onto which the merge is executed
    :param merge_dct: dct merged into dct
    :return: None
    """
    for k, v in merge_dct.items():
        if k in dct and isinstance(dct[k], dict) and isinstance(merge_dct[k], Mapping):
            dict_merge(dct[k], merge_dct[k])
        else:
            dct[k] = merge_dct[k]


# Reads a YAML file as a dictionary
def read_yaml_file(filepath: Path) -> dict:
    with open(filepath, "r") as f:
        try:
            yaml_obj = yaml.safe_load(f)
        except yaml.YAMLError as err:
            logger.debug(err)
            logger.error(f"There was an error reading {filepath}")
            raise yaml.YAMLError
        except FileNotFoundError as fnf_error:
            raise FileNotFoundError(fnf_error)

    return yaml_obj


# Writes a dictionary to a file in YAML format
def write_yaml_file(yaml_obj: dict, filepath: Path):
    with open(filepath, "w") as f:
        try:
            yaml.safe_dump(yaml_obj, f)
        except Exception as err:
            raise Exception(err)


def check_uid(uid: int) -> bool:
    try:
        pwd.getpwuid(uid)
        return True
    except KeyError:
        logger.error(f"uid: {uid} uid doesn't exist")
        return False


def check_gid(gid: int) -> bool:
    try:
        grp.getgrgid(gid)
        return True
    except KeyError:
        logger.error(f"gid: {gid} gid doesn't exist")
        return False


def set_uid_and_gid(uid: int, gid: int, path: Path) -> None:
    try:
        os.chown(path, uid, gid)
    except Exception as err:
        logger.debug(err)
        raise Exception(f"Cannot set uid: {uid} and gid: {gid}")


def create_dir(dir_path: Path, mode, exist_ok: bool):
    try:
        os.makedirs(dir_path, mode=mode, exist_ok=exist_ok)
    except OSError as err:
        logger.debug(err)
        raise OSError(f"Could not create the directory {dir_path}")
