import sys
from pathlib import Path
from provisioner.core import *
from provisioner.repo import *
from provisioner.layer import Layer

uid = gid = 1000
data_path = Path.cwd().parent
bundle_path = Path(data_path, Path("bundle"))
vars_path = Path(data_path, Path("vars"))
defaults_path = Path(vars_path, Path("defaults"))
repo_path = Path(data_path, Path("repo"))
manifests_path = Path(data_path, Path("manifests"))

if __name__ == "__main__":
    yamls = {}; defaults = {}

    layers = ["engine", "image", "instance"]
    shorthands = ["engine", "img", "ins"]

    #the iterator gets exhausted after first use so we want a lambda/list if we want to use a variable
    l = lambda: zip(layers, shorthands)
    
    read_layer_yamls(yamls, layers, vars_path)
    platform_yaml = read_yaml_file(Path(vars_path, "platform.yaml"))

    engine, image, instance = layer_objs = [Layer(yamls[f"{layer}_yaml"][shorthand]) for layer, shorthand in l()]

    read_layer_defaults(yamls, defaults, l(), defaults_path)

    merge_layer_and_default(yamls, defaults, layers)

    for layer in layer_objs:
        include_layer(layer.data, platform_yaml, manifests_path)

    platform = platform_yaml["platform"]
    repo_names = platform["repo"].keys()

    if check_empty_repo_uri(platform):
        sys.exit("You cannot have an empty repo name")

    #TODO: fix uid and gid checks not exiting
    for repo_name in repo_names:
        repo = platform["repo"][repo_name]
        repo_uri = get_repo_uri(repo)
        repo_name = get_repo_name(repo_uri)

        r_path = Path(repo_path, repo_name)
        create_repo_dir(r_path, mode=0o777, exist_ok=True)
        # check_uid_and_gid(uid, gid)
        # set_uid_and_gid(uid, gid, path=r_path)

        if check_key(repo, key="branch"):
            clone_repo(repo_uri, r_path, branch=repo["branch"])

        if not check_key(repo, key="branch") and check_key(repo, key="version"):
            clone_repo(repo_uri, r_path, branch=repo["version"])

    for layer in layers:
        layer = locals()[layer]
        op = layer.op
        enabled = layer.enabled
        repo_name = layer.type

        if not check_key(platform["repo"], repo_name):
          logger.info(f"[*] {repo_name} is not a repo")
          continue

        repo = klopac_repo(platform, repo_name)
        repo_enabled = repo["state"]["enabled"]
        rr_path = Path("")

        if repo_enabled:
            rr_path = Path(repo["outputs"]["file"]["path"])

        logger.debug(f"Operation: {op} / Repo_path: {rr_path} / Repo: {repo_name} / State: {repo_enabled}")

        if op != "create" and repo_enabled and rr_path:
            logger.info("[*] Copying state files...")
            state_path = Path(bundle_path, rr_path)
            dest = Path(repo_path, repo_name)

            logger.info(f"[*] src: {state_path}, dest: {dest}")

            copy_state_file(state_path, dest)
