# DevOps Assignment

## About the Assignment
The assignment consists of multiple steps where each step builds upon the previous. So to avoid having to redo something it is recommended to read the whole assignment before diving in. 

The main objective of the assignment is to evaluate your skills and expertise in regards to the DevOps field. It is by no means intended to be exhaustive and is intended as a speaking framework rather than a piece of production code. While the assignment is not designed for production we do expect quality and thought in the delivered solution.

If anything seems a bit vague it is probably intentional to allow you to make your own assumptions.

## Desired outcomes
The ideal outcome is a solution that can initiate a conversation. It would be great (but not mandatory) if you alongside your solution included a `notes.md` file that described your thought process and choices. This can also be done with comments directly in the solution.

## Assignment
In this repository you will find a small REST server written in go. Your job is to take it from source code to a kubernetes deployment.

Run kubernetes locally using minikube (or in a cluster if you have access to one)

### 1) Dockerization
Create a Dockerfile with your base image of choice. Bonus points if only include what is needed to run the application.

Build the image locally and verify that the server works.

### 2) CICD
Create a CI pipeline with your tool of choice. It should run the provided tests, build the image and push it to a registry (could be https://ttl.sh).

### 3) Kubernetes
Create the necessary kubernetes resources in order to deploy the service to your kubernetes cluster.

### 4) Resilience
Make sure that data is not lost in the case of a pod getting rolled.

### 5) Expose out of cluster
Expose the service outside of the cluster in a way that would be similar to what you would do in a production cluster.

## Feedback
It would be great for future hires if you are willing to provide feedback on the assignment as a whole.


## Caveats
This assignment is relatively new so there might be some strange things please ask if something seems very strange.

The assignment has only been tested on Linux. There is not reason why it should not be completable on other operative systems but speak up if something seems impossible.
