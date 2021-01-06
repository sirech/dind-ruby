#!/bin/bash

set -e
set -o nounset
set -o pipefail

export IMAGE_NAME=${IMAGE_NAME:-'dind-ruby'}
export TAG=2.7.1

goal_build() {
  docker build . -t "${IMAGE_NAME}"
}

goal_test() {
  # check for "deadlock; recursive locking (ThreadError)"
  docker run --privileged --rm -it "${IMAGE_NAME}" sh -c "irb -v" > /dev/null

  received=$(docker run --privileged --rm -it "${IMAGE_NAME}" sh -c "ruby -v")
  if [[ $received =~ .*$TAG.* ]]; then
    exit 0
  else
    echo "expected[$TAG] did not match actual version ${received}"
    exit 1
  fi
}

goal_publish() {
  docker tag "${IMAGE_NAME}" "${IMAGE_NAME}:${TAG}"
  docker push "${IMAGE_NAME}:${TAG}"
}

goal_help() {
  echo "usage: $0 <goal>

    goal:

    build                    -- Build the image
    test                     -- Test that the image is built correctly
    publish                  -- Publish the image to the hub
    "
  exit 1
}

main() {
  TARGET=${1:-}
  if [ -n "${TARGET}" ] && type -t "goal_$TARGET" &>/dev/null; then
    "goal_$TARGET" "${@:2}"
  else
    goal_help
  fi
}

main "$@"
