# udacity-ml-capstone

## Introduction

For my Udacity Machine Learning Capstone project I chose to implement a tennis shot classifier.  The input is a one second video clip and the output is the type of shot. 

![Example Gif](data/frames/lzJt-w3OYKg/00353/00002.jpg)

#### Prediction: Volley

## Usage

The entire repo and all its dependencies have been built into a docker image.  To get the image run the following command:

```bash
docker pull loganjspears/udacity-ml-capstone:1.0.0
```  

To run the jupyter notebook run the following command:
```bash
docker run -it -p 8888:8888 loganjspears/udacity-ml-capstone:1.0.0
```

Go to the outputted URL and select project.ipynb from the file viewer.
