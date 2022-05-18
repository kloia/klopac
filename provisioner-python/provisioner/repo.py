import os
import time
import git
from git import RemoteProgress

class CloneProgress(RemoteProgress):
    def update(self, op_code, cur_count, max_count=None, message=''):
        if message:
            print(message)

def clone_repo(repo_uri: str, repo_path: str, branch: str):
    repo_name = get_repo_name(repo_uri)
    print(f"Cloning into {repo_name} from {repo_uri}")
    git.Repo.clone_from(repo_uri, repo_path, branch=branch, progress=CloneProgress())

def get_repo_uris(platform: dict) -> list[str]:
    return [platform['repo'][repo_name]['uri'] for repo_name in platform['repo'].keys()]

def get_repo_name(uri: str) -> str:
    return uri.split("/")[-1].split(".")[0]

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
    
