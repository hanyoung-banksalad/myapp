package handler

import (
	"os"
	"io"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	myapp "github.com/hanyoung-banksalad/myapp/idl"
)

type imageServer struct{}

var _ myapp.ImageServer = (*imageServer)(nil)

func (s *imageServer) GetImage(req *myapp.GetImageRequest, stream myapp.Image_GetImageServer) error {
	f, err := os.Open("images/" + req.Path)
	if err != nil {
		return status.Error(codes.NotFound, "file not found")
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.WithError(err).Error("failed to close file")
		}
	}()

	// Maximum 16KB size per stream.
	buf := make([]byte, 16*2<<10)
	for {
		num, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := stream.Send(&myapp.GetImageResponse{Data: buf[:num]}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
