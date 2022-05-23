import sys
from pathlib import Path
from provisioner.core import *
from provisioner.repo import *
from provisioner.layer import Layer
from provisioner.config import *

class App:
    def __init__(self, config={}):
        self.config = config

    def run(self):
        layers = ["engine", "image", "instance"]
        shorthands = ["engine", "img", "ins"]

        #the iterator gets exhausted after first use so we want a lambda/list if we want to use a variable
        l = lambda: zip(layers, shorthands)
        
        engine, image, instance = layer_objs = read_layer_yamls(l(), vars_path)
        platform_yaml = read_yaml_file(Path(vars_path, "platform.yaml"))

        defaults = read_layer_defaults(layer_objs, defaults_path)
        merge_layer_and_default(defaults, layer_objs)

        for layer in layer_objs:
            include_layer(layer, platform_yaml, manifests_path)

        platform = platform_yaml["platform"]

        repos = []
        for r in platform["repo"].keys():
            if "uri" not in platform["repo"][r]:
                logger.warning(f"[*] empty uri for {r}")
            else:
                repos.append(Repo(repo=platform["repo"][r], name=r))
        
        #TODO: fix uid and gid checks not exiting
        for repo in repos:
            r_path = Path(repo_path, repo.get_remote_reponame())
            create_repo_dir(r_path, mode=0o777, exist_ok=True)
            # check_uid_and_gid(uid, gid)
            # set_uid_and_gid(uid, gid, path=r_path)
            repo.clone_repo(r_path, repo.branch_or_version())

        for repo in repos:
            layer = locals()[repo.layer]
            download_path = Path("")
            logger.debug(f"Operation: {layer.op} / Repo_path: {download_path} / Repo: {repo.name} / State: {repo.enabled}")

            if repo.enabled:
                download_path = Path(repo.data["outputs"]["file"]["path"])

            if layer.op != "create" and repo.enabled and download_path:
                logger.info("[*] Copying state files...")
                state_path = Path(bundle_path, download_path)
                repo.copy_state_file(state_path)
