def create_repo_names(platform: dict) -> list[str]:
    return [platform['repo'][repo_name]['uri'] for repo_name in platform['repo'].keys()]

def check_empty_repo_name(platform: dict) -> bool:
    for repo_name in platform['repo'].keys():
        if not platform['repo'][repo_name]['uri']:
            return True

    return False

