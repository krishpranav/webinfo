# webinfo
A web information gathering tool made in go - DNS / Subdomains / Ports / Directories enumeration

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

# Installation
```
git clone https://github.com/krishpranav/webinfo
cd webinfo
go get
make
```

Examples
----------

- DNS enumeration:
    
    - `webinfo dns -target target.domain`
    - `webinfo dns -o txt -target target.domain`
    - `webinfo dns -o html -target target.domain`
    - `webinfo dns -plain -target target.domain`

- Subdomains enumeration:

    - `webinfo subdomain -target target.domain`
    - `webinfo subdomain -w wordlist.txt -target target.domain`
    - `webinfo subdomain -o txt -target target.domain`
    - `webinfo subdomain -o html -target target.domain`
    - `webinfo subdomain -i 400 -target target.domain`
    - `webinfo subdomain -i 4** -target target.domain`
    - `webinfo subdomain -c -target target.domain`
    - `webinfo subdomain -db -target target.domain`
    - `webinfo subdomain -plain -target target.domain`

- Directories enumeration:

    - `webinfo dir -target target.domain`
    - `webinfo dir -w wordlist.txt -target target.domain`
    - `webinfo dir -o txt -target target.domain`
    - `webinfo dir -o html -target target.domain`
    - `webinfo dir -i 500,401 -target target.domain`
    - `webinfo dir -i 5**,401 -target target.domain`
    - `webinfo dir -c -target target.domain`
    - `webinfo dir -plain -target target.domain`

- Ports enumeration:
      
    - Default (all ports, so 1-65635) `webinfo port -target target.domain`
    - Specifying ports range `webinfo port -p 20-90 -target target.domain`
    - Specifying starting port (until the last one) `webinfo port -p 20- -target target.domain`
    - Specifying ending port (from the first one) `webinfo port -p -90 -target target.domain`
    - Specifying single port `webinfo port -p 80 -target target.domain`
    - Specifying output format (txt)`webinfo port -o txt -target target.domain`
    - Specifying output format (html)`webinfo port -o html -target target.domain`
    - Specifying multiple ports `webinfo port -p 21,25,80 -target target.domain`
    - Specifying common ports `webinfo port -common -target target.domain`
    - Print only results `webinfo port -plain -target target.domain`

- Full report:
      
    - Default (all ports, so 1-65635) `webinfo report -target target.domain`
    - Specifying ports range `webinfo report -p 20-90 -target target.domain`
    - Specifying starting port (until the last one) `webinfo report -p 20- -target target.domain`
    - Specifying ending port (from the first one) `webinfo report -p -90 -target target.domain`
    - Specifying single port `webinfo report -p 80 -target target.domain`
    - Specifying output format (txt)`webinfo report -o txt -target target.domain`
    - Specifying output format (html)`webinfo report -o html -target target.domain`
    - Specifying directories wordlist `webinfo report -wd dirs.txt -target target.domain`
    - Specifying subdomains wordlist `webinfo report -ws subdomains.txt -target target.domain`
    - Specifying status codes to be ignored in directories scanning `webinfo report -id 500,501,502 -target target.domain`
    - Specifying status codes to be ignored in subdomains scanning `webinfo report -is 500,501,502 -target target.domain`
    - Specifying status codes classes to be ignored in directories scanning `webinfo report -id 5**,4** -target target.domain`
    - Specifying status codes classes to be ignored in subdomains scanning `webinfo report -is 5**,4** -target target.domain`
    - Use also a web crawler for directories enumeration `webinfo report -cd -target target.domain`
    - Use also a web crawler for subdomains enumeration `webinfo report -cs -target target.domain`
    - Use also a public database for subdomains enumeration `webinfo report -db -target target.domain`
    - Specifying multiple ports `webinfo report -p 21,25,80 -target target.domain`
    - Specifying common ports `webinfo report -common -target target.domain`
