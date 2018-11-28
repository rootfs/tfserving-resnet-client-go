# Tensorflow Serving ResNet Go Client

(test against tensorflow `b4add48e633ce9e1ecda25eade56638e0a8808e5` and tensorflow-serving `5446fd973de228693c1652acd4922dc4b177f77a`)

## Generate vendor

```bash
git clone https://github.com/tensorflow/serving.git
git clone https://github.com/tensorflow/tensorflow.git

./gen.sh
```
### Build client

```bash
go build resnet.go
```

### Test

- Start Resnet model server
```console
 docker run -p 8500:8500 -p 8501:8501 -t docker.io/rootfs/resnet_serving &
```
- Test against resnet model

```console
# ./resnet --serving-address localhost:8500 ./cat.jpg
2018/11/28 16:08:32 classes: [286]
```
# Acknowledgement

Inspired by this [inception client](https://mauri870.github.io/blog/posts/tensorflow-serving-inception-go/)
