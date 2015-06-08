#!/bin/bash

#
# This script builds the application from source for multiple platforms.
set -e

# Get the parent directory of where this script is.
DIR=$(realpath $(dirname $0)/..)

# Change into that directory
cd $DIR

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# If its dev mode, only build for ourself
if [ "${TF_DEV}x" != "x" ]; then
    XC_OS=${XC_OS:-$(go env GOOS)}
    XC_ARCH=${XC_ARCH:-$(go env GOARCH)}
fi

# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"386 amd64 arm"}
XC_OS=${XC_OS:-linux darwin windows freebsd openbsd}

# Install dependencies
echo "==> Getting dependencies..."
go get -u -fix ./...

# Delete the old dir
echo "==> Removing old directory..."
rm -Rf ${DIR}/bin
mkdir -p ${DIR}/bin

# Build!
echo "==> Building..."
set +e
gox \
  -os="${XC_OS}" \
  -arch="${XC_ARCH}" \
  -ldflags "-X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" \
  -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}-${GIT_COMMIT:0:8}" \
  ./...
set -e

# Copy our OS/Arch to the bin/ directory
echo "==> Copying binaries for this platform..."
BIN="$(basename ${DIR})-${GIT_COMMIT:0:8}"
BINDIR="${DIR}/pkg/$(go env GOOS)_$(go env GOARCH)"
ln -sv "${BINDIR}/${BIN}.exe" "${DIR}/bin/hd.exe" 2>/dev/null
ln -sv "${BINDIR}/${BIN}" "${DIR}/bin/hd" 2>/dev/null

# Done!
echo
echo "==> Results:"
ls -hl "${DIR}/bin" 
