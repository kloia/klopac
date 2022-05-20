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
            print(err)

def check_uid(uid: int) -> bool:
    try:
        pwd.getpwuid(uid)
        return True
    except KeyError:
        print(f"{uid} uid doesn't exist")
        return False

def check_gid(gid: int) -> bool:
    try:
        grp.getgrgid(gid)
        return True
    except KeyError:
        print(f"{gid} gid doesn't exist")
        return False

def check_uid_and_gid(uid: int, gid: int) -> bool:
    if check_uid(uid) or check_gid(gid):
        return True

    return False

def set_uid_and_gid(uid: int, gid: int, path: str):
    try:
        os.chown(path, uid, gid)
    except Exception as err:
        print(err)

def read_layer_yamls(yamls: dict, layers: List[str], vars_path: Path) -> None:
    for layer in layers:
        file_path = f"{layer}.yaml"
        yamls[f"{layer}_yaml"] = read_yaml_file(Path(vars_path, file_path))
   
def read_layer_defaults(yamls: dict, defaults: dict, layers: zip, defaults_path: Path) -> None:
    logger.info("[*] Reading default files for layers")
    for layer, shorthand in layers:
        yaml_name = f"{layer}_yaml"
        type = yamls[yaml_name][shorthand]['type']
        file_path = f"{shorthand}-{type}.yaml"
        default_yaml_name = f"{layer}_defaults_yaml"
        default_path = Path(defaults_path, file_path)
        logger.info(f"[*] {layer} defaults as key: {default_yaml_name} at path: {default_path}")

        defaults[default_yaml_name] = read_yaml_file(default_path)

def merge_layer_and_default(yamls: dict, defaults: dict, layers: List[str]) -> None:
    for layer in layers:
        layer_yaml = f"{layer}_yaml"
        default_yaml = f"{layer}_defaults_yaml"
        yaml = yamls[layer_yaml]
        default = defaults[default_yaml]

        logger.info(f"[*] Merging {layer} defaults")

        dict_merge(yaml, default)

def include_layer(layer_obj: dict, yaml_to_merge, manifests_path: Path):
    if "branch" in layer_obj[layer_obj['type']]:
        branch_fname = f"{layer_obj['type']}@{layer_obj[layer_obj['type']]['branch']}.yaml"
        manifest_path = Path(manifests_path, layer_obj['runner']['type'], branch_fname)
        dict_merge(yaml_to_merge, read_yaml_file(manifest_path))

    if "branch" not in layer_obj[layer_obj['type']] and "version" in layer_obj[layer_obj['type']]:
        version_fname = f"{layer_obj['type']}-{layer_obj[layer_obj['type']]['version']}.yaml"
        manifest_path = Path(manifests_path, layer_obj['runner']['type'], version_fname)
        dict_merge(yaml_to_merge, read_yaml_file(manifest_path))
