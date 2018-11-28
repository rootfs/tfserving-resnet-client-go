package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	tf_core_framework "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
	pb "tensorflow_serving/apis"

	"google.golang.org/grpc"
)

func main() {
	servingAddress := flag.String("serving-address", "localhost:8500", "The tensorflow serving address")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: " + os.Args[0] + " --serving-address localhost:8500 path/to/img")
		os.Exit(1)
	}

	imgPath, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	imageBytes, err := ioutil.ReadFile(imgPath)
	if err != nil {
		log.Fatalln(err)
	}

	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          "resnet",
			SignatureName: "serving_default",
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"image_bytes": &tf_core_framework.TensorProto{
				Dtype: tf_core_framework.DataType_DT_STRING,
				TensorShape: &tf_core_framework.TensorShapeProto{
					Dim: []*tf_core_framework.TensorShapeProto_Dim{
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(1),
						},
					},
				},
				StringVal: [][]byte{imageBytes},
			},
		},
	}

	conn, err := grpc.Dial(*servingAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to the grpc server: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewPredictionServiceClient(conn)

	resp, err := client.Predict(context.Background(), request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("classes: %v", resp.Outputs["classes"].Int64Val)
}
