import os
import sys
from pathlib import Path
import shutil
from provisioner import logger
from typing import List
import git
from git import RemoteProgress
from provisioner.config import REPO_PATH

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
        # self.state_path = self.data["outputs"]["file"]["path"]

    def get_remote_reponame(self) -> str:
        return self.uri.split("/")[-1].split(".")[0]

    def branch_or_version(self) -> str:
        if "branch" in self.data:
            return self.data["branch"]
        elif "version" in self.data:
            return self.data["version"]
        else:
            logger.error(f"[*] No branch or version found for repo: {self.name}")
            sys.exit(1)

    def clone_repo(self, path: Path, branch: str):
        logger.info(f"Cloning into {self.get_remote_reponame()} from {self.uri}")
        try:
            git.Repo.clone_from(self.uri, path, branch=branch, progress=CloneProgress())
        except Exception as err:
            logger.error(f"Something went wrong when cloning the repo. Make sure the repos do not exist already")
            logger.debug(err)
            sys.exit(1)

    def copy_state_file(self, src: Path):
        dest = Path(REPO_PATH, self.get_remote_reponame())
        logger.info(f"[*] src: {src}, dest: {dest}")

        try:
            shutil.copy(src, dest)
        except Exception as err:
            logger.error(err)

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
    
