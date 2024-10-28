// s3bucket.go
package mod

import (
    "github.com/aws/aws-cdk-go/awscdk/v2/awss3"
    "github.com/aws/constructs-go/constructs/v10"
    "github.com/aws/jsii-runtime-go"
)

type S3BucketProps struct {
    Versioned bool
}

func NewS3Bucket(scope constructs.Construct, id string, props *S3BucketProps) awss3.Bucket {
    bucketProps := &awss3.BucketProps{
        Versioned: jsii.Bool(props.Versioned),
    }

    bucket := awss3.NewBucket(scope, jsii.String(id), bucketProps)
    return bucket
}