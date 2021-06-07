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
/api/_healthcheck
/api/external
```

## Final thoughts

There are definetely multiple ways to tackle this problem. For this test I tried to use one tool  
for each part of the solution.

- Instead of using nginx as a rev proxy I could have served HTTPS directly from Golang and save  
one service.

- Instead of using an EC2 insance the Golang image could be deploy on a non managed service like  
AWS Fargate or even AWS Elastic Beanstalk

- I wouldn't create the images in the EC2 instance. Instead I would use a Jenkins pipeline to run  
the test, then build the image and publish it in a registry. Once done, I could use Ansible to  
pull the images in all the instances that required and deploy de containers.

Looking forward to discuss other approaches.