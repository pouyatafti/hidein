# hidein

simple utility to "hide" a small amount of data in an image (reads png, jpg,
and gif, and outputs png with alpha set to 255).  you can encrypt the input
yourself to have some protection, but consider it a toy and USE AT YOUR OWN
RISK as it's fairly easy to destroy and likely not too difficult to detect.

## usage

```
$ encode input.png msgi.txt >output.png

$ decode output.png [message length] >msgo.txt
```
(also see `test.sh` for an example.)
