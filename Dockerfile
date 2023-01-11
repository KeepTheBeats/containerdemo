FROM ubuntu

RUN apt update -y
RUN apt upgrade -y

RUN apt install -y git
# RUN git clone https://github.com/KeepTheBeats/containerdemo.git

# copy all files and folders in the current path to the /containerdemo in the container
COPY . /containerdemo

RUN chmod +x /containerdemo/containerdemo
RUN chmod +x /containerdemo/run.sh

ENTRYPOINT ["bash", "/containerdemo/run.sh"]
CMD ["./containerdemo", "-inputfrom", "cmd", "-parameter1", "incmdadfsf", "-parameter2", "incmdasdff"]