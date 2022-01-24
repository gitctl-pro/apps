# API Create

## kubebuilder
[kubebuilder](https://book.kubebuilder.io)

```shell
go mod init github.com/gitctl-pro/apps
kubebuilder init --domain gitctl.com
kubebuilder edit --multigroup=false
kubebuilder create api --group apps --version v1 --kind Application
kubebuilder create api --group apps --version v1 --kind Canary
go mod vendor 
make 
```
## go mod


## code-gen
1. 创建code-generator所需的文件
* apis/core/v1/doc.go
* apis/core/v1/register.go

doc.go
```
// +k8s:deepcopy-gen=package
// +groupName=core.sae
// +k8s:openapi-gen=true
package v1
  ```

register.go
  ```
package v1

import (
 "k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = GroupVersion

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
return SchemeGroupVersion.WithResource(resource).GroupResource()
}
  ```
2. *_types.go 添加genclient注释，注意scope：nonNamespaced为cluster
```
// +genclient
// +genclient:nonNamespaced
  ```
4. 执行./hack/update-codegen.sh，然后上级目录中找到生成的代码copy client