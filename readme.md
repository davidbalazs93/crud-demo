# upp-aggregate-healthcheck
The purpose of this service is to serve the functionality of aggregating healthchecks from services and pods in the Kubernetes cluster.

## Usage
 In this section, the aggregate-healthcheck functionalities are described.
### Get services health
### Get services health for certain categories
### Get pods health for a service
### Acknowledge a service
### Sticky categories

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
