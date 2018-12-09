package main

import (
	"flag"
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"

	"github.com/pouyatafti/hkwol/wol"
)

func wake(macStr string) {
	log.Println("waking " + macStr + "...")
	err := wol.Broadcast(macStr)

	if err != nil {
		log.Println("warning: " + err.Error())
	}
}

func main() {
	var pinStr, macStr string
	
	flag.StringVar(&pinStr, "pin", "12345678", "HomeKit pairing PIN")
	flag.StringVar(&macStr, "mac", "00:00:00:00:00:00", "MAC address")
	flag.Parse()

	log.Println("PIN: " + pinStr + ", MAC: " + macStr)

	info := accessory.Info{
		Name:         "WOL",
		Manufacturer: "pouya@nohup.io",
	}

	acc := accessory.NewSwitch(info)

	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			wake(macStr)
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: pinStr}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
