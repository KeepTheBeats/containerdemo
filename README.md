# README #

This project is a demo to put an application into a container image, and pass parameters to the application.

### Application ###

1. read 2 parameters.
2. show the 2 parameters on a web.

### We know the following 3 methods to pass parameters to an application in a container ###

* from command line.
  * example: **./containerdemo -inputfrom cmd -parameter1 asdf -parameter2 zxcv**
* from environment variable.
  * example: **./containerdemo -inputfrom env**
* from config file. 
  * example: **bash run.sh "./containerdemo" "-inputfrom" "config" "-config" "para4Sh.conf"**
    1. Firstly, use the script **run.sh** to render the placeholders in the config file.
    2. Then, execute **containerdemo** to read parameters from the config file.
