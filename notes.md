## 1) Dockerization

The phrase in the project pitch "Bonus points if only include what is needed to run the application." had me create a very simple dockerfile (it can perhaps be even simpler).

```bash
docker build -t uniwise .
```

```bash
docker run \
--env-file ./.env \
--mount type=bind,source="$(pwd)"/secrets/users.json,readonly,target=/app/secrets/users.json \
-p 4444:4444 devops-assignment
```

This creates a working and runnable image. .env and secrets/users.json needs to be loaded in manually due to .dockerignore.