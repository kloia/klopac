import os
import sys
import shutil
import git
from pathlib import Path
from provisioner import logger
from git import RemoteProgress
from provisioner.config import *
from provisioner.layer import Layer
from provisioner.platform import Platform

class CloneProgress(RemoteProgress):
    def update(self, op_code, cur_count, max_count=None, message=''):
        if message:
            logger.info(message)

class Repo:
    __repos = []

    def __init__(self, repo: dict, name: str) -> None:
        self.data = repo
        self.name = name
        self.uri = self.data["uri"]
        self.enabled = self.data["state"]["enabled"]
        self.layer = self.data["from_layer"]
        # self.state_path = self.data["outputs"]["file"]["path"]

    @classmethod
    def get_repos(cls):
        return cls.__repos

    @classmethod
    def set_repos(cls, platform: Platform):
        try:
            # check if the repo URI is empty and create a repo object if it is not
            for r in platform.data["repo"].keys():
                if "uri" not in platform.data["repo"][r]:
                    logger.warning(f"[*] empty uri for {r}")
                else:
                    cls.__repos.append(Repo(repo=platform.data["repo"][r], name=r))
        except KeyError:
            logger.error(f"[*] Your platform YAML does not have a 'repo' key.")
            sys.exit(1)

    @classmethod
    def clone_repos(cls) -> None:
        #TODO: fix uid and gid checks not exiting
        for repo in Repo.get_repos():
            r_path = Path(REPO_PATH, repo.get_remote_reponame())
            Repo.create_dir(r_path, mode=0o777, exist_ok=True)
            # check_uid_and_gid(uid, gid)
            # set_uid_and_gid(uid, gid, path=r_path)
            repo.clone_repo(r_path, repo.branch_or_version())

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

    @staticmethod
    def check_empty_uri(platform: dict) -> bool:
        for repo_name in platform['repo'].keys():
            if not platform['repo'][repo_name]['uri']:
                return True

        return False

    @staticmethod
    def create_dir(dir_path: Path, mode, exist_ok: bool):
        try:
            os.makedirs(dir_path, mode=mode, exist_ok=exist_ok)
        except OSError as err:
            print(err)
    
    @staticmethod
    def copy_states() -> None:
        for repo in Repo.get_repos():
            layer = Layer.get_layer(repo.layer)
            download_path = Path("")

            if repo.enabled:
                download_path = Path(repo.data["outputs"]["file"]["path"])

            logger.debug(f"Operation: {layer.op} / Repo_path: {download_path} / Repo: {repo.name} / State: {repo.enabled}")

            if layer.op != "create" and repo.enabled and download_path:
                logger.info("[*] Copying state files...")
                state_path = Path(BUNDLE_PATH, download_path)
                repo.copy_state_file(state_path)
