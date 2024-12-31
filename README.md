# Skip Platform Take Home Challenge

The goal of this excercise was to integrate on to this repository some quality of life improvements for the developers working on it.

Below I'll list the main problems assessed and a short description of the solution applied. 

* the deployment process is manual and error-prone. Whenever a deployment happens, we suffer a tiny bit of downtime due to the server being down.
    - The deployment has been automated whenever a tag is created on the main branch, a github workflow is triggered that deploys it to a kubernetes cluster.
* there's no standardization for code formatting which leads to inconsistencies
    - golangci-yaml has been added to allow for a more fine grained configuration of the code formatting and it runs as a pre-commit hook. A very simple config has been added just to kickstart its use, configuring the lint format was out of the scope of the exercise.
* the proto-gen script is ran locally which leads to developers forgetting to do it before pushing code upstream
    - The script now runs as a pre-commit hook.
* downloading the tooling dependencies is a manual, undocumented process
    - Devcontainer has been setup to allow developers to have a consistent environment with all tools needed already installed.
* running the server locally is a bit of a pain, since it requires manually standing up a Postgres database (to replicate prod)
    - A docker-compose file has been setup to allow devs to run locally with 1 command the DB + Rest API with live reloading.
* there's no easy way to share a feature update with our colleagues, since we only have a local and production environment.
    - A github workflow when opening a PR has been added that deploys the feature in a kubernetes virtual cluster inside the real kubernetes cluster, and when the PR is closed or merged it destroys the virtual cluster.


## Expected development flow

### Requisites

- [Docker](https://docs.docker.com/desktop/) should be installed.
- VSCode [devcontainer](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension should be installed.

### Setup

Run the VSCode command (Press F1 key to open the VSCode command palette) `Dev Containers: Reopen in container`, this will open the repository on the Docker img specified inside the `.devcontainer` folder on the root folder of the repo.

Open an integrated terminal, and you will have a bash shell inside the devcontainer, with all the required tools installed to work on this repository.

From here you can code as you would normally.

### Nice things to know

There are a couple pre-commit hooks that will execute after trying to commit code. Which will execute the golangci-lint and the proto-gen script.

You can find the hooks at the file `lefthook.yml`. You can also find and add your desired golangci-lint config at `.golangci.yml`.


## Running locally

### Requisites

- [Docker](https://docs.docker.com/desktop/) should be installed.

### Running the API + DB

Copy `.env.template` into an `.env` file in the root of the repository.

From a local terminal (not the one inside the devcontainer) on the root of the repository run:

`docker-compose up -d`

This will start 2 docker containers one for the DB and other for the Rest API, you can access the db at `localhost:5432` and the api at `localhost:8080` and `localhost:8081`.

Any change you make on the src code will be reflected on the instance locally running and it will build the binary and restart the running instance with the new binary.

To stop the containers run `docker-compose down`.

## Deployment

### ⚠️ DISCLAIMER ⚠️

**Please Read Before Proceeding**

To deploy the services on a real Kubernetes cluster, I have included **simple Terraform scripts** to provision a GKE cluster. These scripts:

- **Should not** typically be part of this repository, but since this is a test, I decided to include them.
- Are **not running constantly** due to monetary constraints.

### How to Test the Deployment Flow
If you would like to test the deployment flow, please **contact me**. The cluster takes approximately **15–20 minutes** to be fully operational.

### What Happens If the Cluster Is Down?
- Any **Google Cloud operations** in the GitHub workflows will be **skipped**.
- The workflow will continue executing other parts of the pipeline.

> 🛑 **Note**: Once I receive your message, I can quickly bring the cluster back up for testing.

### Deploying to ephemeral dev environment

When a PR is opened a github workflow will trigger that builds and push the docker image, and then creates a virtual cluster on top of the real GKE cluster, where it deploys the new changes and provides a kubeconfig file as an uploaded artifact on the workflow for the developer to be able to connect to the cluster.

When the PR is either merged or closed the ephemeral virtual cluster will be destroyed.

#### Accessing the Deployed Ephemeral Application

To access the ephemeral application deployed by the pull request workflow:

1. **Download the kubeconfig**:
   - In the pull request Github action, under the **pr-deploy** job, locate the **Artifacts** section and download the `kubeconfig` file.

2. **Set Up kubeconfig**:
   - Set up the `kubeconfig` inside the **devcontainer** to avoid modifying your local environment.


3. **Retrieve the Service External IP**:
   - Run the following command to get the external IP address of the service:
     ```bash
     kubectl get service takehome-api -n default -o jsonpath='{.status.loadBalancer.ingress[0].ip}'
     ```

4. **Test the Application**:
   - Use the external IP to test the application.

### Notes:
- Using the `devcontainer` ensures a clean setup and avoids conflicts with your local Kubernetes configuration.

### Deploying to production

Assumptions: 
- I'm using the same GKE cluster as pro/ephemeral-dev it's not ideal but it's cost effective.
- The GKE cluster is by no means production ready it lacks among other things any certificate manager or DNS, that's why the API is exposed currently on an external IP.

When a tag with the format `v*.*.*` is created on the main branch a github workflow is triggered, which will build and push the docker image and then deploy it on GKE cluster.

#### Accessing the Deployed Application

Since configuring a domain for the application has been omitted, you can access the deployed application by following these steps:

1. Navigate to the **GitHub Action** named **Deploy to GKE**.
2. Under the **deploy** job, locate the step titled **Deployment Info**.
3. In this step, the **EXTERNAL_IP** of the application will be logged, along with example `curl` commands to test the deployed application.


## Next Steps

- Improving devcontainer: Docker-compose should be runnable inside the devcontainer environment.
- GKE cluster is not production ready: adding certificate-manager, nginx-ingress/traefik and a handful of other needed services. That will also allow us to improve the helm chart that deploys the app.
- Improve security: static code analysis, docker image scanning, gosec, implementing a configuration server to better manage secrets, distroless container, etc.
- Improving CI/CD: Adding cache, using official github actions instead of so many scripts.
- If the lists of commands to perform actions grows, maybe add a makefile to centralize them.
- Talk with the devs: 
    - Why there is no testing? Try to figure out if they have any blocker to add tests and help them overcome it.
    - Same for documentation.



