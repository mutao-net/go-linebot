# go-linebot

## Heroku-Cli install
# see https://devcenter.heroku.com/ja/articles/heroku-cli
$ brew tap heroku/brew && brew install heroku

## Heroku GOVERSION
## see https://devcenter.heroku.com/ja/articles/go-support#go-versions
##     https://devcenter.heroku.com/ja/articles/go-dependencies-via-govendor#build-configuration

## gomodで指定する場合
$ vi go.mod
```
// +heroku goVersion go1.16
```