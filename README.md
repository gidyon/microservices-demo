# microservices-demo
A very simple Go-Redis app to demo discovery of multiple services.

### Pre-requisites
1. Docker version > 1.8, docker-compose version > 1.3.1, Swarm version > 0.4

### Steps
1. Clone this repo to a local folder. `git clone https://github.com/gidyon/microservices-demo.git`
2. `cd microservices-demo`
3. Use docker-compose to build and run the web app containers. `docker-compose up -d`
4. Browse to the URL: `http://localhost:8080/demo`
  - <img src="https://farm1.staticflickr.com/666/21705956952_9b3bfea89f_b.jpg" width=300>
5. Every time a container responds to the HTTP request, it will get its counter incremented (on a browser refresh). The counter is being stored (and retrieved) from a REDIS backend database.
6. Using docker-compose, you can scale up or down the web containers as you wish based on the needs and traffic to your application.
