SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

make_go_package_tar() {
  GO_PACKAGES=(
    core
    monitors
    observers
    utils
  )

  # A hack to simplify Dockerfile since Dockerfile doesn't support copying
  # multiple directories without flattening them out
  (cd $SCRIPT_DIR/.. && tar -cf $SCRIPT_DIR/go_packages.tar main.go Makefile scripts/{make-templates,collectd-template-to-go} ${GO_PACKAGES[@]})
}

extra_cflags_build_arg() {
  # If this isn't true then let build use default
  if [[ $DEBUG == 'true' ]]
  then
    echo "--build-arg extra_cflags='-g -O0'"
  fi
}

do_docker_build() {
  local tag=$1
  local dockerfile=$2

  make_go_package_tar

  docker build \
    -t $tag \
    -f $dockerfile \
    --label agent.version=$($SCRIPT_DIR/../VERSIONS agent_version) \
    --label collectd.version=$($SCRIPT_DIR/../VERSIONS collectd_version) \
    --build-arg DEBUG=$DEBUG \
    $(extra_cflags_build_arg) \
    $SCRIPT_DIR/.. 
}