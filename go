#!/bin/bash

export IMAGE_NAME=${IMAGE_NAME:-'dind-ruby'}
export TAG=2.5.1

goal_build() {
  docker build . -t "${IMAGE_NAME}:${TAG}"
}

goal_test() {
  received=$(docker run --privileged --rm -it "${IMAGE_NAME}:${TAG}" sh -c "ruby -v")
  if [[ $received =~ .*$TAG.* ]]; then
    exit 0
  else
    echo "expected[$TAG] did not match actual version ${received}"
    exit 1
  fi
}

goal_help() {
  echo "usage: $0 <goal>

    goal:

    build                    -- Build the image
    test                     -- Test that the image is built correctly
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
