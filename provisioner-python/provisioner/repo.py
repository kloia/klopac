import os
from pathlib import Path
import shutil
from provisioner import logger
from typing import List
import git
from git import RemoteProgress

class CloneProgress(RemoteProgress):
    def update(self, op_code, cur_count, max_count=None, message=''):
        if message:
            logger.info(message)

class Repo:
    def __init__(self, repo: dict, name: str) -> None:
        self.data = repo
        self.name = name
        self.uri = self.data["uri"]
        self.enabled = self.data["state"]["enabled"]
        self.layer = self.data["from_layer"]
        # self.download_path = self.data["outputs"]["file"]["path"]

    def get_remote_reponame(self) -> str:
        return self.uri.split("/")[-1].split(".")[0]

    def check_branch(self) -> bool:
        return "branch" in self.data

    def check_version(self) -> bool:
        return "branch" not in self.data and "version" in self.data

    def clone_repo(self, path: Path, branch: str):
        logger.info(f"Cloning into {self.get_remote_reponame()} from {self.uri}")
        try:
            git.Repo.clone_from(self.uri, path, branch=branch, progress=CloneProgress())
        except Exception as err:
            logger.error(err)

def get_repo_uris(platform: dict) -> List[str]:
    return [platform['repo'][repo_name]['uri'] for repo_name in platform['repo'].keys()]

def klopac_repo(platform: dict, repo_name: str) -> dict:
    return platform['repo'][repo_name]

def check_empty_repo_uri(platform: dict) -> bool:
    for repo_name in platform['repo'].keys():
        if not platform['repo'][repo_name]['uri']:
            return True

    return False

def create_repo_dir(dir_path: Path, mode, exist_ok: bool):
    try:
        os.makedirs(dir_path, mode=mode, exist_ok=exist_ok)
    except OSError as err:
        print(err)
    
def copy_state_file(src_path: Path, dest_path: Path):
    try:
        shutil.copy(src_path, dest_path)
    except Exception as err:
        logger.error(err)
