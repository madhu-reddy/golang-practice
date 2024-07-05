package main

import (
	"fmt"
	"scientiamobile/wurfl"
)

func main() {

	var wengine *wurfl.Wurfl
	var device *wurfl.Device

	wengine, err := wurfl.Create("/usr/share/wurfl/wurfl.zip", nil, nil, -1, wurfl.WurflCacheProviderLru, "100000")

	if err != nil {
		fmt.Println(err)
	}

	// ua := "Dalvik/1.6.0 (Linux; U; Android 4.3; SM-N900T Build/JSS15J)"
	ua := "Mozilla/5.0 (Linux; Android 8.1.0; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.158 Mobile Safari/537.36"

	device, err = wengine.LookupUserAgent(ua)

	deviceid, err := device.GetDeviceID()

	fmt.Println(deviceid)

	fmt.Println(device.GetCapability("device_os"))
	fmt.Println(device.GetVirtualCapability("is_android"))

	device.Destroy()
	wengine.Destroy()
}
