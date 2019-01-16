#libraries
```
go get github.com/aws/aws-lambda-go/lambda
go get  github.com/aws/aws-sdk-go
```
# compile
```
sh compilte.sh
```

# using CLI to upload
```
AWS_PROFILE=personal1 aws lambda create-function --function-name questionresponse --runtime go1.x --role arn:aws:iam::<Account_ID>:role/service-role/lambda --handler main --zip-file fileb://main.zip
```

# Invoking function
```
AWS_PROFILE=personal1 aws lambda invoke --function-name questionresponse --payload '{"Question":"How tall are you?"}' outfile
```

# Update Lambda
```
AWS_PROFILE=personal1 aws lambda  update-function-code --function-name questionresponse --zip-file fileb://main.zip
```