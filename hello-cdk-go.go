// main.go
package main

import (
    "github.com/aws/aws-cdk-go/awscdk/v2"
    "github.com/aws/jsii-runtime-go"
    "hello-cdk-go/mod"  // Importing the custom module
)

func main() {
    app := awscdk.NewApp(nil)

    stack := awscdk.NewStack(app, jsii.String("MyStack"), &awscdk.StackProps{})

    mod.NewS3Bucket(stack, "MyAwesomeBucket", &mod.S3BucketProps{
        Versioned: true,
    })

    app.Synth(nil)
}