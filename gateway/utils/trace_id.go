package utils

import "github.com/google/uuid"

// 生成 TraceID
func GenerateTraceID() string {
	return uuid.New().String()
}
