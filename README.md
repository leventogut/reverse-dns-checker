# reverse-dns-checker

This small program checks if the PTR record matches the hostname and reports it.


# build

go build -o reverse-dns-checker *.go

# example usage
```
buntu:~/environment/go/src/github.com/leventogut/reverse-checker (master) $ ./reverse-dns-checker -domain google.com
Domain:  google.com.
A Record:  74.125.193.100
PTR for 74.125.193.100: ig-in-f100.1e100.net.
A Record:  74.125.193.101
PTR for 74.125.193.101: ig-in-f101.1e100.net.
A Record:  74.125.193.139
PTR for 74.125.193.139: ig-in-f139.1e100.net.
A Record:  74.125.193.138
PTR for 74.125.193.138: ig-in-f138.1e100.net.
A Record:  74.125.193.102
PTR for 74.125.193.102: ig-in-f102.1e100.net.
A Record:  74.125.193.113
PTR for 74.125.193.113: ig-in-f113.1e100.net.
A Record:  2a00:1450:400b:c01::71
No PTR found

Match does NOT exists
ubuntu:~/environment/go/src/github.com/leventogut/reverse-checker (master) $ ./reverse-dns-checker -domain s165-160.m1.turkishpolicy.com
Domain:  s165-160.m1.turkishpolicy.com.
A Record:  104.247.165.160
PTR for 104.247.165.160: s165-160.m1.turkishpolicy.com.

Match exists
ubuntu:~/environment/go/src/github.com/leventogut/reverse-checker (master) $ 
```
