# labelmaker

## Introduction

Labelmaker is a go program for labeling clips for the tennis shot classifier.  It runs a local http server that presents a web form and inline video player.  The user selects which type of shot the clip is and hits save.  That writes the label to an sqlite database and redirects to the next clip.  
![Labelmaker Image](image1.png)

## Usage

The entire repo and all its dependencies have been built into a docker image.  To get the image run the following command:

To run the labelmaker run the following command:
```bash
docker run -it -p 9090:9090 loganjspears/udacity-ml-capstone:1.0.0 bash
> cd labelmaker
> make run
```

To start labeling a video find one in ../data/videos folder and go to the following url:
```
http://localhost:9090/view/:video_id/00000
```

Labels will be saved to labels.db as they are created.  To export them to csv run the following command:
```
make csv
```