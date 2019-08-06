const uuid = require("uuid");
const http = require("http");
const port = 3000;

const requestHandler = (request, response) => {
  let output = "";
  for (let i = 0; i < 2000000; i += 1) {
    uuid();
    output += `${i}`;
    //response.write(`${i} Your random numer is ${uuid()} \n`);
  }
  response.end(output);
  //response.end("");
};

const server = http.createServer(requestHandler);

server.listen(port, err => {
  console.log(`server is listening on ${port}`);
});


