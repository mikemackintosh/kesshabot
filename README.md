# Ke$$ha BOT

Ke$$ha bot is a SSH honeybot. It is designed to tweet everytime someone maliciously attempts to access SSH on your server.


## Building

We bundled dependencies with `godeps`. You can use:

    godep go build -o kesshad main.go


## Usage

For twitter access, please export your Twitter variables:

    TWITTER_CONSUMER_KEY
    TWITTER_CONSUMER_SECRET
    TWITTER_ACCESS_TOKEN
    TWITTER_ACCESS_SECRET


## Security

Don't run this on port 22 directly as `root`. That would be dumb. Run this as an unpriviled user on a port `> 1024` and use IP tables to redirect the port:

    sudo iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 22 -j REDIRECT --to-port 2022

Of course, change to your variables accordingly.
