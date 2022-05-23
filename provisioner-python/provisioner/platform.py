import sys
from provisioner import logger
from provisioner.repo import Repo, create_repo_dir
from provisioner.layer import Layer
from provisioner.config import *
from provisioner.core import dict_merge, read_yaml_file

class Platform:
    def __init__(self, platform: dict) -> None:
      self.raw = platform
      self.data = platform["platform"]

    def include_layer(self, layer: Layer, manifests_path: Path):
        from_layer_obj = {"platform":{"repo":{layer.type:{"from_layer":layer.name}}}}
        if "platform" in self.raw and layer.runner_type == "repo":
            filepath = layer.branch_or_version()
            manifest_path = Path(manifests_path, layer.runner_type, filepath)
            manifest_yaml = read_yaml_file(manifest_path)
            dict_merge(manifest_yaml, from_layer_obj)
            dict_merge(self.raw, manifest_yaml)

    def include_layers(self):
        for layer in Layer.get_layers():
            self.include_layer(layer, MANIFESTS_PATH)
    
    def set_repos(self):
        try:
            repos = []

            # check if the repo URI is empty and create a repo object if it is not
            for r in self.data["repo"].keys():
                if "uri" not in self.data["repo"][r]:
                    logger.warning(f"[*] empty uri for {r}")
                else:
                    repos.append(Repo(repo=self.data["repo"][r], name=r))

            self.repos = repos

        except KeyError:
            logger.error(f"[*] Your platform YAML does not have a 'repo' key.")
            sys.exit(1)

    def clone_repos(self) -> None:
        #TODO: fix uid and gid checks not exiting
        for repo in self.repos:
            r_path = Path(REPO_PATH, repo.get_remote_reponame())
            create_repo_dir(r_path, mode=0o777, exist_ok=True)
            # check_uid_and_gid(uid, gid)
            # set_uid_and_gid(uid, gid, path=r_path)
            repo.clone_repo(r_path, repo.branch_or_version())

    def copy_states(self) -> None:
        for repo in self.repos:
            layer = Layer.get_layer(repo.layer)
            download_path = Path("")

            if repo.enabled:
                download_path = Path(repo.data["outputs"]["file"]["path"])

            logger.debug(f"Operation: {layer.op} / Repo_path: {download_path} / Repo: {repo.name} / State: {repo.enabled}")

            if layer.op != "create" and repo.enabled and download_path:
                logger.info("[*] Copying state files...")
                state_path = Path(BUNDLE_PATH, download_path)
                repo.copy_state_file(state_path)
