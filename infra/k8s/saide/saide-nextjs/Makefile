DEPLOYMENT_NAME = saide-nextjs-deployment
SERVICE_NAME = saide-nestjs-service

DEPLOYMENT_FILE = $(DEPLOYMENT_NAME).yaml
SERVICE_FILE = $(SERVICE_NAME).yaml

KUBE_APPLY = kubectl apply -f
KUBE_DELETE = kubectl delete -f

all: up

up: saide-nextjs-deployment.yml saide-nextjs-service.yml
	@echo "Deploying saide-nextjs to Kubernetes..."
	$(KUBE_APPLY) $(DEPLOYMENT_FILE)
	$(KUBE_APPLY) $(SERVICE_FILE)
	@echo "Deployment complete."

down: saide-nextjs-deployment.yml saide-nextjs-service.yml
	@echo "Deleting saide-nextjs from Kubernetes..."
	$(KUBE_DELETE) $(DEPLOYMENT_FILE)
	$(KUBE_DELETE) $(SERVICE_FILE)
	@echo "Deletion complete."

status:
	@echo "Fetching deployment status..."
	kubectl get deployments $(DEPLOYMENT_NAME)
	@echo "Fetching service status..."
	kubectl get svc $(SERVICE_NAME)

.PHONY: all up down status
