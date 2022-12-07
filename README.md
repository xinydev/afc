# Ask for Confirmation

Conduct secondary confirmation to the user when executing sensitive operations (such as kubectl remove, kubectl apply, and rm -rf) to avoid operation problems.

It is simple to become lost while interacting with several Kubernetes at once.
Unexpected losses may result, for instance, if you fail to switch clusters before applying or deleting.

Before executing sensitive operations (apply, delete), AFC will display a notice asking the user to confirm before doing the actual operations. minimize the likelihood of operational errors

AFC is very practical and works with all instructions. If the CLI was created by Cobra, native shell completion is also supported.

[中文说明](README-CN.md)

## Install

```shell
go install github.com/xinydev/afc@latest

# completion for zsh. run `afc completion --help` for other shells
afc completion zsh > "${fpath[1]}/_afc"
```

## Usage

### kubectl

### With alias

add below line in  ~/.zshrc

```shell
alias k="afc --danger='delete,apply' --notice='kubectl config get-contexts' --cmd='kubectl' run -- "
```

```shell
source ~/.zshrc
```

```shell
➜ k get pods
NAME     READY   STATUS    RESTARTS       AGE
nginx    1/1     Running   4 (2d2h ago)   57d
nginx2   1/1     Running   4 (2d2h ago)   57d

➜ k delete pods -n default nginx
CURRENT   NAME      CLUSTER   AUTHINFO   NAMESPACE
*         kind-k1   kind-k1   kind-k1    
          kind-k2   kind-k2   kind-k2    
Confirm???n # press n+enter 
exit...

➜ k delete pods -n default nginx 
CURRENT   NAME      CLUSTER   AUTHINFO   NAMESPACE
*         kind-k1   kind-k1   kind-k1    
          kind-k2   kind-k2   kind-k2    
Confirm???y # press y+enter
pod "nginx" deleted


```

### without alias

```shell

➜  afc --danger='delete,apply,get' --notice='kubectl config get-contexts' --cmd='kubectl' run -- get pods
CURRENT   NAME      CLUSTER   AUTHINFO   NAMESPACE
*         kind-k1   kind-k1   kind-k1    
          kind-k2   kind-k2   kind-k2    
Confirm???y
NAME     READY   STATUS    RESTARTS       AGE
nginx    1/1     Running   4 (2d2h ago)   57d
nginx2   1/1     Running   4 (2d2h ago)   57d

```

### helm

```shell
alias helm="afc --danger='delete,apply,get' --notice='kubectl config get-contexts' --cmd='helm' run -- "

````

### rm -rf

```shell
alias rm="afc --danger='-rf,-fr' --notice='echo dangers!!!' --cmd='rm' run -- "
```

```shell

mkdir folder1
rm -rf folder1
dangers!!!
Confirm???n
exit...

```

## shell completion

You must first set the afc shell completion if you wish to use the kubectl(helm...) shell completion.

```shell
afc completion --help
```
