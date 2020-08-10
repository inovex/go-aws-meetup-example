# Go AWS Meetup Example
Minimalistic example service to be deployed as a lambda function on AWS.
It's intended to show the best practices we explained during our meetup "Building an MVP with serverless Go".

## How to use this
This is just a simple demo service that reads and writes some items to DynamoDB.
It is only used for education purposes and doesn't provide any real value apart from that.
If you want to use this as a way to get started with Go and AWS, feel free to dig through the source and learn from our experience.

## Build and Run
To build and run the software locally, you can use make.
```bash
make build
```
This will generate an executable file called `main` in the project directory.
You can put `.aws.access` and `.aws.secret` files in the working directory, containing the access key and secret access key of a set of user credentials, respectively.
The example service will use those to authenticate with AWS when you run the service locally.

To get more information, run
```bash
./main -h
```
The service assumes AWS resources (specifically DynamoDB tables) to be present and accessible via the provided credentials.

## Deploy
Deployments are handled by terraform.
To get started, look in the `terraform/` directory.
There is a terraform wrapper script called `terraformw.sh` which simplifies some things during initialization and for everyday use.
Mainly, it will set the stage variable to the first positional argument given:
```bash
./terraformw.sh dev apply
```
will effectively run
```bash
terraform -var="stage=dev" apply
```
To get started in your own AWS account, you need:

  - a set of credentials with administrator access to your account (create/modify/delete resources) configured to use with the AWS cli...
  - S3 buckets called `go-meetup-tf-state-<stage>` in the `eu-central-1` region...
  - DynamoDB tables called `tf-state-lock-<stage>` in the `eu-central-1` region...

where `<stage>` is equal to `dev`, `int` or `prd`.
If you want to deploy in another region, make some changes in `main.tf` and `main.go`.
