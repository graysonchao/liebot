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
    
Then reverse proxy 443 to the port that liebot listens on (8443 IIRC).

Note that liebot _DOES NOT TERMINATE SSL!_ That's why you use a reverse proxy like Nginx and terminate SSL there instead. Seriously, DON'T put liebot on an open port and just let it run - unless you're down to get MITMed.

Test request:

    curl -XPOST -d "text=I'm the guy who sucks" hostname.com/comic

### In Slack
Set up a [slash command integration](https://api.slack.com/slash-commands) and point it at `/comic` on the host that's serving liebot.
