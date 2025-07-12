# Atlas

## Introduction

Atlas is a project hosted on Amazon Web Services (AWS) that provides a URL shortening service. It is designed as a practical example of building and deploying a serverless application using various AWS services.

The infrastructure is defined and managed using the AWS Cloud Development Kit (CDK) with TypeScript. The backend services are implemented as AWS Lambda functions written in both Go and TypeScript.

> Atlas was the original codename for the project.

## Architecture Overview

The project implements a serverless architecture on AWS, leveraging several key services to provide a scalable and resilient backend system.

*   **Infrastructure as Code**: The entire cloud infrastructure is managed through the AWS CDK in TypeScript, allowing for version-controlled and repeatable deployments.

*   **Compute**: AWS Lambda is used for all compute operations. The functions are written in Go and TypeScript, demonstrating polyglot development within a single project.

*   **API Layer**: Amazon API Gateway provides the public-facing entry points for the service, including a REST API for URL shortening and a WebSocket API for real-time updates.

*   **Data Storage**: Amazon DynamoDB is used as the primary data store for persisting shortened URL data and WebSocket connection information.

*   **Messaging**: Amazon SQS is used for asynchronous processing, decoupling the initial URL submission from the backend processing logic.

*   **Code Architecture**: The Go-based Lambda functions are structured following the principles of **Clean Architecture**. This approach separates the core business logic (domain and use cases) from external concerns like database access or API gateways (infrastructure), resulting in a more maintainable and testable codebase. The project also utilizes **AWS Lambda Layers** to share common code and dependencies across multiple functions.

## Project Structure

The repository is organized into the following main directories:

*   `bin/`: Contains the entry point for the CDK application.
*   `lib/`: Defines the AWS resources and constructs that form the `AtlasInfraStack`.
*   `lambdas/`: Contains the source code for the Lambda functions, separated by language (`go` and `js`).
*   `layers/`: Contains the source code for the shared Lambda Layer.
*   `build-aux/`: Includes build scripts for the project.
*   `assets/`: Stores static assets, such as images for documentation.

## Getting Started

Follow these instructions to set up and deploy the project in your own AWS account.

### Prerequisites

Before you begin, ensure you have the following installed:

*   [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
*   [Node.js and npm](https://nodejs.org/en/download/)
*   [Go Compiler](https://go.dev/doc/install)

### Setup and Deployment

1.  **Configure AWS CLI**

    First, configure the AWS CLI with your credentials. This allows the CDK to deploy resources to your account.

    ```shell
    aws configure
    ```

2.  **Clone and Install Dependencies**

    Clone the repository and install the root Node.js dependencies.

    ```shell
    git clone https://github.com/AlphaTechnolog/atlas-infra.git ~/work/atlas-infra
    cd ~/work/atlas-infra
    npm install
    ```

3.  **Bootstrap CDK Environment**

    Deploy the `CdkToolkit` stack to your AWS account. This stack contains resources required for the CDK to operate.

    ```shell
    npm run cdk bootstrap
    ```
    > For more information on bootstrapping, refer to the [official AWS CDK documentation](https://docs.aws.amazon.com/cdk/v2/guide/bootstrapping.html).

    After a successful bootstrap, you will see a `CDKToolkit` stack in your AWS CloudFormation console.

    ![CloudFormation Bootstrap](./assets/cloudformation-bootstrap.png)

4.  **Install Service Dependencies**

    Install the dependencies for the individual JavaScript-based Lambda functions and layers.

    ```shell
    # Install dependencies for JS Lambdas
    cd lambdas/js/
    for x in *; do cd $x && npm install && cd -; done
    cd ../../

    # Install dependencies for the Lambda Layer
    cd layers/atlas-url-shortener-layer && npm install && cd -
    ```

5.  **Build All Project Artifacts**

    Compile the TypeScript code and build the Go binaries for all services.

    ```shell
    npm run build:all
    ```

6.  **Deploy the Application Stack**

    Finally, deploy the `AtlasInfraStack`, which contains all the application's resources.

    ```shell
    npm run cdk deploy
    ```

    This command will provision all the necessary AWS resources for the URL shortener backend.

## Frontend Application

A frontend application designed to interact with this backend is available in a separate repository.

*   **Project Repository**: [atlas-frontend](https://github.com/AlphaTechnolog/atlas-frontend)