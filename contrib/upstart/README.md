## Build

    godep go build -o kesshad main.go
    cp kesshad /usr/sbin


## Make User

    sudo useradd -s /sbin/nologin kessha


## Create RSA key

    mkdir /etc/kessha
    ssk-keygen -t rsa -f /etc/kessha/id_rsa


## Create Service

    `/etc/systemd/system/kesshad.service`:
    description "Kesshad SSH Honeypot"

    start on runlevel [2345]
    stop on runlevel [!2345]

    respawn
    respawn limit 5 60
    setuid kessha

    script
        [ -r /etc/default/kesshad ] && . /etc/default/kesshad
        exec /usr/sbin/kesshad > /var/log/kesshad.log 2>&1
    end script


# Add Twitter Secrets to Defaults

`/etc/default/kesshad`:

    export TWITTER_CONSUMER_KEY=
    export TWITTER_CONSUMER_SECRET=
    export TWITTER_ACCESS_TOKEN=
    export TWITTER_ACCESS_SECRET=


## Start Service

    sudo start kesshad
