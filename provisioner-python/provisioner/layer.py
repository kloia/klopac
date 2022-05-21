from pathlib import Path
import provisioner.core as core
from provisioner import logger

class Layer:
    def __init__(self, layer: dict, name: str, shorthand: str) -> None:
        self.raw = layer
        self.data = layer[shorthand]
        self.op = self.data['operation']['type']
        self.type = self.data['type']
        self.runner_type = self.data['runner']['type']
        self.enabled = self.data['enabled']
        self.shorthand = shorthand
        self.name = name

    def write_to_yaml(self, path: Path) -> None:
        core.write_yaml_file(self.data, path)

    def get_branch_or_version(self):
        # logger.debug(self.data[self.type])
        return self.data[self.type]

    def get_branch(self):
        return self.data[self.type]["branch"]

    def get_version(self):
        return self.data[self.type]["version"]

    def get_branch_filename(self) -> Path:
        return Path(f"{self.type}@{self.get_branch()}.yaml")

    def get_version_filename(self) -> Path:
        return Path(f"{self.type}-{self.get_version()}.yaml")

    def get_default_path(self, defaults_path: Path) -> Path:
        filepath = Path(f"{self.shorthand}-{self.type}.yaml")
        return Path(defaults_path, filepath)
