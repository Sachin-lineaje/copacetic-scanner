# 🏭 Copacetic Scanner Plugin Template

## Development Pre-requisites

The following tools are required to build and run this template:

- `git`: for cloning this repo
- `Go`: for building the plugin
- `make`: for the Makefile

## Example Development Workflow

This is an example development workflow for this template.

```shell
# clone this repo
git clone https://sachin097@bitbucket.org/veedna/copacetic-scanner.git

# change directory to the repo
cd copacetic-scanner

# build the copa-lineaje-scanner binary
make

# add copa-lineaje-scanner binary to PATH (path might change depending on linux/mac) 
# Note: add absoulate path because Go throw's error for relative paths
export PATH=$PATH:absolute-path-to-dist/darwin_arm64/release/

# test plugin with example config
copa-lineaje-scanner testdata/fake_fix_plan.json
# this will print the report in JSON format
# {"apiVersion":"v1alpha1","metadata":{"os":{"type":"FakeOS","version":"42"},"config":{"arch":"amd64"}},"updates":[{"name":"foo","installedVersion":"1.0.0","fixedVersion":"1.0.1","vulnerabilityID":"VULN001"},{"name":"bar","installedVersion":"2.0.0","fixedVersion":"2.0.1","vulnerabilityID":"VULN002"}]}

# run copa with the scanner plugin (copa-lineaje-scanner) and the fix plan file
copa-lineaje patch -i $IMAGE -r testdata/fake_fix_plan.json --scanner lineaje-scanner
# this is for illustration purposes only
# it will fail with "Error: unsupported osType FakeOS specified"
```