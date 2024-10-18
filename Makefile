# Makefile

SHELL := /bin/bash

.PHONY: init plan apply destroy clean

all: init plan

# Format Terraform code
format:
	@echo "Formating Terraform..."
	cd terraform && terraform fmt -check

check_resources:
	@echo "Validating resources config..."
	cd go \
		&& go get . \
		&& go mod tidy \
		&& go run -tags checker checker.go

# Create Terraform resources
update_resources: 
	@echo "Generating Terraform resources..."
	cd go \
		&& go get . \
		&& go mod tidy \
		&& go run -tags builder builder.go

# Initialize Terraform
init:
	@echo "Initializing Terraform..."
	@./init_backend.sh

# Plan Terraform changes
plan: init
	@echo "Planning Terraform changes..."
	cd terraform && terraform plan

# Apply Terraform changes
apply: init
	@echo "Applying Terraform changes..."
	cd terraform && terraform apply -auto-approve

# Destroy Terraform-managed infrastructure
destroy: init
	@echo "Destroying Terraform-managed infrastructure..."
	cd terraform && terraform destroy -auto-approve && rm users.tf

# Clean up temporary files
clean:
	@echo "Cleaning up..."
	cd terraform && rm -f config.s3.tmp

# Help target
help:
	@echo "Available targets:"
	@echo "  init    	 : Initialize Terraform with backend configuration"
	@echo "  plan    	 : Plan Terraform changes"
	@echo "  apply   	 : Apply Terraform changes"
	@echo "  destroy 	 : Destroy Terraform-managed infrastructure"
	@echo "  clean   	 : Clean up temporary files"
	@echo "  update_users: Create/Update Terraform resources"
	@echo "  help    	 : Show this help message"