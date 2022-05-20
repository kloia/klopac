from pathlib import Path
import provisioner.core as core

class Layer:
    def __init__(self, layer: dict) -> None:
        self.data = layer
        self.op = self.data['operation']['type']
        self.type = self.data['type']
        self.enabled = self.data['enabled']

    def write_to_yaml(self, path: Path) -> None:
        core.write_yaml_file(self.data, path)

