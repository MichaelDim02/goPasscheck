<img src="https://mcdim.xyz/media/passchecklogo1.png" alt="goPasscheck logo" style="height: 200px;"/>

# goPasscheck
<p float="left">
<img src="https://img.shields.io/badge/version-0.1-blue"/>
<img src="https://img.shields.io/badge/language-Go-lightblue">
<img src="https://img.shields.io/badge/license-GNU%20GPLv3+-yellow">

How strong is your password? **goPasscheck** can be used to intelligently assess the strength of a password, taking into account multiple variables and statistics provided by `lykan`.

## Installation
To install, run the following commands
```
$ git clone https://github.com/MicaelDim02/gopasscheck && cd gopasscheck/src
$ sudo make install
```
To uninstall, run the following in `gopasscheck/src`:
```
$ sudo make uninstall
```

## Usage
Check a password:
```
gopasscheck -p password
```
Optionally, you can provide a dictionary: 
```
gopasscheck -p password -w /usr/share/dict/words
```
