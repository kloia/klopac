from provisioner.core import *

if __name__ == "__main__":
    ins_yaml = read_yaml_file("instance.yml")
    engine_yaml = read_yaml_file("engine.yml")
    print(engine_yaml)
