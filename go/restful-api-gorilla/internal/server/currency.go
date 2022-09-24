package server

import (
	"context"
	"io"
	"time"

	"github.com/hashicorp/go-hclog"
	"gitlab.com/prasetiyohadi/go-usvc/internal/data"
	proto "gitlab.com/prasetiyohadi/go-usvc/internal/server/proto/currency"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Currency is a gRPC server
type Currency struct {
	rates         *data.ExchangeRates
	log           hclog.Logger
	subscriptions map[proto.Currency_SubscribeRatesServer][]*proto.RateRequest
	proto.UnimplementedCurrencyServer
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	c := &Currency{r, l, make(map[proto.Currency_SubscribeRatesServer][]*proto.RateRequest), proto.UnimplementedCurrencyServer{}}
	go c.handleUpdates()

	return c
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)
	for range ru {
		c.log.Info("Got updated rates")

		// loop over subscribed clients
		for k, v := range c.subscriptions {

			// loop over subscribed rates
			for _, rr := range v {
				r, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.log.Error("Unable to get updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}

				// create the response and send to the client
				err = k.Send(
					&proto.StreamingRateResponse{
						Message: &proto.StreamingRateResponse_RateResponse{
							RateResponse: &proto.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: r},
						},
					},
				)
				if err != nil {
					c.log.Error("Unable to send updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}
			}
		}
	}
}

// GetRate implements interface of CurrencyServer
func (c *Currency) GetRate(ctx context.Context, rr *proto.RateRequest) (*proto.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())

	// Validate parameters base currency can not be the same as destination
	if rr.Base == rr.Destination {
		// create the grpc error and return to the client
		err := status.Errorf(
			codes.InvalidArgument,
			"Base rate %s can not be equal to destination rate %s",
			rr.Base.String(),
			rr.Destination.String(),
		)

		return nil, err
	}

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &proto.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate}, nil
}

// SubscribeRates implements the gRPC bi-direction streaming method for the server
func (c *Currency) SubscribeRates(src proto.Currency_SubscribeRatesServer) error {

	// handle client messages
	for {
		rr, err := src.Recv() // Recv is a blocking method which retunrs on client data
		// io.EOF signals that the client has closed the connection
		if err == io.EOF {
			c.log.Info("Client has closed the connection")
			break
		}

		// any other error means the transport between the server and the client is unavailable
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			return err
		}

		c.log.Info("Handle client request", "request_base", rr.GetBase(), "request_dest", rr.GetDestination())

		// get the current subscriptions for this client
		rrs, ok := c.subscriptions[src]
		if !ok {
			rrs = []*proto.RateRequest{}
		}

		// check if already in the subscribe list and return a custom gRPC error
		for _, r := range rrs {
			// if we already have subscribed to this currency return an error
			if r.Base == rr.Base && r.Destination == rr.Destination {
				c.log.Error("Subscription already active", "base", rr.Base.String(), "dest", rr.Destination.String())

				grpcError := status.New(codes.InvalidArgument, "Subscription already active for rate")
				grpcError, err = grpcError.WithDetails(rr)
				if err != nil {
					c.log.Error("Unable to add metadata to error message", "error", err)
					continue
				}

				// Cannot return error as that will terminate the connection, instead must send an error which
				// can be handled by the client Recv stream.
				rrs := &proto.StreamingRateResponse_Error{Error: grpcError.Proto()}
				src.Send(&proto.StreamingRateResponse{Message: rrs})
			}
		}

		// if all is ok then add to the collection
		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}

	return nil
}
