from __future__ import annotations

import logging
import validator.core as core
from pathlib import Path
from typing import List
from validator.config import config
from validator.platform import Platform


class Layer:
    __layers = []
    __defaults = []

    def __init__(self, layer: dict, name: str, shorthand: str) -> None:
        self.raw = layer
        self.data = layer[shorthand]
        self.shorthand = shorthand
        self.name = name
        self.type = self.data["type"]
        self.runner_type = self.data["runner"]["type"]
        self.enabled = self.data["enabled"]
        self.default_path = Path(
            config.defaults_path, f"{self.shorthand}-{self.type}.yaml"
        )

        self.layers.append(self)

    @classmethod
    @property
    def layers(cls) -> List[Layer]:
        return cls.__layers

    @classmethod
    @property
    def defaults(cls) -> List[dict]:
        return cls.__defaults

    @property
    def manifest_path(self):
        return self._manifest_path

    """Getter and setter for manifest_version
    The value depends on the value of the defaults for the specific layers.
    If "branch" exists this will be used, otherwise we look for "version".
    If none are found there must be a problem with our config.
    (!) Also sets "manifest_path" to the correct filepath for the manifest file
    
    layer_type: This refers to the "self.data[self.type]" object 
    ex. for the engine layer this can be -> self.data["k3s"] or self.data["rke2"]
    """

    @property
    def manifest_version(self):
        return self._manifest_version

    # TODO: check what happens when the runnner_type is not "repo" and we try to access branch/manifest_path
    @manifest_version.setter
    def manifest_version(self, layer_type: dict):
        # We don't want to include the manifest if it is not the repo type
        if self.runner_type != "repo":
            return

        if "branch" in layer_type:
            self._manifest_version = layer_type["branch"]
            self._manifest_path = Path(
                config.manifests_path,
                self.runner_type,
                f"{self.type}@{self.manifest_version}.yaml",
            )
        elif "version" in layer_type:
            self._manifest_version = layer_type["version"]
            self._manifest_path = Path(
                config.manifests_path,
                self.runner_type,
                f"{self.type}-{self.manifest_version}.yaml",
            )
        else:
            raise ValueError(f"[!] No branch or version exists for repo {self.name}")

    """DEFAULTS METHODS"""
    # Reads the defaults for each Layer and adds the dictionary to __defaults
    @classmethod
    def read_defaults(cls):
        logging.info("[*] Reading default files for layers")

        for layer in Layer.layers:
            logging.info(f"[*] {layer.name} defaults at path: {layer.default_path}")
            cls.defaults.append(core.read_yaml_file(layer.default_path))

    """Merges the Layer object with its default values
    After this call the layer object will contain the default values it uses
    """

    @staticmethod
    def merge_layers_and_defaults():
        for layer, default in zip(Layer.layers, Layer.defaults):
            logging.info(f"[*] Merging {layer.name} defaults...")
            core.dict_merge(layer.raw, default)

            # We can safely try to set the "manifest_version" now after the merge
            layer.manifest_version = layer.data[layer.type]

    """MANIFEST METHODS"""

    @staticmethod
    def include_manifests(platform: Platform):
        for layer in Layer.layers:
            layer.include_manifest(platform)

    # Merges the Layer manifest with the Platform object
    def include_manifest(self, platform: Platform):
        from_layer_obj = {"platform": {"repo": {self.type: {"from_layer": self.name}}}}

        if self.runner_type == "repo":
            manifest_yaml = core.read_yaml_file(self.manifest_path)
            logging.info(
                f"[*] Merging {self.name} repo manifest at: {self.manifest_path}"
            )

            # This merge with from_layer_obj is necessary to access which layer a repo belongs to
            core.dict_merge(manifest_yaml, from_layer_obj)
            core.dict_merge(platform.raw, manifest_yaml)

    """UTILITY METHODS"""
    # Reads and creates Layer objects from YAMLs
    @staticmethod
    def read_yamls(layers_zip: zip) -> None:
        logging.info("[*] Reading layer YAMLs")

        for layer, shorthand in layers_zip:
            yaml_path = Path(config.vars_path, f"{layer}.yaml")
            logging.info(f"[*] {layer} at path: {yaml_path}")
            Layer(core.read_yaml_file(yaml_path), layer, shorthand)

    @staticmethod
    def get_layer(name: str) -> Layer:
        for layer in Layer.layers:
            if layer.name == name:
                return layer

        raise ValueError(f"[!] No layer with name {name} exists")

    # Utility function for debugging purposes
    def write_to_yaml(self, path: Path) -> None:
        core.write_yaml_file(self.data, path)
