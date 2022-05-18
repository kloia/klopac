import os

def get_repo_uris(platform: dict) -> list[str]:
    return [platform['repo'][repo_name]['uri'] for repo_name in platform['repo'].keys()]

def get_repo_name(uri: str) -> str:
    return uri.split("/")[-1].split(".")[0]

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
    
