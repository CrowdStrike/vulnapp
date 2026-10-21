# CrowdStrike's VulnApp
[![Docker Repository on Quay](https://quay.io/repository/crowdstrike/vulnapp/status "Docker Repository on Quay")](https://quay.io/repository/crowdstrike/vulnapp)

## Usage - Generic Kubernetes

> [!IMPORTANT]
> This project intentionally incorporates security vulnerabilities and simulated attack scenarios to demonstrate CrowdStrike Falcon platform capabilities.
> These behaviors are expected by design, and the project is intended solely for demonstration and testing in controlled environments, not for production use.

```
kubectl apply -f  https://raw.githubusercontent.com/crowdstrike/vulnapp/main/vulnerable.example.yaml
```

Get web address:
```
watch -n 1 echo 'http://$(kubectl get service vulnerable-example-com  -o yaml -o=jsonpath="{.status.loadBalancer.ingress[0].ip}")/'
```

If the `ip` field is not present try:
```

watch -n 1 echo 'http://$(kubectl get service vulnerable-example-com  -o yaml -o=jsonpath="{.status.loadBalancer.ingress[0].hostname}")/'
```

Delete the app:
```
kubectl delete -f  https://raw.githubusercontent.com/crowdstrike/vulnapp/main/vulnerable.example.yaml
```

## Usage - OpenShift

The OpenShift-specific deployment uses a `Route` with automatic edge TLS termination, configures a ServiceAccount and RoleBinding to run the app as UID 0, and takes advantage of the Topology view's app grouping.

### Web console

1. Switch to the project you want to deploy the app to
1. Click the **(+)** icon in the top right
1. Copy and paste the contents of `vulnerable.openshift.yaml`
1. Click **Create**

To open the webpage, return to the Topology view click the URL link on the deployment icon.

To delete the app, click the kebab menu on the `vulnapp` application, then **Delete Application**. You'll have to manually delete the `vulnapp` ServiceAccount and `vulnapp-anyuid` RoleBinding.

### Command line

```
# Deploy the app
oc apply -f https://raw.githubusercontent.com/crowdstrike/vulnapp/main/vulnerable.openshift.yaml
# Get the URL
oc get route vulnapp
# Delete the app
oc delete -f https://raw.githubusercontent.com/crowdstrike/vulnapp/main/vulnerable.openshift.yaml
```
