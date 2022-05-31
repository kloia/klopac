import logging
import shutil
from git.repo import Repo as GitRepo
from pathlib import Path
from git import RemoteProgress
from provisioner.config import config
from provisioner.core import check_gid, check_uid, set_uid_and_gid
from provisioner.layer import Layer
from provisioner.platform import Platform


class CloneProgress(RemoteProgress):
    def update(self, op_code, cur_count, max_count=None, message=""):
        if message:
            logging.info(message)


class Repo:
    __repos = []

    def __init__(self, repo: dict, name: str) -> None:
        self.data = repo
        self.name = name
        self.uri = repo["uri"]
        self.remote_name = self.uri.split("/")[-1].split(".")[0]
        self.layer = repo["from_layer"]
        self.state_enabled = repo["state"]["enabled"]
        self.branch = repo
        self.state_path = repo

    @classmethod
    @property
    def repos(cls):
        return cls.__repos

    @classmethod
    def set_repos(cls, platform: Platform):
        try:
            # check if the repo URI is empty and create a repo object if it is not
            for repo in platform.data["repo"].keys():
                logging.info(f"[*] Adding the following repo: {repo}")
                cls.__repos.append(Repo(repo=platform.data["repo"][repo], name=repo))
        except KeyError as err:
            logging.error(err)
            raise KeyError(f"[!] There was an error setting up the repo for {repo}.")

    """Getter and setter for branch
    The exact value depends on whether "branch" or "version" is set inside the manifest file
    If "branch" is set this will be used, otherwise we look for "version"
    If none are found there must be something wrong with our config.
    """

    @property
    def branch(self):
        return self._branch

    @branch.setter
    def branch(self, repo):
        if "branch" in repo:
            self._branch = repo["branch"]
        elif "version" in repo:
            self._branch = repo["version"]
        else:
            raise ValueError(f"[!] No branch or version found for repo: {self.name}")

    """Getter and setter for state_path
    If "state_enabled" is true we set state_path, otherwise it should be None
    If the path to the state does not exist we set "state_path" to None
    """

    @property
    def state_path(self):
        return self._state_path

    @state_path.setter
    def state_path(self, repo):
        try:
            if self.state_enabled:
                self._state_path = Path(
                    config.bundle_path, repo["outputs"]["file"]["path"]
                )
            else:
                self._state_path = None
        except KeyError:
            raise KeyError(
                f"[!] {self.name} repo does not have a state file but its [state] is ENABLED"
            )

    """CLONE METHODS"""

    @classmethod
    def clone_repos(cls) -> None:
        # TODO: uid and gid checks
        for repo in cls.repos:
            r_path = Path(config.repo_path, repo.remote_name)
            repo.clone_repo(r_path)
            # if check_uid(uid) and check_gid(gid):
            #     set_uid_and_gid(uid, gid, path=r_path)

    def clone_repo(self, path: Path):
        logging.info(f"[*] Cloning into {self.remote_name} from {self.uri}")
        try:
            GitRepo.clone_from(
                self.uri, path, branch=self.branch, progress=CloneProgress()
            )
        except Exception as err:
            logging.debug(err)
            raise Exception(
                f"[!] Something went wrong while cloning the repo. Make sure the repo folders do not exist already"
            )

    """STATE METHODS"""

    @staticmethod
    def copy_states() -> None:
        for repo in Repo.repos:
            layer = Layer.get_layer(repo.layer)
            logging.debug(
                f"Operation: {layer.op} / Repo_path: {repo.state_path} / Repo: {repo.name} / State: {repo.state_enabled}"
            )

            if layer.op != "create" and repo.state_enabled and repo.state_path:
                logging.info("[*] Copying state files...")
                repo.copy_state_file()

    def copy_state_file(self):
        dest = Path(config.repo_path, self.remote_name)
        logging.info(f"[*] src: {self.state_path}, dest: {dest}")

        try:
            if self.state_path:
                shutil.copy(self.state_path, dest)
        except Exception as err:
            logging.debug(err)
            raise Exception(f"Error while copying the state file")
