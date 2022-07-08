# aws_lambda_demo
Authentication service demo using lambda function, aws api gateway.

### Install sam cli.
1. Install by home brew
``` bash
brew tap aws/tap
brew install aws-sam-cli
```

### Workflow
![](image.png)

### How the JWT Authorizer work.
https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-jwt-authorizer.html


### Lambda authorizer of api can't work in local.
We only test api gateway authorizer in aws console.
ISSUE: https://github.com/aws/aws-sam-cli/issues/137
