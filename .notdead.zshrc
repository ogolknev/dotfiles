########################
### NOTDEAD'S .ZSHRC ###
########################

source "$HOME/.config/broot/launcher/bash/br"

export PATH="$HOME/go/bin:$PATH"
export PATH="$HOME/bin:$PATH"
export PATH="$HOME/.config/scripts:$PATH"

export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # Загружает nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion" # Добавляет автодополнение

export PYENV_ROOT="$HOME/.pyenv"
export PATH="$PYENV_ROOT/bin:$PATH"
eval "$(pyenv init --path)"
eval "$(pyenv init -)"

if [[ ! -f $HOME/.zi/bin/zi.zsh ]]; then
  print -P "%F{33}▓ %F{160}Installing (%F{33}z-shell/zi%F{160})…%f"
  command mkdir -p "$HOME/.zi" && command chmod go-rwX "$HOME/.zi"
  command git clone -q --depth=1 --branch "main" https://github.com/z-shell/zi "$HOME/.zi/bin" && \
    print -P "%F{33}▓ %F{34}Installation successful.%f%b" || \
    print -P "%F{160}▓ The clone has failed.%f%b"
fi
source "$HOME/.zi/bin/zi.zsh"
autoload -Uz _zi
(( ${+_comps} )) && _comps[zi]=_zi

zicompinit

zi ice as"command" from"gh-r" \
  atclone"./starship init zsh > init.zsh; ./starship completions zsh > _starship" \
  atpull"%atclone" src"init.zsh"
zi light starship/starship

zi ice lucid wait has'fzf'
zi light Aloxaf/fzf-tab

zi ice lucid wait as'completion'
zi light zsh-users/zsh-completions

zi ice wait lucid atinit"ZI[COMPINIT_OPTS]=-C; zicompinit; zicdreplay"
zi light z-shell/F-Sy-H

zi ice wait lucid atload"!_zsh_autosuggest_start"
zi load zsh-users/zsh-autosuggestions

zi light z-shell/zsh-eza

zi ice has'zoxide'
zi light z-shell/zsh-zoxide

alias cat='bat'
alias rm='trash-put'