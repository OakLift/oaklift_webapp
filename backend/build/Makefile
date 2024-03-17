
.PHONY: all
all: user-guard nexusdb

.PHONY: user-guard
USER_GUARD_DOCKERFILE=../user-guard/build/user-guard.Dockerfile
user-guard:
	DOCKER_BUILDKIT=1 docker build \
		--file=${USER_GUARD_DOCKERFILE} \
		--tag=prod/user-guard:latest \
		..


.PHONY: nexusdb
NEXUS_DB_DOCKERFILE="../nexusdb/build/nexusdb.Dockerfile"	
nexusdb:
	DOCKER_BUILDKIT=1 docker build \
		--file=${NEXUS_DB_DOCKERFILE} \
		--tag=prod/nexusdb:latest \
		..
