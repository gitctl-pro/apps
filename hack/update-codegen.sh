#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

mkdir -p ./tmp
SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

#rm -rf "${APIS_PKG}/${GROUP}" && mkdir -p "${APIS_PKG}/${GROUP}" && cp -r "${APIS_PKG}/${VERSION}/" "${APIS_PKG}/${GROUP}"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
bash "${CODEGEN_PKG}"/generate-groups.sh "client,informer,lister" \
  github.com/gitctl-pro/apps/client github.com/gitctl-pro/apps/apis \
  "apps:v1 core:v1" \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  --output-base "${SCRIPT_ROOT}/../../.." \
  --v 6 \
   --output-base "${SCRIPT_ROOT}/tmp"

rm -rf ./client
mv ./tmp/github.com/gitctl-pro/apps/client  ./
rm -rf ./tmp