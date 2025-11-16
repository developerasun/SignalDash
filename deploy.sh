docker build -t signaldash-backend:latest .
docker service rm signaldash_backend
docker stack up -c ./service.yaml signaldash
docker service ls