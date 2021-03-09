To use Traefik as reverse proxy to route HTTP requests to 2 containers cd into this folder and run the command `docker-compose up -d`. <br/>
Image for the backend service from Task 1 is pulled from Docker Hub. Simple to do list manager is used as second backend service. Here is the link to it: https://github.com/prologic/todo <br/>
Sample requests:
- First backend service: GET -> http://api.localhost/api/todos
- Second backend service: GET -> http://mystery.localhost
