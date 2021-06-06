# Application for Devops Engineer at Holded

## 1. Infrastructure Test

For this point I tried to keep it simple. I provisioned a EC2 instance and set the docker  
installation commands in the "user data". At this point I have a server with all the requiered  
tools to complete the next steps.

Note:  
I realize that the steps from provisioning the instance all the way to deploy the application  
could be done on a single run by setting a provisioner remote and local exec. This way  
terraform is able to interact with ansible and run the playbook. I'll do this in a separete branch.


## 2.1 Application (CI/CD)

In this point I'm using Ansible to deploy all the solution inside the remote machine.  
As I mentioned before this can be done in one run, but I'm showing how terraform appends  
the IP from the new instances into the ansible inventory so I can run it separately.


## 2.2 Application (Coding)

