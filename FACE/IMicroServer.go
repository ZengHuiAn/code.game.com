package FACE

import "github.com/micro/go-micro/v2"

type IMicroServer interface {
	GetServerName() string
	GetServerAddr() string

	RunMicroServer()
	GetService() micro.Service
	InitMicroServer()
}