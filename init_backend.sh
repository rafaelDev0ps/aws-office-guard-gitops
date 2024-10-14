#!/bin/bash

cd terraform

# Create a temporary file with substituted values
envsubst < config.s3.hcl > config.s3.tmp

# Initialize Terraform with the temporary config file
terraform init -backend-config=config.s3.tmp

# Remove the temporary file
rm config.s3.tmp