# exchangeME
ExchangeMe is an application that allows users to track (fictional) Gold investments and get real-time information on Gold prices. It is written in Go with [Fyne](https://developer.fyne.io/).


##### An illustration of v1.0.0
![alt text](https://github.com/petrostrak/exchangeME/blob/main/exchangeME.png)


##### To build exchangeME from source
    $ fyne package -appVersion 1.0.0 -appBuild 1 -name exchangeME -release

##### To install exchangeME on debian:

* Extract exchangeME.tar.xz
    
    `mkdir exchangeME && tar -xvf exchangeME.tar.xz -C exchangeME`

* Mount and run make file:

    `cd exchangeME && sudo make install`
