import sys
from pathlib import Path
from provisioner.core import *
from provisioner.repo import *

uid=1000
gid=1000
data_path = ".."
bundle_path = f"{data_path}/bundle"
vars_path = f"{data_path}/vars"
repo_path = f"{data_path}/repo"
manifests_path = f"{data_path}/manifests"

if __name__ == "__main__":
    instance_yaml = read_yaml_file(f"{vars_path}/instance.yaml")
    engine_yaml = read_yaml_file(f"{vars_path}/engine.yaml")
    image_yaml = read_yaml_file(f"{vars_path}/image.yaml")
    platform_yaml = read_yaml_file(f"{vars_path}/platform.yaml")

    instance_defaults_yaml = read_yaml_file(f'{vars_path}/defaults/ins-{instance_yaml["ins"]["type"]}.yaml')
    image_defaults_yaml = read_yaml_file(f'{vars_path}/defaults/img-{image_yaml["img"]["type"]}.yaml')
    engine_defaults_yaml = read_yaml_file(f'{vars_path}/defaults/engine-{engine_yaml["engine"]["type"]}.yaml')

    dict_merge(instance_yaml, instance_defaults_yaml)
    dict_merge(image_yaml, image_defaults_yaml)
    dict_merge(engine_yaml, engine_defaults_yaml)

    engine = engine_yaml["engine"]
    image = image_yaml["img"]
    instance = instance_yaml["ins"]

    layers = ["engine", "image", "instance"]

    for layer in layers:
        layer_yaml = locals()[layer]
        include_layer(layer_yaml, layer, platform_yaml, manifests_path)

    platform = platform_yaml["platform"]
    repo_names = platform["repo"].keys()

    if check_empty_repo_uri(platform):
        sys.exit("You cannot have an empty repo name")

    #TODO: fix uid and gid checks not exiting
    for repo_name in repo_names:
        repo = platform["repo"][repo_name]
        repo_uri = get_repo_uri(repo)
        repo_name = get_repo_name(repo_uri)

        r_path = f"{repo_path}/{repo_name}"
        create_repo_dir(r_path, mode=0o777, exist_ok=True)
        # check_uid_and_gid(uid, gid)
        # set_uid_and_gid(uid, gid, path=r_path)

        if check_key(repo, key="branch"):
            clone_repo(repo_uri, r_path, branch=repo["branch"])

        if not check_key(repo, key="branch") and check_key(repo, key="version"):
            clone_repo(repo_uri, r_path, branch=repo["version"])

    for layer in layers:
        layer_yaml = locals()[layer]
        op = get_layer_operation(layer_yaml)
        enabled = check_layer_enabled(layer_yaml)
        repo_name = layer_yaml["type"]

        if not check_key(platform["repo"], repo_name):
          logger.info(f"[*] {repo_name} is not a repo")
          continue

        repo = klopac_repo(platform, repo_name)
        repo_enabled = repo["state"]["enabled"]
        rr_path = ""

        if repo_enabled:
            rr_path = repo["outputs"]["file"]["path"]

        logger.debug(f"Operation: {op} / Repo_path: {rr_path} / Repo: {repo_name} / State: {repo_enabled}")

        if op != "create" and repo_enabled and rr_path:
            logger.info("[*] Copying state files...")
            state_path = Path(bundle_path, rr_path)
            dest = Path(repo_path, repo_name)

            logger.info(f"[*] src: {state_path}, dest: {dest}")

            copy_state_file(state_path, dest)
