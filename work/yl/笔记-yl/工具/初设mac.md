```
推荐设置
#1.1 设置 zsh 为默认
chsh -s $(which zsh)
#1.2 安装 Git 客户端
打开 Git 官方下载页面 ，通过 Binary installer 下载安装

为啥不通过 brew 安装？因为现在还没有安装 brew 呀。。。

#1.3 安装 oh my zsh
需要先安装 git

sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"

安装字体

# clone in any tmp dir
git clone https://github.com/powerline/fonts.git --depth=1
# install
cd fonts
./install.sh
# clean-up a bit
cd ..
rm -rf fonts

安装 oh my zsh 插件

# 命令自动完成插件
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
# 命令高亮插件
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
编辑 ~/.zshrc

ZSH_THEME="agnoster"

...

plugins=(
  git
  extract
  zsh-autosuggestions
  zsh-syntax-highlighting
)



#1.4 安装 Homebrew
https://brew.sh/index_zh-cn
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
方法二、使用国内镜像安装

cd ~
curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install >> brew_install
编辑brew_install文件
#BREW_REPO = "https://github.com/Homebrew/brew".freeze
BREW_REPO = "git://mirrors.ustc.edu.cn/brew.git".freeze
#CORE_TAP_REPO = "https://github.com/Homebrew/homebrew-core".freeze
CORE_TAP_REPO = "git://mirrors.ustc.edu.cn/homebrew-core.git".freeze  # 没有不用新增
安装

/usr/bin/ruby ~/brew_install
此时，应该会停在:

==> Tapping homebrew/core

Cloning into '/usr/local/Homebrew/Library/Taps/homebrew/homebrew-core'...
出现这个原因是因为源不通，代码来不下来，解决方法就是更换国内镜像源：

手动执行下面这句命令，更换为中科院的镜像：

git clone git://mirrors.ustc.edu.cn/homebrew-core.git/ /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core --depth=1
安装 brew cask (可选)
brew cask
然后替换源

几个镜像:

https://git.coding.net/homebrew/homebrew.git - Coding
https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git - 清华
https://mirrors.ustc.edu.cn/brew.git - 中科大
cd "$(brew --repo)"
git remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git

cd "$(brew --repo)/Library/Taps/homebrew/homebrew-core"
git remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-core.git

# 默认不安装cask 有需要的可以替换
# 使用 brew cask 安装
cd "$(brew --repo)/Library/Taps/homebrew/homebrew-cask"
git remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-cask.git
vi ~/.zshrc

# Brew Bottle Mirrors
export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.ustc.edu.cn/homebrew-bottles
#1.5 安装 NodeJS 版本管理 （可选）
https://github.com/nvm-sh/nvm/blob/master/README.md#installation-and-update
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.34.0/install.sh | bash
设置 nvm 下载镜像地址

vi ~/.zshrc
# NVM mirrors
export NVM_NODEJS_ORG_MIRROR=https://npm.taobao.org/mirrors/node
export NVM_IOJS_ORG_MIRROR=https://npm.taobao.org/mirrors/iojs
设置 npm 镜像

vi ~/.npmrc

registry=https://registry.npm.taobao.org/
disturl=https://npm.taobao.org/dist
sass_binary_site=https://npm.taobao.org/mirrors/node-sass
electron_mirror=https://npm.taobao.org/mirrors/electron/
puppeteer_download_host=https://npm.taobao.org/mirrors
chromedriver_cdnurl=https://npm.taobao.org/mirrors/chromedriver
operadriver_cdnurl=https://npm.taobao.org/mirrors/operadriver
phantomjs_cdnurl=https://npm.taobao.org/mirrors/phantomjs
selenium_cdnurl=https://npm.taobao.org/mirrors/selenium
node_inspector_cdnurl=https://npm.taobao.org/mirrors/node-inspector
安装 nodejs

# 查看远程版本
nvm ls-remote
# 安装对应版本
nvm install v8.16.1
nvm install v10.16.3
#1.6 安装 GoLang 环境
下载安装包 https://golang.org/doc/install ，版本 1.12+

配置环境变量

vi ~/.zshrc

# GoLang
export GOROOT="/usr/local/go"
export GOPATH="$HOME/.go"
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
#export GO111MODULE=on # 建议在 IDE 中开启，不开启全局
export GOPROXY=https://goproxy.io
export GOPRIVATE=git.medlinker.com
source ~/.zshrc

go version

#必要软件
#Docker
官方下载 Docker Desktop ，并安装
国内阿里云下载 Docker Desktop
#IDE
推荐使用 IntelliJ 全家桶，不强制要求，但是后续教程均以此为基础。

PHP：PHPStorm
GoLang: GoLand
```

