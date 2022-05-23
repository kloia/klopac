from provisioner import logger
from provisioner.layer import Layer
from provisioner.config import *
from provisioner.core import dict_merge, read_yaml_file

class Platform:
    def __init__(self, platform: dict) -> None:
        self.raw = platform
        self.data = platform["platform"]

    def include_layer(self, layer: Layer, manifests_path: Path):
        from_layer_obj = {"platform":{"repo":{layer.type:{"from_layer":layer.name}}}}
        if "platform" in self.raw and layer.runner_type == "repo":
            filepath = layer.branch_or_version()
            manifest_path = Path(manifests_path, layer.runner_type, filepath)
            manifest_yaml = read_yaml_file(manifest_path)
            dict_merge(manifest_yaml, from_layer_obj)
            dict_merge(self.raw, manifest_yaml)

    def include_layers(self):
        for layer in Layer.get_layers():
            self.include_layer(layer, MANIFESTS_PATH)
