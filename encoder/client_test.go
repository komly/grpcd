package encoder

import (
	"testing"
	"log"
	"github.com/komly/grpcd/encoder/fixtures"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"golang.org/x/net/context"
)

type impl struct {

}

func (i *impl) TestMethod(ctx context.Context, msg *fixtures.TestMessage) (*fixtures.TestMessage, error){
	return msg, nil
}

func TestNewClient(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	s := grpc.NewServer()
	fixtures.RegisterTestServiceServer(s, &impl{})
	reflection.Register(s)
	go s.Serve(ln)

	defer s.Stop()

	cl, err := New(ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cl.Invoke(nil)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Resp: %+v", resp)
}
