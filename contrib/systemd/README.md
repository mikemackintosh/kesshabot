## Build

    godep go build -o kesshad main.go
    cp kesshad /usr/sbin


## Make User

    sudo adduser kessha -s /sbin/nologin


## Create RSA key

    mkdir /etc/kessha
    ssk-keygen -t rsa -f /etc/kessha/id_rsa


## Create Service

`/etc/systemd/system/kesshad.service`:

    [Unit]
    Description=kessha Service
    After=network.target

    [Service]
    EnvironmentFile=-/etc/default/kesshad
    Type=simple
    User=kessha
    ExecStart=/usr/sbin/kesshad
    Restart=on-failure
    RestartSec=5
    StartLimitInterval=60s
    StartLimitBurst=3

    [Install]
    WantedBy=multi-user.target


## Add Twitter Secrets to Defaults

`/etc/default/kesshad`:

    export TWITTER_CONSUMER_KEY=
    export TWITTER_CONSUMER_SECRET=
    export TWITTER_ACCESS_TOKEN=
    export TWITTER_ACCESS_SECRET=


## Start Service

    sudo systemctl start kesshad
