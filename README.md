# rename
A CLI written in golang to set aliases in bash/zsh.

This is completely over engineering the method of setting an alias but it was fun to do and a project for me when learning golang.

### How to use this tool

Clone or download this repo and cd into it. From there run `make build` which will build and install a go binary `rename` to your `$GOPATH`.

The binary expects two arguments, first the orginal command, followed by the alias you want to set.
This will work out of the box, simply running `rename "git status" gs` will edit your profile with a new alias.
However, the go program can't make any changes to the terminal that is running on (if I'm wrong here and there is a way please let me know!). So the change won't take affect until you source the ~/.zshrc or ~/.bashrc file, or open a new terminal. 

One way around it is to write a small bash function like 

``` 
function rename() {
    if [ "$#" -eq 2 ]; then
        $GOPATH/bin/rename $1 $2
        source ~/.zshrc
    else
        echo "rename takes in two arguments"
    fi
}
```

Now if you type `rename "echo hello", hello` the alias will be available on the same terminal!