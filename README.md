# DNS:Recap

`dnsrecap` memorize relation of domain name and IP address of captured DNS packet and provide HTTP API to look up domain name by IP address.

```bash
% go get github.com/m-mizutani/dnsrecap
% sudo ./dnsrecap &
% curl -v https://google.com
---- snip ----
* Connected to google.com (172.217.174.110) port 443 (#0)
---- snip ----
% curl http://localhost:5080/addr/172.217.174.110 | jq
{
  "addr": "172.217.174.110",
  "records": [
    {
      "type": "A",
      "name": "google.com",
      "data": "172.217.174.110",
      "time": "2020-07-21T23:17:00.711518Z"
    }
  ]
}
```
