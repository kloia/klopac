from pathlib import Path
from typing import List
import yaml
import pwd
import grp
import os
from collections import Mapping
from provisioner import logger
from provisioner.layer import Layer

# should we use the following package
# https://github.com/zerwes/hiyapyco
# https://gist.github.com/angstwad/bf22d1822c38a92ec0a9
# a: {a:1,b:2}
# b: {b:3, c:4}
# merged_dict = {**a, **b}
def dict_merge(dct, merge_dct):
    """ Recursive dict merge. Inspired by :meth:``dict.update()``, instead of
    updating only top-level keys, dict_merge recurses down into dicts nested
    to an arbitrary depth, updating keys. The ``merge_dct`` is merged into
    ``dct``.
    :param dct: dict onto which the merge is executed
    :param merge_dct: dct merged into dct
    :return: None
    """
    for k, v in merge_dct.items():
        if (k in dct and isinstance(dct[k], dict)
                and isinstance(merge_dct[k], Mapping)):
            dict_merge(dct[k], merge_dct[k])
        else:
            dct[k] = merge_dct[k]

def read_yaml_file(file_path: Path) -> dict:
    with open(file_path, "r") as f:
        try:
            yaml_obj = yaml.safe_load(f)
        except yaml.YAMLError as err:
            print(err)

    return yaml_obj

def write_yaml_file(yaml_obj: dict, file_path: Path):
    with open(file_path, "w") as f:
        try:
            yaml.safe_dump(yaml_obj, f)
        except Exception as err:
            logger.error(err)

def check_uid(uid: int) -> bool:
    try:
        pwd.getpwuid(uid)
        return True
    except KeyError:
        logger.error(f"{uid} uid doesn't exist")
        return False

def check_gid(gid: int) -> bool:
    try:
        grp.getgrgid(gid)
        return True
    except KeyError:
        logger.error(f"{gid} gid doesn't exist")
        return False

def check_uid_and_gid(uid: int, gid: int) -> bool:
    if check_uid(uid) or check_gid(gid):
        return True

    return False

def set_uid_and_gid(uid: int, gid: int, path: str):
    try:
        os.chown(path, uid, gid)
    except Exception as err:
        logger.error(err)

def read_layer_yamls(layers, vars_path: Path) -> List[Layer]:
    return [Layer(read_yaml_file(Path(vars_path, f"{layer}.yaml")), layer, shorthand) for layer, shorthand in layers]
   
def read_layer_defaults(layers: List[Layer], defaults_path: Path) -> List[dict]:
    logger.info("[*] Reading default files for layers")
    defaults = []

    for layer in layers:
        default_path = layer.get_default_path(defaults_path)
        logger.info(f"[*] {layer.name} defaults at path: {default_path}")

        defaults.append(read_yaml_file(default_path))

    return defaults

def merge_layer_and_default(defaults: List[dict], layers: List[Layer]) -> None:
    for layer, default in zip(layers, defaults):
        logger.info(f"[*] Merging {layer.name} defaults")
        dict_merge(layer.raw, default)

def include_layer(layer: Layer, yaml_to_merge, manifests_path: Path):
    if "branch" in layer.get_branch_or_version():
        manifest_path = Path(manifests_path, layer.runner_type, layer.get_branch_filename())
        dict_merge(yaml_to_merge, read_yaml_file(manifest_path))

    elif "branch" not in layer.get_branch_or_version() and "version" in layer.get_branch_or_version():
        manifest_path = Path(manifests_path, layer.runner_type, layer.get_version_filename())
        dict_merge(yaml_to_merge, read_yaml_file(manifest_path))
