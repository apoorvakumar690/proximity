package middleware

import (
	"fmt"
	log "proximity/pkg/utils/logger"

	"github.com/gin-gonic/gin"
	pkgErrors "github.com/pkg/errors"
)

// HandlePanic ... rest panic handler
func HandlePanic(c *gin.Context) {
	defer func(c *gin.Context) {
		r := recover()
		var stackTrace string
		if r != nil {
			err, ok := r.(error)
			if ok {
				// Logs the error
				stackTrace = fmt.Sprintf("%+v", pkgErrors.New(err.Error()))
				log.Error("GO-BOILERPLATE.PANIC", "Unexpected panic occured", log.Priority1, nil, map[string]interface{}{"error": err.Error(), "stackTrace": stackTrace})

				// Notice error in apm
				// apm.APM.NoticeError(apm.FromContext(c), err)

				// Forms error message
				c.JSON(500, gin.H{
					"message":    "Panic: Unexpected error occured.",
					"error":      err.Error(),
					"stackTrace": stackTrace,
				})
			} else {
				// Logs the error
				log.Error("GO-BOILERPLATE.PANIC", "Panic recovery failed to parse error", log.Priority1, nil, map[string]interface{}{"error": r})

				// Notice error in apm
				// apm.APM.NoticeError(apm.FromContext(c), errors.New("GO-BOILERPLATE.UNRECOVERED.PANIC"))

				// Forms error message
				c.JSON(500, gin.H{
					"message": "Panic: Unexpected error occured, failed to parse error",
				})
			}
		}
	}(c)
	c.Next()
}
