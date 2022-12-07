# AFC

在执行敏感操作的时候，向用户进行二次确认，防止操作错误

当同时要和多个kubernetes进行交互的时候，很容易搞混。比如忘记切换集群，直接进行apply或者delete，造成意外的损失

AFC会在执行敏感操作的时候(apply,delete)之前，提示一条消息，提示当前的环境，并且要求用户输入y才可以真正的进行执行，
减少操作错误的可能性

使用AFC很方便，支持所有的命令。如果cli是cobra开发的，那还支持原生的shell补全

## 安装

```shell
go install github.com/xinydev/afc@latest
```

## 使用

### kubectl

### 通过别名使用

在 ~/.zshrc增加

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
Confirm???n # 手动输入n
exit...

➜ k delete pods -n default nginx 
CURRENT   NAME      CLUSTER   AUTHINFO   NAMESPACE
*         kind-k1   kind-k1   kind-k1    
          kind-k2   kind-k2   kind-k2    
Confirm???y # 手动输入y
pod "nginx" deleted


```

### 直接使用

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

## shell补全

如果希望使用kubectl,helm等cobra开发的cli的shell补全功能，只需要给afc配置补全即可

```shell
afc completion --help
```
