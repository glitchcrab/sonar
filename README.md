# Sonar

Sonar deploys a debugging container to a Kubernetes cluster.

## Configuration

### Global flags

```
--kube-config (default: '/home/$user/.kube/config')

Absolute path to the kubeconfig file to use.

--kube-context (default: current context in kube config)

Name of the kube context to use to use. NOTE: not yet implemented.

--name (default: 'debug')

Name given to all the created resources. This will be automatically
prepended with 'sonar-', so a provided name of 'test' will result
in a deployment named 'sonar-debug'. Provided name can be a max of
50 characters.

--namespace (default: 'default')

Namespace to deploy resources to.
```

### Create

#### Options

```
--image (default: 'busybox:latest')

Name of the image to use. Image names may be provided with or without a
tag; if no tag is detected then 'latest' is automatically used.

--pod-cmd (default: 'sleep')

Command to use as the entrypoint.

--pod-args (default: '24h')

Args to pass to the command.

--pod-userid (default: 1000)

User ID to run the container as (set in deployment's SecurityContext).

--podsecuritypolicy (default: false)

Create a PodSecurityPolicy and the associated ClusterRole and Binding.
The PSP will inherit the value set via --pod-userid and configure the
minimum value of the RunAs range accordingly.

--privileged (default: false)

Allow the pod to run as a privileged pod; must be provided at the same
time as --podsecuritypolicy to have any effect.

--networkpolicy (default: false)

Apply a NetworkPolicy which allows all ingress and egress traffic.
```

#### Examples

```
sonar create
```
 - accept all defaults. Creates a deployment in namespace `default` called `sonar-debug`.  The pod image will be `busybox:latest` with `sleep 24h` as the initial command.

```
sonar create --image glitchcrab/ubuntu-debug:v1.0 --pod-cmd sleep \
    --pod-args 1h
```
 - uses the provided image, command and args.

```
sonar create --podsecuritypolicy --pod-userid 0 --privileged
```

- creates a deployment which runs as root. Also creates a PodSecurityPolicy (and associated RBAC) which allows the pod to run as root/privileged.

```
sonar create --networkpolicy
```

 - also creates a NetworkPolicy which allows all ingress and traffic to the Sonar pod.

### Delete

#### Options

```
--force (default: false)

Skips all interaction and deletes all resources created by Sonar.
```

#### Examples

```
sonar delete
```
 - deletes all resources which match the defaults. This will result in deleting all resources in namespace 'default' which are named 'sonar-debug'.

```
sonar delete --name test --namespace kube-system
```
 - deletes all resources in namespace 'kube-system' named 'sonar-test'.
