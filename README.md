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

![image](https://user-images.githubusercontent.com/35423021/128667605-3bcbb57a-99ec-40d7-ba23-a9fe376e4d3d.png)
