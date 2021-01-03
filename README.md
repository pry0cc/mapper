# Mapper - a tool to help the distributed scannning of hosts
I found myself frequently having a lot of masscan outputs in the form of host:port, and I wanted to scan them line by line with nmap service scanning, mapper helps you do that.

### Installation

If you have a properly configured GOPATH and $GOPATH/bin is in your PATH, then run this command for a one-liner install, thank you golang!
```
go get -u github.com/pry0cc/mapper
```


#### Cat a very large unsorted wordlist.

#### port.txt
```
8.8.8.8:53
8.8.8.8:80
8.8.8.8:443
1.1.1.1:53
1.1.1.1:80
```

```
cat port.txt | mapper
```
This will start a separate nmap process for each line for example `§8.8.8.8:53` will translate to `nmap -T4 -sV -p53 8.8.8.8 -oX output/8.8.8.8_53.xml`
This will create an output directory and place output in the format of `ip_port.xml`


Enjoy :)
