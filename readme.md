# upp-aggregate-healthcheck
The purpose of this service is to serve the functionality of aggregating healthchecks from services and pods in the Kubernetes cluster.

## Usage
 In this section, the aggregate-healthcheck functionalities are described.
### Get services health

### Get pods health for a service
 Pods health is evaluated by querying the health endpoint of apps inside the pods. Given a pod, if there is at least one check that fails, the pod health will be considered warning or critical, based on the severity level of the check that fails.
### Acknowledge a service
 When a service is unhealthy, there is a possibility to acknowledge the warning. By acknowledging all the services that are unhealthy, the general status of the aggregate-healthcheck will become healthy (it will also mention that there are 'n' services acknowledged).
### Sticky categories
 Categories can be sticky, meaning that if one of the services become unhealthy, the category will be disabled, meaning that it will be unhealthy, until manual re-enabling it. There is an endpoint for enabling a category.
## Running locally
 There is a limited number of functionalities that can be used locally, because we are querying all the apps, inside the pods and there is no current solution of accessing them outside of the cluster, without using port-forwarding.
 The list of functionalities that can be used outside of the cluster are:
  * Add/Remove acknowledge
  * Enable/Disable sticky categories
  
## Endpoints

### Service endpoints

### Admin endpoints
 * `__health` -
 * `__gtg`
    - params:
    - returns a __503 Service Unavailable__ status code in the following cases:
       - if at least one of the provided categories is disabled (see sticky functionality)
       - if at least one of the checked services is unhealthy
    - returns a __200 OK__ status code otherwise
