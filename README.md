# liebot
    I'm the guy who sucks

Liebot is a Slack bot that searches ohnorobot and the Achewood comic index page.

It uses the [Slack slash command API](https://api.slack.com/slash-commands).

Example usage:

![Example image]
(http://i.imgur.com/zrLtnYg.png)

## To install

### On some WAN-facing host
    go get github.com/graysonchao/liebot
    go install github.com/graysonchao/liebot
    nohup liebot > liebot.txt &
    
Then reverse proxy the listening port (pretty sure it's 8443) to 80 or 443.

Test request:
    curl -XPOST -d "text=I'm the guy who sucks" hostname.com/comic

### In Slack
Set up a [slash command integration](https://api.slack.com/slash-commands) and point it at `/comic` on the host that's serving liebot.
