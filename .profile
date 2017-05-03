# maintain: Kamontat Chantrachirathumrong
# version:  1.2.0
# since:    03/05/1017

[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*

# set default variable
export USER="Kamontat"
export LC_ALL=en_US.UTF-8
export LANG=en_US.UTF-8
# set bin library location
export PATH="/usr/bin:$PATH"                                                                 # user bin
export PATH="/bin:$PATH"                                                                     # bin
export PATH="/usr/sbin:$PATH"                                                                # user sbin
export PATH="/sbin:$PATH"                                                                    # sbin
export PATH="/usr/local/bin:$PATH"                                                           # local bin
export PATH="/usr/local/sbin:$PATH"                                                          # local sbin
export PATH="/usr/local/git/bin:$PATH"                                                       # git
export PATH="/usr/local/sbin:$PATH"                                                          # new local sbin folder
# set python3 location
export PATH="/Library/Frameworks/Python.framework/Versions/3.5/bin:$PATH"                    # python version 3.5
# set rvm, ruby, gems location
export PATH="$HOME.rvm/bin:$PATH"                                                            # rvm bin
export PATH="$HOME/.rvm/gems/ruby-2.4.0/bin:$PATH"                                           # gems ruby
export PATH="$HOME/.rvm/gems/ruby-2.4.0@global@/bin:$PATH"                                   # gems global ruby
export PATH="$HOME/.rvm/rubies/ruby-2.4.0/bin:$PATH"                                         # ruby
# java home setting
export JAVA_HOME="/Library/Java/JavaVirtualMachines/jdk1.8.0_152.jdk/Contents/Home:$PATH"    # java home
export PATH="$JAVA_HOME/bin:$PATH"                                                     	     # java path

# new key
alias vim='/usr/local/Cellar/vim/8.0.0589/bin/vim'
alias c='clear'
alias srm='sudo rm -rf'
alias lsa='ls -la'