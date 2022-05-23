from pathlib import Path
from provisioner.core import *
from provisioner.repo import *
from provisioner.layer import Layer
from provisioner.platform import Platform
from provisioner.config import *

class App:
    def __init__(self, config={}):
        self.config = config

    def run(self):
        # The iterator gets exhausted after first use so we want a lambda/list if we want to use a variable
        # We zip layers because some of the filenames and dictionary keys use the shorthand version
        zip_layers = lambda: zip(LAYERS, SHORTHANDS)
        
        # Read YAML files for different layers
        read_layer_yamls(zip_layers())

        # Platform is not a layer so it is parsed separately
        platform = Platform(read_yaml_file(Path(VARS_PATH, "platform.yaml")))

        # Read default YAML files for layer types
        Layer.read_layer_defaults()

        # Merge the Layers and defaults lists
        Layer.merge_layer_and_default()

        # Include the manifest YAMLs for different components to the platform object
        platform.include_layers()

        # We can now set the repo object on our platform instance
        Repo.set_repos(platform)

        # Clone repos that need to be cloned
        Repo.clone_repos()

        # Copy state files to the necessary repos
        Repo.copy_states()
