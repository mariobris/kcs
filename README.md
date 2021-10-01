# kcs
Changes KUBECONFIG environment variable in the current shell. Inspired by https://github.com/jsen-/kx and no extra featuers added. I just needed binary for ARM and did not want to play with rust.

## Installation
You can clone repository and build binary with `go build` and copy it into directory in your $PATH. Add to your .zshrc, .bashrc, ... following section
```
# export KCS_KUBEDIR=~/.kube/configs
kx() {
    eval $(kcs)
}
```

## Usage
```
kx
```

## Configuration
- `$KCS_KUBEDIR` will let you choose KUBECONFIG from this directory
- to configure interactive look, see https://github.com/peco/peco
- to change default collors create $HOME/.peco/config.json file
```
mkdir -p $HOME/.peco
echo "{
    "Prompt": "[kube]",
    "Style": {
        "Basic": ["on_black", "white"],
        "SavedSelection": ["bold", "on_black", "green"],
        "Selected": ["bold", "on_black", "yellow"],
        "Query": ["yellow", "bold"],
        "Matched": ["yellow", "on_black"]
    }
}" > $HOME/.peco/config.json
```
