package internal

import (
	"fmt"
	"github.com/MarkusAJacobsen/ImgurUploader"
)

func UploadImage(img string) (res *imgurUploader.ImgurResponse, err error) {
	iu := imgurUploader.GetDefaultUploader()
	iu.Config.ClientID = "e04557eb4a298af"

	req := imgurUploader.ImgurUploadBody{
		Image: img,
	}

	res, err = iu.Upload(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)

	return res, nil
}