# kubesync
Kubernetes bi-directional sync utility providing the following
* Convenience around selecting contexts and namespaces
* Extra layers of protection against pushing the wrong thing to an environment via diffs and confirmation
* Easy deployment accepting updated image (or other value) from command line updating local resource and pushing to remote
* Selective or full capture of resources from remote to local files
* Selective or full update or create of remote resources from local
* Promotion of updated settings from one context/namespace to another (i.e. staging to production)
* A daemon service which monitors local and remote keeping one in sync with the other or both

# Example Usage
I want to create a deployment, test it out in staging, then move it to production. It is a new one.
```
$ vi myDeploy.yaml
(create deployment on disk)
$ kubesync env nonprod staging               # You have set your context to nonprod and namespace to staging
$ kubesync push deployment myDeploy[.yaml]   # Now your deployment is in staging
$ echo "I've tested my deployment in staging and am ready to go to prod"
$ kubesync promote deployment myDeploy nonprod/staging prod/prod
$ echo "yay, I'm done and didn't break things!"
$ echo "but wait, someone found a bug!"
$ kubesync revert deployment myDeployment prod/prod
```

# Usage:
## Sub Commands
* configure                  - Configure server settings and defaults writing them to file (can set where promote goes from and to?)
* env                        - Prompt the user for which context
* pull                       - Pull a resource form Kubernetes to local file
* push                       - Push a resource from local file to Kubernetes
* deploy                     - Accept updated param via command line updating local resource as well as remote
* promote                    - Promote a build from one environment/namespace to another (default is deployment?)
* revert                     - Fall back to the last thing that was deployed
* daemon                     - Run as a daemon (logs to to std out, potentially give -b option for built in background operation)

## Global Options
* -f --force - Do not ask to confirm or compare diffs
* -o --output - Name of file to be written locally (should use default extension if none given)
* -c --context - Kubernetes context
* -n --namespace - Kubernetes namespace to interact with (uses currently configured options if not specified)

## Command Details
### CONFIGURE

### ENV
#### Sub Commands
  * (empty) - Returns current config overview
  * get
  * set

This command should potentially also help manage environment settings when used outside of kubesync (i.e. you run kubesync env [context] [namespace] then run kubectl get pods, you would be using the context and namespace set by kubesync)

#### Params
  * context
  * namespace

#### Examples
* env context [context name] - Accept a new setting for context (non interactive)
* env namespace [namespace]  - Accept a new setting for namespace (non interactive)

### PULL
$ kubesync pull [resource type] [resource name] [ -o myresource.yaml ]

$ kubesync pull
$ kubesync pull deployments
$ kubesync pull services

### PUSH
Update remote with local version
```
$ kubesync push [ deployment.yaml | deployment | MyDeploymentName ]
```
* Straight forward in the case of deployment.yaml
* For `deployment` it will guess at the extension and fail if there are more than one file with the same name but different extensions
* For `MyDeploymentName` it should load all of the resources in the current directory and determine which one has this name

Push everything recurisvely starting in $PWD
```
$ kubesync push
```

Push all deployment type resources recursively
```
$ kubesync push deployments
```

### DEPLOY

### PROMOTE
Feature Ideas
* Contextually determine what resource you're wanting to promote by $CWD
* Configure the from and to through kubesync configure

### REVERT

### DAEMON

# Configuration
* Config File (default ~/.kubesync/config)
* File format (default yaml)
* Kubernetes config (default ~/.kube/config)
* Root of local config path (default $PWD) (unless can be detected by familiar directory structure)

# Naming & Directory Structure
* File Naming
* Directory Structure

# Roadmap
## V1
- [ ] Pull a resource from Kube and save to file
  - [ ] Sanatize file removing unneeded elements
  - [ ] Sort (optional) for consistent/predictable files
  - [ ] If no file exists
    - [ ] Save it using default format
  - [ ] If file exists
    - [ ] Load the existing file
    - [ ] Sanatize both files (file on disk should ideally already be sanatized)
    - [ ] Sort both files
    - [ ] Diff them
    - [ ] Present diffs to user and wait for user to choose option
      - [ ] Update local file
      - [ ] Bail
      - [ ] Selectively update local file
- [ ] Push a resource to Kube
  - [ ] Sanatize file removing unneeded elements
  - [ ] Sort resource
  - [ ] Pull remote version
  - [ ] If it doesn't exist
    - [ ] Just push it
  - [ ] If it exists
    - [ ] Load the remote resource
    - [ ] Sanatize remote resource
    - [ ] Sort remote resource
    - [ ] Diff the two
    - [ ] Present diffs to user and wait for user to choose option
      - [ ] Update remote resource
      - [ ] Bail
      - [ ] Selectively update remote resource
      - [ ] Aloow to sync remote to local for select elements?

## V2
- [ ] Add daemon functionality
- [ ] Enable git syncing up/down

# Notes

## Feature Brain Dump
KubeSync is a tool which...
* Syncs from Kubernetes manifest files on disk to Kubernetes
* Syncs from Kubernetes to manifest file[s] on disk
* Allows you to deploy from file allowing you to confirm the diff between what is in use currently and what is being deployed
* Runs as a daemon and keeps both files and kube sync'd
* Has the option to interact with GitHub pulling and pushing changes
* Has the ability to pull all configuration files from Kubernetes and write them to disk
* Has the ability to sync all files recursively from disk to Kubernetes
  * Can detect order needed to restore files such that everything works
* When there are changes on both sides
  * Present options for which changes to keep
  * When things exist that can't be resolved, ability to alert or notify
* Manage kubernetes context and namespace (potentially allow selection)
  * Confirm (with extra confirmation to select contexts and namespaces)
* Keep track of the last thing that changed for each resource and allow easy quick revert
* Allow promoting an image or settings from one cluster or context to another

## Sub Command Thoughts
Separate what we're doing by what we're chainging in the resource?
* Deploy - Would only change the image
* Env - Would only change environment variables
* Configure - Would change other configuration?
* Scale - Would scale the deployment up or down

## TODO
* Make viper read env variables to override config values
  * https://scene-si.org/2017/04/20/managing-configuration-with-viper/
