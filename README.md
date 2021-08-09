# go-linebot

## Heroku-Cli install
see https://devcenter.heroku.com/ja/articles/heroku-cli

```
$ brew tap heroku/brew && brew install heroku
```

## Heroku GOVERSION
https://devcenter.heroku.com/ja/articles/go-support#go-versions
    
https://devcenter.heroku.com/ja/articles/go-dependencies-via-govendor#build-configuration

## gomodで指定する場合
$ vi go.mod

```
// +heroku goVersion go1.16
```
![image](https://user-images.githubusercontent.com/35423021/128667673-aeb72f8e-fedf-437f-852d-307c2bfc5a10.png)
