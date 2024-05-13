DEPLOYMENT_NAME = y-ws-deployment
SERVICE_NAME = y-ws-service
INGRESS_NAME = y-ws-ingress

DEPLOYMENT_FILE = $(DEPLOYMENT_NAME).yml
SERVICE_FILE = $(SERVICE_NAME).yml
INGRESS_FILE = $(INGRESS_NAME).yml

KUBE_APPLY = kubectl apply -f
KUBE_DELETE = kubectl delete -f

all: up

up: y-ws-deployment.yml y-ws-service.yml y-ws-ingress.yml
	@echo "Deploying y-ws to Kubernetes..."
	$(KUBE_APPLY) $(DEPLOYMENT_FILE)
	$(KUBE_APPLY) $(SERVICE_FILE)
	$(KUBE_APPLY) $(INGRESS_FILE)
	@echo "Deployment complete."

down: y-ws-deployment.yml y-ws-service.yml y-ws-ingress.yml
	@echo "Deleting y-ws from Kubernetes..."
	$(KUBE_DELETE) $(DEPLOYMENT_FILE)
	$(KUBE_DELETE) $(SERVICE_FILE)
	$(KUBE_DELETE) $(INGRESS_FILE)
	@echo "Deletion complete."

status:
	@echo "Fetching deployment status..."
	kubectl get deployments $(DEPLOYMENT_NAME)
	@echo "Fetching service status..."
	kubectl get svc $(SERVICE_NAME)
	@echo "Fetching ingress status..."
	kubectl get ingress $(INGRESS_NAME)

.PHONY: all up down status
