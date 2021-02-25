

- 查看php.ini和php-fpm.comf和nginx.conf的文件位置

  ```
  php -i |grep php.ini
  sudo php-fpm -t
  sudo nginx -t
  
  # mac 启动php-fpm:
  sudo php-fpm
  # mac 关闭php-fpm
  sudo killall php-fpm
  ```

  

- 查看是bash用户或者zsh用户

  ```
  echo $SHELL
  ```

  

- #### Homebrew从本地文件安装软件

  ```
  有时候会遇到有些文件curl下载不下来，导致brew不能正常工作，但是文件通过浏览器是可以下载的。
  
  以下是解决方案：
  
  1. 手动下载压缩文件（一定要下载brew提示下载失败的文件）；
  
  2. 执行brew --cache获取brew缓存路径，将下载的文件放入缓存目录；
  
  3. 重新执行安装命令，brew会发现缓存中有了文件，就不去下载了，OK。
  ```

  

- Home-brew 安装nginx和php

  ```
  地址：https://juejin.im/post/5c8fb28a6fb9a07103548318#heading-2
  
  nginx:
  Docroot is: /usr/local/var/www
  /usr/local/etc/nginx/nginx.conf
  
  php:
  php@7.2
  To enable PHP in Apache add the following to httpd.conf and restart Apache:
      LoadModule php7_module /usr/local/opt/php@7.2/lib/httpd/modules/libphp7.so
  
      <FilesMatch \.php$>
          SetHandler application/x-httpd-php
      </FilesMatch>
  
  Finally, check DirectoryIndex includes index.php
      DirectoryIndex index.php index.html
  
  The php.ini and php-fpm.ini file can be found in:
      /usr/local/etc/php/7.2/
  
  php@7.2 is keg-only, which means it was not symlinked into /usr/local,
  because this is an alternate version of another formula.
  
  If you need to have php@7.2 first in your PATH run:
    echo 'export PATH="/usr/local/opt/php@7.2/bin:$PATH"' >> ~/.zshrc
    echo 'export PATH="/usr/local/opt/php@7.2/sbin:$PATH"' >> ~/.zshrc
  
  For compilers to find php@7.2 you may need to set:
    export LDFLAGS="-L/usr/local/opt/php@7.2/lib"
    export CPPFLAGS="-I/usr/local/opt/php@7.2/include"
  
  
  To have launchd start php@7.2 now and restart at login:
    brew services start php@7.2
  Or, if you don't want/need a background service you can just run:
    php-fpm
  
  
  php目录 /usr/local/Cellar/php@7.1/版本号/bin
  php-fpm目录 /usr/local/Cellar/php@7.1/版本号/sbin
  配置目录 /usr/local/etc/php/7.1
  ```

  

