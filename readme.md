# upp-aggregate-healthcheck
Aggregate healthcheck that is currently used for Kubernetes cluster

## Usage

## Running locally

## Endpoints

### Service endpoints

### Admin endpoints
 * `__health` -
 * `__gtg` 
    - returns a __503 Service Unavailable__ status code in the following cases:
       - if at least one of the provided categories is disabled (see sticky functionality)
       - if at least one of the checked services is unhealthy
    - returns a __200 OK__ status code otherwise
