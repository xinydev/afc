# Ask for Confirmation

When executing sensitive operations (such as kubectl remove, kubectl apply, and rm -rf), it is important to conduct secondary confirmation with the user to avoid operational problems. Interacting with several Kubernetes at once can be confusing and may result in unexpected losses, such as failing to switch clusters before applying or deleting.

To minimize the likelihood of operational errors, AFC displays a notice asking the user to confirm before executing sensitive operations like apply and delete. AFC is practical and works with all instructions. If the CLI was created by Cobra, native shell completion is also supported.

[中文说明](README-CN.md)

## Install

```shell
go install github.com/xinydev/afc@latest

# completion for zsh. run `afc completion --help` for other shells
afc completion zsh > "${fpath[1]}/_afc"
```

## Usage

### kubectl

#### With alias

Add the following line to `~/.zshrc`:

```shell
alias k="afc --danger='delete,apply' --notice='kubectl config get-contexts' --cmd='kubectl' run -- "
```

Then run:

```shell
source ~/.zshrc
```

Now you can use `k` instead of `kubectl`. For example:

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

#### Without alias

```shell
➜ afc --danger='delete,apply,get' --notice='kubectl config get-contexts' --cmd='kubectl' run -- get pods

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
```

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

## Shell completion

To use the kubectl (helm...) shell completion, you must first set the afc shell completion.

```shell
afc completion --help
```
