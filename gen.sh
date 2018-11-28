# git clone https://github.com/tensorflow/serving.git
# git clone https://github.com/tensorflow/tensorflow.git

 mkdir -p vendor
 PROTOC_OPTS='-I tensorflow -I serving --go_out=plugins=grpc:vendor'
 eval "protoc $PROTOC_OPTS serving/tensorflow_serving/apis/{predict,prediction_service,model,regression,classification,inference,input,get_model_metadata}.proto"
 eval "protoc $PROTOC_OPTS serving/tensorflow_serving/config/*.proto"
 eval "protoc $PROTOC_OPTS serving/tensorflow_serving/util/*.proto"
 eval "protoc $PROTOC_OPTS serving/tensorflow_serving/sources/storage_path/*.proto"
 eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/framework/*.proto"
 eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/example/*.proto"
 eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/lib/core/*.proto"
 eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/{saver,saved_model,meta_graph}.proto"
