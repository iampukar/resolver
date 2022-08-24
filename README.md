# resolver
A light-weight tool for DNS resolution written in Go! 

## Installing

Requires [Go](https://golang.org/dl/)

```bash
go get github.com/iampukar/resolver
```

OR, 

```bash
git clone https://github.com/iampukar/resolver/
cd resolver
```

## How To Use:

Examples: 
- `go run resolver/dns.go URL_TO_RESOLVE`

## Working  

```
$ go run resolver/dns.go realmadrid.com
Getting DNS Records for realmadrid.com ...

[IP] 23.76.236.186

[MX] Error : lookup realmadrid.com on 127.0.0.53:53: no such host

[NS] pdns80.ultradns.biz.
[NS] pdns80.ultradns.org.
[NS] pdns80.ultradns.com.
[NS] pdns80.ultradns.net.

[CNAME] realmadrid.com.

[TXT] google-site-verification=5L1fWU_QgAtKuh4-mYulZxvQamlaZbp7zZaQHLYPFZE
[TXT] facebook-domain-verification=k56h93u7s1iwbkuvuqi5z1ipzwx43w
[TXT] _globalsign-domain-verification=vLgoogWqACpf9j2zykchDN2wCyTtXrayGZtsnrpJdN
[TXT] 077txg62782t4yh8mqbrg5q50jks5fgj
[SPF] v=spf1 include:_spf-ssg-a.microsoft.com

[DMARC] v=DMARC1; p=none; rua=mailto:globaldpadmin@corp.realmadrid.com; ruf=mailto:globaldpadmin@corp.realmadrid.com; fo=1
```

```
$ go run resolver/dns.go google.com
Getting DNS Records for google.com ...

[IP] 2404:6800:4009:81b::200e
[IP] 142.250.199.174

[MX] 10 smtp.google.com.

[NS] ns3.google.com.
[NS] ns1.google.com.
[NS] ns4.google.com.
[NS] ns2.google.com.

[CNAME] google.com.

[TXT] atlassian-domain-verification=5YjTmWmjI92ewqkx2oXmBaD60Td9zWon9r6eakvHX6B77zzkFQto8PQ9QsKnbf4I
[TXT] MS=E4A68B9AB2BB9670BCE15412F62916164C0B20BB
[TXT] docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e
[TXT] google-site-verification=wD8N7i1JTNTkezJ49swvWW48f8_9xveREV4oB-0Hf5o
[TXT] webexdomainverification.8YX6G=6e6922db-e3e6-4a36-904e-a805c28087fa
[TXT] docusign=1b0a6754-49b1-4db5-8540-d2c12664b289
[TXT] globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=
[TXT] facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95
[SPF] v=spf1 include:_spf.google.com ~all

[DMARC] v=DMARC1; p=reject; rua=mailto:mailauth-reports@google.com
```


```
$ go run resolver/dns.go dummydummywebsite.com
Getting DNS Records for dummydummywebsite.com ...

[IP] Error : lookup dummydummywebsite.com: no such host

[MX] Error : lookup dummydummywebsite.com on 127.0.0.53:53: no such host

[NS] Error : lookup dummydummywebsite.com on 127.0.0.53:53: no such host

[CNAME] Error : lookup dummydummywebsite.com: no such host
[TXT] Error : lookup dummydummywebsite.com on 127.0.0.53:53: no such host

[TXT] Error : lookup _dmarc.dummydummywebsite.com on 127.0.0.53:53: no such host
```
