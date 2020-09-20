package config

const defaultYAML string = `
service:
    name: omo.bla.acm
    address: :9701
    ttl: 15
    interval: 10
logger:
    level: info
    dir: /var/log/msa/
client:
    retry: 3
    timeout: 2
msa:
    account: omo.msa.account
    group: omo.msa.group
`
