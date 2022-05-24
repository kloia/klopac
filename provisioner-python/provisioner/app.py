from pathlib import Path
import provisioner.core as core
from provisioner import logger
from provisioner.repo import Repo
from provisioner.layer import Layer
from provisioner.platform import Platform
from provisioner.config import *

class App:
    def __init__(self, config={}):
        self.config = config

    def run(self):
        try:
            # The iterator gets exhausted after first use so we want a lambda/list if we want to use a variable
            # We zip layers because some of the filenames and dictionary keys use the shorthand version
            zip_layers = lambda: zip(LAYERS, SHORTHANDS)
            
            # Read and create Layers from YAML files for different layers
            Layer.read_yamls(zip_layers())

            # Platform is not a layer so it is parsed separately
            platform = Platform(core.read_yaml_file(Path(VARS_PATH, "platform.yaml")))

            # Read default YAML files for layer types
            Layer.read_defaults()

            # Merge the Layers and defaults lists
            Layer.merge_layers_and_defaults()

            # Include the manifest YAMLs for different components to the platform object
            Layer.include_manifests(platform)

            # We can now set the repo object on our platform instance
            Repo.set_repos(platform)

            # Clone repos that need to be cloned
            Repo.clone_repos()

            # Copy state files to the necessary repos
            Repo.copy_states()

        except Exception as err:
            logger.error(err)