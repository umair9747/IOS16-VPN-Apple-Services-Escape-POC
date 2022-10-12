# IOS16-VPN-Apple-Services-Escape-POC

I found a twitter account (https://twitter.com/mysk_co) mentioning that they experimented around iOS 16 and found out that it communicates with Apple services outside an active VPN tunnel. Worse, it leaks DNS requests. To make it easier for users to test out this issue and see if any of the IP addresses related to Apple are actually leaked when the device is connected to a VPN tunnel, I tried to develop a POC.

## Setup

This POC only works on windows and assumes that you have <a href="https://www.wireshark.org/">Wireshark(and tshark)</a> installed in C:\Program Files\Wireshark. If you have installed it in a different location, you can edit line no.77 in <a href="main.go">main.go</a> and run **go build**.

In order to setup your testing environment, the following steps can be followed:
* Enable mobile hotspot on windows by going to Settings > Network & Internet > Mobile Hotspot
* Connect your iPhone to this hotspot
* Identify the interface being used for this connection, you will need this later.

## Usage

```
Usage of ios-poc.exe:
  -d int
        duration for capturing the packets. (default 60)
  -i string
        interface to capture the packets from.

Example Usage:
PS C:\Users\Hp\OneDrive\Desktop\ios> .\ios-poc.exe -d 30 -i "Local Area Connection* 2"

APPLE VPN CONNECTION ESCAPE POC
Developed by 0x9747

#################
Changed to new directory: C:\Program Files\Wireshark
Running tshark for 30 seconds on Local Area Connection* 2
#################

Unqiue IP Addresses Discovered: [192.168.137.140 149.34.244.169 17.57.145.116]

Checking for VPN-related IP Addresses:
149.34.244.169

Apple IP Addresses Leaked:
17.57.145.116
```

## Tested On
This POC successfully worked under the following testing condition:
* Device: iPhone 13 (iOS 16)
* VPN: Proton VPN
* Apps Opened: Health, Maps, Wallet, Find My, Weather, iTunes Store

## Credits
This tool was inspired from the tweet by <a href="https://twitter.com/mysk_co/">Mysk</a> and I would like to thank them for this discovery.
