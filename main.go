
package main

import (
  "fmt"
  "github.com/vishvananda/netlink"
)

func main() {

  netlink_done := make(chan struct{})
  defer close(netlink_done)

  addr_ch := make(chan netlink.AddrUpdate, 10)
  link_ch := make(chan netlink.LinkUpdate, 10)
  route_ch := make(chan netlink.RouteUpdate, 10)

  if err := netlink.AddrSubscribe(addr_ch, netlink_done); err != nil { return }
  if err := netlink.LinkSubscribe(link_ch, netlink_done); err != nil { return }
  if err := netlink.RouteSubscribe(route_ch, netlink_done); err != nil { return }

  exitCh := make(chan struct{})

  go func() {
    defer func() { close(exitCh) }()

    for {
      select {
      case msg := <-addr_ch:
        fmt.Println("addr channel")
        la := msg.LinkAddress
        ip := la.IP
        fmt.Printf("Address: %d.%d.%d.%d\n", ip[0], ip[1], ip[2], ip[3])
      case _ = <-link_ch:
        fmt.Println("link channel")
      case _ = <-route_ch:
        fmt.Println("route channel")
      }
    }

  }()

  <-exitCh
}

