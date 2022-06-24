# repo config validations
- Check if repo manifest file exists -> Is this necessary since Provisioner handles this?
- Check if the terraform state file exists (platform.repo.terraform.state.file.path) -> Is this necessary since Provisioner handles this?
# runner validations
- Need to check if a layer is enabled before validating it
- For a given layer if it is enabled we should check its /executor/repo/runner yaml files
- Check if a given runner version exists on dockerhub
- For the klopac runner itself, we should check if the provided version exists on dockerhub

# compatibility validations
- For AWS, check if the given user has permissions to create necessary resources

# input/output flow
- If a layer enabled, check if the necessary prerequirements are provided
- Check the "required" field for each layer if they are enabled
- For example, if you enabled the engine layer, you must provide the "required" fields for the instance layer

# state validation
- Check if the state files are in the correct place if the state is enabled
- State drift?
