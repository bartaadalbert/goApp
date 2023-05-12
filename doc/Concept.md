# AsciiArtify - Kubernetes Deployment Tools Analysis

## Introduction
AsciiArtify is a startup focused on developing a new software product that converts images into ASCII art using Machine Learning. The team consists of experienced software developers who are eager to launch their own product but lack DevOps expertise. In order to facilitate code storage and collaboration, AsciiArtify has chosen GitHub as their version control system. As they plan to use Kubernetes for deploying and scaling their software product, they have evaluated three Kubernetes deployment tools for local development: minikube, kind, and k3d.

## Tools Overview
### Minikube
Minikube is the most widely used local Kubernetes installer. It offers an easy to use guide on installing and running single Kubernetes environments across multiple operating systems. It deploys Kubernetes as a container, VM or bare-metal and implements a Docker API endpoint that helps it push container images faster. It has advanced features like load balancing, filesystem mounts, and FeatureGates, making it a favourite for running Kubernetes locally. However, concerns have arisen regarding its scalability limitations.

### Kind (Kubernetes IN Docker)
Kind is a tool that allows the creation of local Kubernetes clusters using Docker containers. AsciiArtify sees potential in using Kind for local testing and development or CI.

### K3d
K3d is a platform-agnostic, lightweight wrapper that runs K3s in a docker container using Rancher Kubernetes Engine (RKE). It helps run and scale single or multi-node K3S clusters quickly without further setup while maintaining a high availability mode. K3s helps you run a simple, secure and optimised Kubernetes environment on a local computer using a virtual machine such as VMWare or VirtualBox.

## Comparison
Here is a comparison of the three Kubernetes deployment tools based on various factors:

| Tool      | Supported OS & Architectures | Automation Capabilities | Additional Features        | Ease of Use | Deployment Speed | Stability | Documentation & Community Support | Configuration Complexity |
|-----------|-----------------------------|-------------------------|----------------------------|-------------|------------------|-----------|----------------------------------|--------------------------|
| Minikube  | Windows, macOS, Linux                   | Limited                 | Dashboard, Ingress, Addons | Easy        | Moderate         | Good      | Extensive                        | Moderate                 |
| Kind      | Windows, macOS, Linux                    | Moderate                | N/A                        | Easy        | Fast             | Good      | Moderate                         | Easy                     |
| K3d       |Windows, macOS, Linux                   | High                    | Load Balancer, Registry   | Easy        | Very Fast        | Good      | Good                             | Easy                     |

## Demo
To showcase the recommended tool, we have prepared a short demo that demonstrates the key capabilities of the chosen tool. The demo involves deploying a "GO APP" application on Kubernetes. You can find the detailed steps and code examples in the [demo](demo) folder of this repository.

## Conclusion
After evaluating the three Kubernetes deployment tools, we recommend the use of k3d for the PoC of the AsciiArtify project. It offers a good balance between ease of use, deployment speed, stability, and community support. Additionally, it provides valuable features like a load balancer and a container registry. While Minikube and Kind have their own advantages, k3d aligns better with the requirements and goals of AsciiArtify.

Demo: Testing and Deploying in Kubernetes with Minikube, Kind, and K3d

In this demo, we showcase the process of testing and deploying applications in Kubernetes using three popular tools: Minikube, Kind, and K3d. Each tool offers a unique approach to local Kubernetes cluster management. Please see the link https://asciinema.org/a/584385

For further information and a comprehensive comparison, please refer to the [Concept.md](doc/Concept.md) file in this repository.

## Next Steps
Docker and Podman are great container management engines and serve the same purpose in building, running, and managing containers. 
Containerization is a very efficient method of virtualization that is available out there. It makes it much easier for developers to test, build and deploy large-scale applications.Containers and Virtual machines are similar, but they’re not the same.
In Virtual Machines (VMs), the system’s resources are distributed between multiple virtual machines. The VMs run separately and can have different operating systems in them.
Containers, on the other hand, only virtualize software layers above the operating system level, which means you can have multiple virtualized environments with only one operating system. 

Here's the summary of the Podman vs Docker comparison in a table format:

|        | Docker                                   | Podman                                   |
|--------|------------------------------------------|------------------------------------------|
| Daemon | Uses a Daemon                             | Has a daemon-less architecture            |
| Privileges | Requires root privileges                | Can run root-less                         |
| Security | Less secure (containers have root privileges) | More secure (containers don't have root privileges) |
| Image Building | Docker builds images on its own       | Podman uses buildah to build images       |
| Swarm and Compose | Supports Docker swarm and Docker compose | Does not support Docker swarm or Docker compose |
| Independence | Monolithic independent tool            | Relies on other third-party tools         |

Please note that this table provides a brief overview of some key differences between Docker and Podman. There may be additional factors to consider based on your specific use case and requirements.

Based on the analysis and conclusions, AsciiArtify should proceed with implementing k3d for local Kubernetes cluster development. They should further explore the features and capabilities of k3d and leverage its strengths in the development and testing phases of their project.

Please refer to the [demo](demo) folder and documentation for detailed instructions on how to use k3d for local Kubernetes cluster setup and deployment.

Demo: Testing Docker Images with Dive

In this demo, we showcase the process of testing Docker images using Dive, a tool for exploring and analyzing Docker image layers.

    Installation: Start by installing Dive on your local machine. Here's a link https://asciinema.org/a/584389 to the asciinema recording that demonstrates the installation process: Dive Installation

    Analyzing Docker Images: Dive allows you to dive deep into the layers of Docker images and analyze their contents. Watch this asciinema recording to see how Dive helps in analyzing a sample Docker image: Analyzing Docker Images with Dive

    Identifying Optimization Opportunities: Dive helps in identifying optimization opportunities in Docker images, such as redundant or unnecessary layers. Check out this asciinema recording to understand how Dive can help optimize your Docker images: Optimizing Docker Images with Dive

By incorporating Dive into your Docker image testing workflow, you can gain insights into image composition, identify potential issues, and optimize your Docker images for better performance and security.

---

*Note: This document provides an overview and analysis of the Kubernetes deployment tools for the AsciiArtify startup. The specific implementation details and configurations may vary based on the requirements and constraints of the project.*