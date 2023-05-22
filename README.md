# YAML Kubernetes Manifests with Prompt Examples

This repository contains a collection of YAML Kubernetes manifests with prompts. These prompts can be used with the `kubectl-ai` plugin to generate the corresponding manifest files. Below is a list of the available examples:

| NAME                   | PROMPT                  | DESCRIPTION                          | EXAMPLE                                               |
|------------------------|-------------------------|--------------------------------------|-------------------------------------------------------|
| app.yaml               | Create deployment       | Deploy the application to Kubernetes  | [Example](./yaml/app.yaml)                          |
| app-livenessProbe.yaml | Add liveness probe      | Configure liveness probe for the app | [Example](./yaml/app-livenessProbe.yaml)            |
| app-readinessProbe.yaml| Add readiness probe     | Configure readiness probe for the app| [Example](./yaml/app-readinessProbe.yaml)           |
| app-volumeMounts.yaml  | Add volume mounts       | Mount volumes in the application     | [Example](./yaml/app-volumeMounts.yaml)             |
| app-cronjob.yaml       | Create cron job         | Schedule a cron job in Kubernetes    | [Example](./yaml/app-cronjob.yaml)                  |
| app-job.yaml     | Create job   | Run a job         | [Example](./yaml/app-job.yaml)                |
| app-multicontainer.yaml| Create multi-container  | Define a pod with multiple containers| [Example](./yaml/app-multicontainer.yaml)           |
| app-resources.yaml     | Specify resource limits | Set resource limits for the app       | [Example](./yaml/app-resources.yaml)                |
| app-secret-env.yaml    | Use secrets as env vars | Inject secrets as environment vars   | [Example](./yaml/app-secret-env.yaml)               |

## Usage

To use these prompts, follow these steps:

1. Install the `kubectl-ai` plugin by downloading the binary for your operating system from the [kubectl-ai GitHub releases](https://github.com/sozercan/kubectl-ai/releases) page.

2. Make the `kubectl-ai` binary executable and move it to a directory in your system's `PATH`. sudo mv kubectl-ai /usr/local/bin/

3. Navigate to the directory where you want to generate the YAML manifest.

4. Run the `kubectl ai` command followed by the desired prompt. For example, to create a deployment, you can use the following command:

5. Follow the prompts and provide the required information.

6. The corresponding YAML manifest file will be generated in the current directory based on your inputs.

Feel free to explore the examples and generate the desired Kubernetes manifests using the prompts.


