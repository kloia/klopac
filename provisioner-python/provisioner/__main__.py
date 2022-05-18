import sys
from provisioner.core import *
from provisioner.repo import *

uid=1000
gid=1000
data_path = ".."
vars_path = f"{data_path}/vars"
repo_path = f"{data_path}/repo"
manifests_path = f"{data_path}/manifests"

if __name__ == "__main__":
    instance_yaml = read_yaml_file(f"{vars_path}/instance.yaml")
    engine_yaml = read_yaml_file(f"{vars_path}/engine.yaml")
    image_yaml = read_yaml_file(f"{vars_path}/image.yaml")
    platform_yaml = read_yaml_file(f"{vars_path}/platform.yaml")

    instance_defaults_yaml = read_yaml_file(f"{vars_path}/defaults/ins-{instance_yaml['ins']['type']}.yaml")
    image_defaults_yaml = read_yaml_file(f"{vars_path}/defaults/img-{image_yaml['img']['type']}.yaml")
    engine_defaults_yaml = read_yaml_file(f"{vars_path}/defaults/engine-{engine_yaml['engine']['type']}.yaml")

    dict_merge(instance_yaml, instance_defaults_yaml)
    dict_merge(image_yaml, image_defaults_yaml)
    dict_merge(engine_yaml, engine_defaults_yaml)

    engine = engine_yaml['engine']
    image = image_yaml['img']
    instance = instance_yaml['ins']

    if check_key(engine[engine['type']], key='branch'):
        engine_manifest_path = f"{manifests_path}/{engine['runner']['type']}/{engine['type']}@{engine[engine['type']]['branch']}.yaml"
        dict_merge(platform_yaml, read_yaml_file(engine_manifest_path))

    if not check_key(engine[engine['type']], key='branch') and check_key(engine[engine['type']], key='version'):
        engine_manifest_path = f"{manifests_path}/{engine['runner']['type']}/{engine['type']}-{engine[engine['type']]['version']}.yaml"
        dict_merge(platform_yaml, read_yaml_file(engine_manifest_path))

    if check_key(image[image['type']], key='branch'):
        image_manifest_path = f"{manifests_path}/{image['runner']['type']}/{image['type']}@{image[image['type']]['branch']}.yaml"
        dict_merge(platform_yaml, read_yaml_file(image_manifest_path))

    if not check_key(image[image['type']], key='branch') and check_key(image[image['type']], key='version'):
        image_manifest_path = f"{manifests_path}/{image['runner']['type']}/{image['type']}-{image[image['type']]['version']}.yaml"
        dict_merge(platform_yaml, read_yaml_file(image_manifest_path))

    if check_key(instance[instance['type']], key='branch'):
        instance_manifest_path = f"{manifests_path}/{instance['runner']['type']}/{instance['type']}@{instance[instance['type']]['branch']}.yaml"
        dict_merge(platform_yaml, read_yaml_file(instance_manifest_path))

    if not check_key(instance[instance['type']], key='branch') and check_key(instance[instance['type']], key='version'):
        instance_manifest_path = f"{manifests_path}/{instance['runner']['type']}/{instance['type']}-{instance[instance['type']]['version']}.yaml"
        dict_merge(platform_yaml, read_yaml_file(instance_manifest_path))

    platform = platform_yaml['platform']
    repo_names = platform['repo'].keys()

    if check_empty_repo_uri(platform):
        sys.exit("You cannot have an empty repo name")

    #TODO: fix uid and gid checks not exiting
    for repo_name in repo_names:
        repo = platform['repo'][repo_name]
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

