# udacity-ml-capstone

## Introduction

For my Udacity Machine Learning Capstone project I chose to implement a tennis shot classifier.  The input is a one second video clip and the output is the type of shot. 

![Example](example.jpg)

#### Prediction: Volley

## Installation

The entire repo and all its dependencies have been built into a docker image.  To get the image run the following command:

```bash
docker pull loganjspears/udacity-ml-capstone:1.0.0
```  

To run the jupyter notebook run the following command:
```bash
docker run -it -p 8888:8888 loganjspears/udacity-ml-capstone:1.0.0
```

## Usage

The docker run command should output a URL that you can paste in your browser.  It will look something like this:  
```
Copy/paste this URL into your browser when you connect for the first time,
    to login with a token:
        http://localhost:8888/?token=3d720d16f8c6cae9d920d6558ea438681170cf847dd52da3
```

After loading the directory view select "project.ipynb" to view the Jupyter notebook for this project.  The docker image contains preprocessed data not available in the git repo.  If you would like to rerun the preprocessing steps (they take a LONG TIME), then run the docker container with the bash command and delete the directors under audio, clips, features, and frames.

To use the labelmaker view its [README](labelmaker/README.md)

## Technologies Used

- [Python 2.x](https://www.python.org)
- [Tensorflow](https://www.tensorflow.org)
- [Keras](https://keras.io)
- [Docker](https://www.docker.com)
- [ffmpeg](https://ffmpeg.org)
- [ImageMagick](http://www.imagemagick.org/script/index.php)
- [Go](https://golang.org)
- [Glide](https://glide.sh)
