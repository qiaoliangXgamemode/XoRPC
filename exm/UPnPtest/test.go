package main

import (
	"fmt"
	"log"

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/internetgateway1"
)

func main() {
	// 创建一个UPnP发现器
	discoverer, err := goupnp.DiscoverDevices(internetgateway1.URN_WANPPPConnection_1)
	if err != nil {
		log.Fatal("Failed to discover UPnP devices:", err)
	}

	// 遍历设备列表
	for _, device := range discoverer {
		fmt.Println("Device:", device.Root.Device.FriendlyName)

		// 获取设备的服务列表
		services, err := device.Root.Device.Services()
		if err != nil {
			log.Fatal("Failed to get services for device:", err)
		}

		// 遍历服务列表
		for _, service := range services {
			fmt.Println("Service:", service.ServiceType)
		}
	}
}
