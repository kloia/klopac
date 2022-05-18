from typing import List
import yaml
import pwd
import grp
import os
from collections import Mapping

# should we use the following package
# https://github.com/zerwes/hiyapyco
# https://gist.github.com/angstwad/bf22d1822c38a92ec0a9
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

def read_yaml_file(file_path: str) -> dict:
    with open(file_path, "r") as f:
        try:
            yaml_obj = yaml.safe_load(f)
        except yaml.YAMLError as err:
            print(err)

    return yaml_obj

def write_yaml_file(yaml_obj: dict, file_path: str):
    with open(file_path, "w") as f:
        try:
            yaml.safe_dump(yaml_obj, f)
        except Exception as err:
            print(err)

def check_key(dict: dict, key: str) -> bool:
    if key in dict:
        return True
    return False

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

def include_layer(layer_obj: dict, layer_name: str, yaml_to_merge, manifests_path):
    manifest_path = f"{layer_name}_manifest_path"
    if check_key(layer_obj[layer_obj['type']], key='branch'):
        manifest_path = f"{manifests_path}/{layer_obj['runner']['type']}/{layer_obj['type']}@{layer_obj[layer_obj['type']]['branch']}.yaml"
        dict_merge(yaml_to_merge, read_yaml_file(manifest_path))

    if not check_key(layer_obj[layer_obj['type']], key='branch') and check_key(layer_obj[layer_obj['type']], key='version'):
        manifest_path = f"{manifests_path}/{layer_obj['runner']['type']}/{layer_obj['type']}-{layer_obj[layer_obj['type']]['version']}.yaml"
        dict_merge(yaml_to_merge, read_yaml_file(manifest_path))
