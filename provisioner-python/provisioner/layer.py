from __future__ import annotations

from pathlib import Path
import sys
import provisioner.core as core
from provisioner.config import *
from provisioner import logger
from typing import List

class Layer:
    __layers = []
    __defaults = []

    def __init__(self, layer: dict, name: str, shorthand: str) -> None:
        self.raw = layer
        self.data = layer[shorthand]
        self.op = self.data['operation']['type']
        self.type = self.data['type']
        self.runner_type = self.data['runner']['type']
        self.enabled = self.data['enabled']
        self.shorthand = shorthand
        self.name = name

        self.__layers.append(self)

    @classmethod
    def read_layer_defaults(cls):
        logger.info("[*] Reading default files for layers")

        for layer in Layer.get_layers():
            default_path = layer.get_default_path(DEFAULTS_PATH)
            logger.info(f"[*] {layer.name} defaults at path: {default_path}")
            cls.__defaults.append(core.read_yaml_file(default_path))

    @classmethod
    def merge_layer_and_default(cls):
        for layer, default in zip(Layer.get_layers(), Layer.get_defaults()):
            logger.info(f"[*] Merging {layer.name} defaults...")
            core.dict_merge(layer.raw, default)

    @classmethod
    def get_layer(cls, name: str) -> Layer:
        for layer in cls.__layers:
            if layer.name == name: return layer

        logger.error(f"No layer with name {name} exists")
        sys.exit(1)
        
    @classmethod
    def get_layers(cls) -> List[Layer]:
        return cls.__layers

    @classmethod
    def get_defaults(cls) -> List[dict]:
        return cls.__defaults
            
    def write_to_yaml(self, path: Path) -> None:
        core.write_yaml_file(self.data, path)

    def branch_or_version(self):
        if "branch" in self.data[self.type]:
            return Path(f"{self.type}@{self.get_branch()}.yaml")
        elif "version" in self.data[self.type]:
            return Path(f"{self.type}-{self.get_version()}.yaml")
        else:
            logger.error(f"[*] No branch or version found for repo: {self.data[self.type]}")
            sys.exit(1)

    def get_branch(self):
        try:
            return self.data[self.type]["branch"]
        except KeyError:
            logger.error(f"[*] Type mismatch: cannot find a branch for repo {self.type}")
            sys.exit(1)

    def get_version(self):
        try:
            return self.data[self.type]["version"]
        except KeyError:
            logger.error(f"[*] Type mismatch: cannot find a version for repo {self.type}")
            sys.exit(1)

    def get_default_path(self, defaults_path: Path) -> Path:
        filepath = Path(f"{self.shorthand}-{self.type}.yaml")
        return Path(defaults_path, filepath)
