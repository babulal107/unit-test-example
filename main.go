package main

import (
	speech "cloud.google.com/go/speech/apiv1"
	"context"
	"fmt"
	"github.com/unit-test-example/employee"
	"google.golang.org/api/option"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"log"
	"os"
)

func main() {
	emp := employee.Employee{}
	data, err := emp.GetData()
	if err != nil {
		fmt.Println("Error GetData : ", err)
	}
	fmt.Println("Data : ", data)

	//PostTranscribeRequest("gs://idstats-audio/pere_babulal_transcoder/1234_Rm9Q7_213123213STE2323_2IoOs_2314324324324SDFDsf_AFKkc_sample_video_2.mp3")

	//GetTranscribePullData("1833849626958597602")
}

func PostTranscribeRequest(gCloudAudioUrl string) {

	var (
		audioLanguageCode = "en-IN"
	)
	ctx := context.Background()

	// Get current workspace directory
	currentDir, _ := os.Getwd()

	// Creates a client.
	client, err := speech.NewClient(ctx, option.WithCredentialsFile(currentDir+"/configs/google-cloud-idstats-vision-api-key.json"))
	if err != nil {
		log.Println("Failed to create client: ", err)
		return
	}

	log.Println("=======LongRunningRecognizeRequest======>")

	req := &speechpb.LongRunningRecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:                   speechpb.RecognitionConfig_ENCODING_UNSPECIFIED,
			SampleRateHertz:            16000,
			LanguageCode:               audioLanguageCode,
			EnableAutomaticPunctuation: true,
			DiarizationConfig:          &speechpb.SpeakerDiarizationConfig{EnableSpeakerDiarization: true},
			EnableWordTimeOffsets:      true,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: gCloudAudioUrl},
		},
	}

	op, err := client.LongRunningRecognize(ctx, req)
	if err != nil {
		log.Println("LongRunningRecognizeRequest : ", err)
		return
	}

	log.Println("LongRunningRecognize request operation name : ", op.Name())
}

func GetTranscribePullData(operationName string) {

	ctx := context.Background()

	// Get current workspace directory
	currentDir, _ := os.Getwd()

	// Creates a client.
	log.Println("=======LongRunningRecognizeOperation Pulling Request======>")
	client, err := speech.NewClient(ctx, option.WithCredentialsFile(currentDir+"/configs/google-cloud-idstats-vision-api-key.json"))
	if err != nil {
		log.Println("Failed to create client: ", err)
		return
	}
	op := client.LongRunningRecognizeOperation(operationName)
	resp, err := op.Wait(ctx)
	if err != nil {
		log.Println("LongRunningRecognizeOperation Wait :", err.Error())
		return
	}
	log.Println("=======LongRunningRecognizeOperation  Response======>")
	fmt.Println("response: ", resp)
	for _, result := range resp.Results {
		for index, alt := range result.Alternatives {
			fmt.Println("alternatives : index : ", index, alt)
		}
	}
}
