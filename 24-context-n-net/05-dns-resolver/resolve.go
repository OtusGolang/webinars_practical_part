package main

import (
	"context"
	"log"
	"net"
)

func main() {
	resolver := net.Resolver{
		PreferGo:     true,
		StrictErrors: false,
		Dial:         nil,
	}

	ctx := context.Background()

	addrs, err := resolver.LookupIPAddr(ctx, "mail.ru")
	if err != nil {
		log.Fatalf("cannot lookup addr: %v", err)
	}

	for _, addr := range addrs {
		log.Printf("addr of mail.ru: %s", addr)
	}

	mxes, err := resolver.LookupMX(ctx, "bk.ru")
	if err != nil {
		log.Fatalf("cannot lookup mx: %v", err)
	}

	for _, mx := range mxes {
		log.Printf("mx of mail.ru: %s (pref=%d)", mx.Host, mx.Pref)
	}

	txts, err := resolver.LookupTXT(ctx, "mail.ru")
	if err != nil {
		log.Fatalf("cannot lookup txt: %v", err)
	}

	for _, txt := range txts {
		log.Printf("txt of mail.ru: %s", txt)
	}
}
