# kubesync
Kubernetes bi-directional sync utility

KubeSync is a tool which...
* Syncs from Kubernetes manifest files on disk to Kubernetes
* Syncs from Kubernetes to manifest file[s] on disk
* Allows you to deploy from file allowing you to confirm the diff between what is in use currently and what is being deployed
* Runs as a daemon and keeps both files and kube sync'd
* Has the option to interact with GitHub pulling and pushing changes
* Has the ability to pull all configuration files from Kubernetes and write them to disk
* Has the ability to sync all files recursively from disk to Kubernetes
  * Can detect order needed to restore files such that everything works
