package helpers

import (
	"artics-server/config"
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"

	"github.com/disintegration/imaging"
)

var rekognitionClient *rekognition.Client = nil

func cachedInit() *rekognition.Client {
	if rekognitionClient != nil {
		return rekognitionClient
	}
	cfg, err := config.AWS()
	if err != nil {
		panic(err)
	}
	rekognitionClient = rekognition.NewFromConfig(cfg)
	return rekognitionClient
}

func convertToSupportedFormat(buffer *[]byte) ([]byte, error) {
	r := bytes.NewReader(*buffer)
	decoder, err := imaging.Decode(r)
	if err != nil {
		return nil, err
	}
	var transformedBuf bytes.Buffer
	err = imaging.Encode(&transformedBuf, decoder, imaging.JPEG, imaging.JPEGQuality(80))
	if err != nil {
		return nil, err
	}
	return transformedBuf.Bytes(), nil

}

func CheckImage(ImgBuffer *[]byte) (bool, error) {
	client := cachedInit()
	JPEGBuffer, err := convertToSupportedFormat(ImgBuffer)
	if err != nil {
		return false, err
	}
	input := &rekognition.DetectModerationLabelsInput{
		Image: &types.Image{
			Bytes: JPEGBuffer,
		},
	}
	resp, err := client.DetectModerationLabels(context.TODO(), input)

	if err != nil {

		return false, err
	}
	isPostAllowed := allowPost(resp.ModerationLabels)
	if isPostAllowed {
		return true, nil
	}

	return false, nil
}

func allowPost(labels []types.ModerationLabel) bool {
	if len(labels) == 0 {
		return true
	}
	arr := []Toy{
		{
			Confidence: 99.94709777832031,
			Name:       "Explicit Nudity",
			ParentName: "",
		},
		{
			Confidence: 99.94709777832031,
			Name:       "Adult Toys",
			ParentName: "Explicit Nudity",
		},
	}
	labelLen := len(labels)
	if len(arr) != labelLen {
		return false
	}
	var response bool = false
	for i := 0; i < len(arr); i++ {
		if &arr[i].Name == labels[i].Name {
			response = true
		} else {
			response = false
			break
		}
	}
	if response == false {
		for i := 0; i < labelLen; i++ {
			parentName := *labels[i].ParentName
			if parentName == "Explicit Nudity" || parentName == "Violence" ||
				parentName == "Visually Disturbing" ||
				parentName == "Rude Gestures" ||
				parentName == "Drugs" ||
				parentName == "Tobacco" ||
				parentName == "Alcohol" ||
				parentName == "Gambling" ||
				parentName == "Hate Symbols" {
				return false
			} else {
				return true
			}

		}
	} else {
		return true
	}
	return false
}

type Toy struct {
	Name       string
	ParentName string
	Confidence float64
}
