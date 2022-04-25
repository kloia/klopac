## Name: platform_validations

1. Check platform provider auth type
2. Check provider name (azure or aws)
3. Check provider type (eks or ec2 for aws, vm or vmss for azure)
4. Check if access key and secret key are defined for aws
5. Check if tenant and subscription are defined for azure

## Name: repo_config_validation

  1. Check repo folder
  2. Check existance of state files for each runner 
  3. Check definition of required values


## Name: runner_validation
1. Check specified runner definition; runner and runner versions are exists?
2. Dry run

## Name: compatilibility_validation
1. Check tool version(which is specified in repo manifest) and runner's tool version compability

## Name: pac_input_flow_validation
1. Check required repositories existance. Skip and print out a "warning" if an application repo does not exists. Abort if infrastructure repo does not exists.

## Name: input_output_validation
1. Check runners inputs and outputs compability.

## Name: state_validation
1. Check state files existance
2. Check state files for a corruption