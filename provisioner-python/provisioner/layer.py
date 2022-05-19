from pathlib import Path
from provisioner.core import write_yaml_file

class Layer:
    def __init__(self, layer_obj: dict) -> None:
        self.data = layer_obj
        self.op = self.data['operation']['type']
        self.type = self.data['type']
        self.enabled = self.data['enabled']

    def write_to_yaml(self, path: Path) -> None:
        write_yaml_file(self.data, path)

