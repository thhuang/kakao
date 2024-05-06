KUBE_APPLY = kubectl apply -f
KUBE_DELETE = kubectl delete -f

all: deploy

deploy: y-ws-deployment.yml y-ws-service.yml
	@echo "Deploying y-ws to Kubernetes..."
	$(KUBE_APPLY) y-ws-deployment.yml
	$(KUBE_APPLY) y-ws-service.yml
	@echo "Deployment complete."

delete: y-ws-deployment.yml y-ws-service.yml
	@echo "Deleting y-ws from Kubernetes..."
	$(KUBE_DELETE) y-ws-deployment.yml
	$(KUBE_DELETE) y-ws-service.yml
	@echo "Deletion complete."

status:
	@echo "Fetching service status..."
	kubectl get svc y-ws-service
	@echo "Fetching deployment status..."
	kubectl get deployments y-ws-deployment

.PHONY: all deploy delete status
