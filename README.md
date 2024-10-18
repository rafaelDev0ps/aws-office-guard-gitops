# aws-office-guard-gitops
This is a project focused to integrate AWS IAM features in GitOps processes. Btw, this is a project for learning purposes, there's no LLM usage, just me and the effort of my little brain.

### Creating a user in AWS
Open a PR following the branch pattern `user/<username>`. Modify the `config_users.yaml` with all the information required. Check the example
```yaml
- name: robert
  description: robert is a great guy
  tags:
  - name: robert
  - project: orion
  permissions:
  - effect: Allow
    actions:
    - "ec2:Describe*"
    - "iam:GetUsers"
    resources:
    - "*"
```  
You can also pass the resource ARN in the resources list.

### Hey, what happens under the hood?
1. Golang `builder.go` script will generate a Terraform file based on a configuration file explained before
2. The information will fill out the template (like Python Jinja)
3. Terraform will read and apply this template 
  
### Is that? There's nothing else?
- There's also a config file validator in `checker.go` script. 
- This project is integrated with Snyk code vuln. scanner.
  
Other features will came up in the future, I'm tired now, I need rest.