# Koddi-framework-starter

Starter project for _framework_-based services. To begin a new service, it is recommended to use the GitHub "use this Template" button to create a new repo with a copy of the starter.

# Updating this project for your use

## Project Files

1) modify the go.mod file to change the name of your module: "module github.com/TravelMedia/_your-koddi-framework-name_"
2) modify the main.go imports to reflect you newly defined module (e.g. module github.com/TravelMedia/_your-koddi-framework-name_/config)
3) modify the .gitignore so that you ignore the GoLang Execuable (_your-koddi-framework-name_)
4) in `.circleci/kube/<type>.yaml` replace `koddi-framework-starter` with the appropriate name of your project in all places. You only need to replace this in the yamls that are relevant for this service (i.e. if its a REST service, `deployment.yaml` and `service.yaml` are useful, others are not.)

Be sure to choose a DNS-compliant name for your project (all lowercase, no punctuation except hyphens). We recommend
going as simple as possible when in doubt. Don't forget to change your project's module name to match its github.com url
in _go.mod_.

## Docker

The `Dockerfile` in the root directory should work for any project based on this platform. However, if you install new external dependencies (like C libraries, command line tools required, etc.) you will need to update the `Dockerfile` to include these libraries at runtime.

### Running Docker locally
Create a Personal Access Token for your Github Account. If you don't know how, Google it. You'll get a Password/token you should copy into your 1Password under your github entry. **Note: This is a temporary measure and will be replaced with a build user.**

```
%> docker build -t koddi/<projectname> --build-arg ACCESS_TOKEN_USR="<your_user>" --build-arg ACCESS_TOKEN_PWD="<access_token>" .`
...
...
# Example of running the rest service with container port 8080 mapped to your localhost port 8080
%> docker run -p 8080:8080 koddi/<projectname> 
```

You should now be able to access your service on `http://localhost:8080`. If it's broken, it won't work in production either, so you should investigate and fix.

## Builds and Deployments

`.circleci/config` describes the process to checkout, test, build, and release this project on CircleCI. Out of the box, this project supports the following build workflow:

* Docker build, push image, update two kubernetes specs (`deployment.yaml` and `service.yaml` by default)

The following Kubernetes Specs are supported:

* `deployment.yaml` - A kubernetes deployment open on port 8080.
* `service.yaml` - A kubernetes service that binds to the `deployment.yaml` pods, and creates a public ELB open on port 80.

Update the kubernetes specs and build configuration to your liking. If you need to support a new type of object or build workflow, implement it and make a PR to this repo.

### First build and deploy

If this is your first time there is a [video tutorial that runs through these steps](https://drive.google.com/open?id=1_tw36zaWxZCq6xW-0XVpbqwIKQidbRvq) that you can watch.

1) Go to AWS and create a new ECR repo with the project name that you want.
2) locally, do a `docker build -t <ecr_repo_url>/<ecr_repo_name>:latest .`
3) Upon a successful build, follow the instructions in ECR `Push Commands` to push your first build to the repo using the `latest` tag. This will make sure your build has something to pull down initially.
4) Go into CircleCI, login with Github, navigate in the top left to the Org where your repo lives.
5) Click `Add projects`. Find your project in the list and click `Set Up Project` and then `Start Building`
6) Click `Add manually` on the modal window and then click `Start Building`.
7) When on the main project page, click `Project Settings` and `Environment Variables` to setup a few key things:
* **AWS_ACCESS_KEY_ID** - The AWS Access Key for the user that will be used to build operations. (Your user for now)
* **AWS_SECRET_ACCESS_KEY** - The AWS Secret Access Key for the user that will be used to build operations. (Your user for now)
* **AWS_ACCOUNT_ID** - You can find the account id by clicking on My Account in the upper right of the AWS console.
* **AWS_ECR_ACCOUNT_URL** - The url part of the ECR repo. E.g. `815599370552.dkr.ecr.us-east-1.amazonaws.com`
* **AWS_REGION** - The region where ECR and the EKS cluster are located.
* **GH_ACCESS_TOKEN** - the Personal Access Token for the user that will clone the repo.
* **GH_USER** - The user that will clone the repo.
8) You'll be returned to the `Pipelines` screen where your first build will begin. If all is well you should see a build end in success. Checking it out will show that the deployment to the cluster was successful.
9) Use kubectl to check the objects were created and that your system is healthy. If its a REST service you can obtain the ELB url and hit the health endpoint. It will take several mins for it to spin up.

### To be implemented
* Kubernetes `Jobs`

## SSH configuration

If you don't already have an SSH key on your machine, run this command (substituting your actual koddi email address):

`ssh-keygen -t rsa -C "your.email@koddi.com"`

You can just hit enter three times to use the default key location and not use a passphrase.

Then, copy the output of `cat ~/.ssh/id_rsa.pub`. Go to your github.com settings -> SSH keys, and add it as a new key.

Finally, run the following command to configure Git (and Go) to authenticate via SSH.

`git config --global url."git@github.com:".insteadOf "https://github.com/"`


### Configure Git/Go to access private Github repository

go env -w GOPRIVATE=github.com/TravelMedia


This will use your SSH key.  See: 
https://golang.org/cmd/go/#hdr-Module_configuration_for_non_public_modules

## Updating dependencies

Any external packages imported by your Go code are tracked, and their versions pinned, in go.mod and go.sum.  Go uses "SemVer" version numbers in its dependency resolution, we will follow that recommendation and add a git tag in the
format "vX.X.X", where "X" is an integer representing major, minor, and patch versions respectively.

### Getting versioned modules

To get the specific version of a module, you need to add the version information (commit tag) to your `go get` command like this:

`go get github.com/TravelMedia/koddi-framework@v1.0.0`

The result of that would be your `go.mod` file looking like this:

```
module github.com.org//{yourModName}

go 1.14

require github.com/TravelMedia/koddi-framework v1.0.5
```

If your `go.mod` looks something like this:

> github.com/TravelMedia/koddi-framework v0.1

_then it is an indication that the tag was not named correctly or that your `go get` command was not done to the correct tag name_
