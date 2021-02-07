# JAMA - Jackson's Awesome Media Application
A lightweight(suckless) solution to view a library of videos.

JAMA is a lightweight http application that serves the videos (basic file server), and a REST endpoint to easily provide an interface to create your own client.

## TODO:
    1. [ ] Serve files
    2. [ ] Create endpoint for media selection (TV/Movie/Other)
        * Serve each folder, include type (directory or normal file)
    3. [ ] Create dmenu client for MPV
        * Probably will use jq and curl, so that JSON is the default endpoint as opposed to plaintext or html

## Why Golang?
Golang is nice for this kind of project because it's fast, its standard library is awesome, and static linking executable should prove nice for deployment.
