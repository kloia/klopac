import logging
import argparse
from pathlib import Path
from enum import Enum


class LogLevel(str, Enum):
    INFO = "info"
    DEBUG = "debug"


class Config:
    def __init__(self, data_path: Path) -> None:
        self.uid: int = 1000
        self.gid: int = 1000
        self.data_path: Path = data_path
        self.bundle_path: Path = Path(self.data_path, Path("bundle"))
        self.vars_path: Path = Path(self.data_path, Path("vars"))
        self.defaults_path: Path = Path(self.vars_path, Path("defaults"))
        self.repo_path: Path = Path(self.data_path, Path("repo"))
        self.manifests_path: Path = Path(self.data_path, Path("manifests"))
        self.layers = ["engine", "image", "instance"]
        self.shorthands = ["engine", "img", "ins"]

    @staticmethod
    def parse_args():
        parser = argparse.ArgumentParser()
        parser.add_argument(
            "--dev",
            action="store_true",
            help="set the current environment as development",
        )
        parser.add_argument(
            "--log-level", help="set log-level. possible values are 'info' and 'debug'"
        )
        parser.add_argument(
            "--data-path", help="set the data path to an arbitrary location"
        )
        return parser.parse_args()


def set_config() -> Config:
    args = Config.parse_args()
    data_path = Path("/data")

    if args.log_level == LogLevel.DEBUG:
        logger = logging.getLogger("provisioner")
        logger.setLevel(logging.DEBUG)

    if args.dev:
        data_path = Path.cwd().parent

    if args.data_path:
        data_path = Path(args.data_path)

    return Config(data_path)


config = set_config()
