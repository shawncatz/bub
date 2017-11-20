git + hub = github
==================

hub is a command line tool that wraps `git` in order to extend it with extra
features and commands that make working with GitHub easier.

``` sh
$ hub clone rtomayko/tilt

# expands to:
$ git clone git://github.com/rtomayko/tilt.git
```

hub is best aliased as `git`, so you can type `$ git <command>` in the shell and
get all the usual `hub` features. See "Aliasing" below.


Installation
------------

Dependencies:

* **git 1.7.3** or newer

#### Homebrew

`hub` can be installed through Homebrew:

``` sh
$ brew install hub
$ hub version
git version 1.7.6
hub version 2.2.3
```

#### Chocolatey

`hub` can be installed through [Chocolatey](https://chocolatey.org/) on Windows.

``` sh
> choco install hub
> hub version
git version 2.7.1.windows.2
hub version 2.2.3
```

#### Fedora Linux

On Fedora you can install `hub` through DNF:

``` sh
$ sudo dnf install hub
$ hub version
git version 2.9.3
hub version 2.2.9
```

#### Standalone

`hub` can be easily installed as an executable. Download the latest
[compiled binaries](https://github.com/shawncatz/bub/releases) and put it anywhere
in your executable path.

#### Source

To install hub from source:

``` sh
$ git clone https://github.com/shawncatz/bub.git
$ cd hub
$ make install prefix=/usr/local
```

Prerequisites for compilation are:

* `make`
* [Go 1.8+](http://golang.org/doc/install)
* Ruby 1.9+ with Bundler - for generating man pages

If you don't have `make`, Ruby, or want to skip man pages (for example, if you
are on Windows), you can build only the `hub` binary:

``` sh
$ ./script/build
```

You can now move `bin/hub` to somewhere in your PATH.

Finally, if you've done Go development before and your `$GOPATH/bin` directory
is already in your PATH, this is an alternative installation method that fetches
hub into your GOPATH and builds it automatically:

``` sh
$ go get github.com/shawncatz/bub
```

Aliasing
--------

Using hub feels best when it's aliased as `git`. This is not dangerous; your
_normal git commands will all work_. hub merely adds some sugar.

`hub alias` displays instructions for the current shell. With the `-s` flag, it
outputs a script suitable for `eval`.

You should place this command in your `.bash_profile` or other startup script:

``` sh
eval "$(hub alias -s)"
```

#### PowerShell

If you're using PowerShell, you can set an alias for `hub` by placing the
following in your PowerShell profile (usually
`~/Documents/WindowsPowerShell/Microsoft.PowerShell_profile.ps1`):

``` sh
Set-Alias git hub
```

A simple way to do this is to run the following from the PowerShell prompt:

``` sh
Add-Content $PROFILE "`nSet-Alias git hub"
```

Note: You'll need to restart your PowerShell console in order for the changes to be picked up.

If your PowerShell profile doesn't exist, you can create it by running the following:

``` sh
New-Item -Type file -Force $PROFILE
```

### Shell tab-completion

hub repository contains tab-completion scripts for bash and zsh. These scripts
complement existing completion scripts that ship with git.

[Installation instructions](etc)

* [hub bash completion](https://github.com/shawncatz/bub/blob/master/etc/hub.bash_completion.sh)
* [hub zsh completion](https://github.com/shawncatz/bub/blob/master/etc/hub.zsh_completion)
* [hub fish completion](https://github.com/shawncatz/bub/blob/master/etc/hub.fish_completion)

Meta
----

* Home: <https://github.com/shawncatz/bub>
* Bugs: <https://github.com/shawncatz/bub/issues>
* Authors: <https://github.com/shawncatz/bub/contributors>
