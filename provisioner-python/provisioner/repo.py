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

def clone_repo(repo_uri: str, repo_path: str, branch: str):
    repo_name = get_repo_name(repo_uri)
    logger.info(f"Cloning into {repo_name} from {repo_uri}")
    try:
        git.Repo.clone_from(repo_uri, repo_path, branch=branch, progress=CloneProgress())
    except Exception as err:
        logger.error(err)

def get_repo_uris(platform: dict) -> List[str]:
    return [platform['repo'][repo_name]['uri'] for repo_name in platform['repo'].keys()]

def get_repo_name(uri: str) -> str:
    return uri.split("/")[-1].split(".")[0]

def klopac_repo(platform: dict, repo_name: str) -> dict:
    return platform['repo'][repo_name]

def get_repo_uri(repo: dict) -> str:
    return repo['uri']

def check_empty_repo_uri(platform: dict) -> bool:
    for repo_name in platform['repo'].keys():
        if not platform['repo'][repo_name]['uri']:
            return True

    return False

def create_repo_dir(dir_path: str, mode, exist_ok: bool):
    try:
        os.makedirs(dir_path, mode=mode, exist_ok=exist_ok)
    except OSError as err:
        print(err)
    
def copy_state_file(src_path: Path, dest_path: Path):
    try:
        shutil.copy(src_path, dest_path)
    except Exception as err:
        logger.error(err)
