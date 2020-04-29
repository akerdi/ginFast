#!/bin/sh

VERSION=1.0.0
COMMITS=`git log| grep ^commit| wc -l | xargs`

RUN_IN() {
  CWD="`pwd`"
  cd "$2"
  eval $1 || exit 1
  cd "$CWD"
}

PRINT() {
  echo "\n$1"
}

PRINT "start pack backend\n"
RUN_IN "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o staticBuilds/ginFast_linux" "."
PRINT "complete pack backend"

PRINT "start build & copy frontend to public\n"
RUN_IN "npm run build && cp -r dist ../public" "frontend"
PRINT "end build & copy frontend to public"

PRINT "material all done\n"
PRINT "commits: $VERSION.$COMMITS\n"

if [ ! -z "$DOCKER_USER" -a ! -z "$DOCKER_PASS" ]; then
  PRINT "packing docker image\n"
  docker login docker.pkg.github.com -u "$DOCKER_USER" -p "$DOCKER_PASS" || exit 1
  PRINT "docker had login\n"
  TAG="docker.pkg.github.com/shaohung001/ginfast/ginfast:$VERSION.$COMMITS"
  PRINT $TAG
  docker build -t $TAG .
  docker push $TAG
fi

exit 0

