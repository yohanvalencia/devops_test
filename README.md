# Application for Devops Engineer at Holded

## 1. Infrastructure Test

For this point I tried to keep it simple. I provisioned a EC2 instance and set the docker  
installation commands in the "user data". At this point I have a server with all the requiered  
tools to complete the next steps.

```
terraform apply -auto-approve
```

## 2.1 Application (CI/CD)

I'll use Ansible to deploy all the solution inside the remote machine. The IP from  
the new instances is appended into the ansible inventory.

```
ansible-playbook holded_test.yaml
```

## 2.2 Application (Coding)

I created two endpoints with the following path:

```
/_healthcheck
/external
```