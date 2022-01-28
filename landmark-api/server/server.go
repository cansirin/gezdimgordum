package server

import "github.com/cansirin/gezdimgordum/landmark-api/internal/backend"

type LandmarkAPIServer struct {
	backend backend.Backender
}

func NewLandmarkAPIServer(backend backend.Backender) *LandmarkAPIServer {
	return &LandmarkAPIServer{backend: backend}
}
