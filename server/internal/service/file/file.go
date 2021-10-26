package file

import (
	"caoguo/internal/config"
	"mime/multipart"
	"net/http"

	"github.com/xxjwxc/public/mylog"

	"github.com/xxjwxc/public/message"

	"github.com/xxjwxc/public/myfile"

	"github.com/gin-gonic/gin"
)

// Upload file deal
type Upload struct {
}

// UploadMult 上传文件 多个
// @Router /upload [post]
func (u *Upload) UploadMult(c *gin.Context) {
	op := myfile.NewWebFile("file/image/", true)
	// Multipart form
	form, _ := c.MultipartForm()
	if form == nil {
		mylog.ErrorString("/upload is empty")
		msg := message.GetErrorMsg(message.ParameterInvalid)
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	var files []*multipart.FileHeader
	for _, v := range form.File {
		files = append(files, v...)
	}

	//op := myfile.NewWebFile("file/image/", true)
	result := make([]string, 0, len(files))
	for _, file := range files {
		src, err := op.SaveOne(file)
		if err != nil {
			return
		}
		result = append(result, config.GetFileHost()+src)
	}

	msg := message.GetSuccessMsg()
	msg.Data = result

	c.JSON(http.StatusOK, msg)
}
