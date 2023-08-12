<img align="center" width="800px" width="400px" src="./assets/img/logo.png">

# PasswordME

PasswordME is a tiny cross-platform password manager, which can generate passwords for different services using your favorite password

**The key features:**
- hack-resistant
- determented
- on-flight generation
- save password in encrypt file (soon)

## Overview
<img align="center" width="450" width="450" src="./assets/img/scheme.jpg">

## Problem
I'm sure most of us have one good password that we came up with once and now use everywhere, in all services. 
Clearly this is a bad thing.

With the streamlining of technology, hackers are finding more and more vulnerabilities in the security systems of the biggest companies, let alone the smaller ones. That's why we often hear about a database leak.

And now imagine if your only "good" password was leaked in some database. I think you know what a malicious person can do with it.

**PasswordME will help to solve this problem.**

## Implementation Status
- [x] Fyne GUI v1
- [x] On-Flight Generation
- [x] Passwords Repository
- [x] Save just-generated password 
- [x] React GUI v2
- [ ] PasswordME CLI v1

## Installing 
```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
git clone https://github.com/avelex/passwordme
cd passwordme/
make build
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
