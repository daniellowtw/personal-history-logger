# Personal history logger

This is a simple link collector server that allows you to log when and what links you viewed. It is an experiment by me to collect some data about my browsing history and hopefully do some data analysis on it. Use at your own discretion.

# Installation

```
go get github.com/daniellowtw/personal-history-logger
./personal-history-logger -help
```

Please note that this should ideally be served behind some https reverse proxy so that you don't leak the site you are visiting in the packets.

# Using

## Posting

Sending request to the `/post?url=<base64 encoded url>` endpoint will capture the url, relevant meta tags of the website, and the timestamp.

## Bookmarklet

Visit `/bookmarklet` to drag the bookmarklet to your bookmark bar for easy logging.

## Retrieving

The entries are captured and saved in the log file sharded by days. There is no front end for viewing these.

# TODO

* Add some auth checking
* Add SSL support
