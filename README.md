# hkwol

simple program serving as a wake-on-lan switch for Apple HomeKit.  I run this on a Raspberry Pi to wake my boxes with voice commands to Siri. 

usage is simple.  The following lines build and launch an executable that introduces a new device named "WOL" you can control using Apple's HomeKit, to send a magic wake-on-lan packet to the provided MAC address. You can use Apple's Home app to add "WOL" as a new device, with manual pairing using the specified PIN.

```
go build cmd/hkwold.go
./hkwold -pin <pairing PIN> -mac <MAC address>
```
