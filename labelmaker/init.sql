CREATE TABLE labels (
 video_id text NOT NULL,
 clip_id integer NOT NULL,
 label integer NOT NULL,
 UNIQUE(video_id, clip_id) ON CONFLICT REPLACE
);