package dbtest

import (
	"context"
	"net"
)

//WaitForTCP blocks until a connection can be made via TCP
func WaitForTCP(ctx context.Context, rAddr string) error {
	dialer := net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", rAddr)
	//For loop to get around OS Dial Timeout
	for err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		conn, err = dialer.DialContext(ctx, "tcp", rAddr)
	}
	conn.Close()
	return nil
}
