package common

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	PDerr = status.Error(codes.PermissionDenied, "Доступ запрещён")
	NFerr = status.Error(codes.NotFound, "Данные не найденны")
	INerr = status.Error(codes.Internal, "Внутренняя ошибка сервера")
	DLerr = status.Error(codes.DataLoss, "Ошибка запроса к БД")
)
