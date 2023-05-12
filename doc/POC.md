You can find the demo instructions for deploying and configuring ArgoCD on Kubernetes in the AsciiArtify repository.

Link to AsciiArtify repository https://github.com/bartaadalbert/goApp

In the doc/POC.md file in Markdown format, located in the main branch, you will find detailed steps on how to access the ArgoCD interface.
 Asciiinema recording:
 [![Installation Process](https://asciinema.org/a/584476)](https://asciinema.org/a/584476)

Example demo from the official ArgoCD website  https://argo-cd.readthedocs.io/en/stable/

Please review this repository and documentation to get all the necessary information for deploying and configuring ArgoCD in the context of your AsciiArtify project.

If you have any further questions or need additional assistance, please let me know!

#create namespace argocd
#apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
#patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
#port-forward svc/argocd-server -n argocd 8080:443
#argocd admin initial-password -n argocd
PASSWORD IS :giEL2Gp3iD4GPUeP
LOGIN IS: admin