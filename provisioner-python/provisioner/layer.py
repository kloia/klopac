from pathlib import Path
import provisioner.core as core
import sys
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
