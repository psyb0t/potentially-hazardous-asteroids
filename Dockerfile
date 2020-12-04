FROM ubuntu
COPY build/potentially-hazardous-asteroids /potentially-hazardous-asteroids
WORKDIR /
ENTRYPOINT ["/potentially-hazardous-asteroids"]
