# Ask for Confirmation

When performing sensitive operations(kubectl delete,kubectl apply,rm -rf),
conduct secondary confirmation to the user to prevent operation errors

When interacting with multiple kubernetes at the same time,
it is easy to get confused.   
For example, if you forget to switch clusters before apply or delete,
unexpected losses may be caused

AFC will prompt a message to prompt the current environment
before performing sensitive operations (apply, delete),  
and ask the user to enter y before perform real operations. Reduce the possibility of operational errors

Using AFC is very convenient and supports all commands. If the CLI is developed by cobra,
it also supports native shell completion

[中文说明](README-CN.md)

## Install

```shell
go install github.com/xinydev/afc@latest
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

If you want to use the shell completion of the kubectl(helm...),you need set the afc shell completion first

```shell
afc completion --help
```