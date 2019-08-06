const http = require("http");
const port = 3001;
const requestHandler = (req, res) => {
    res.end("Hello World");
}
const server = http.createServer(requestHandler);

server.listen(port, err =>{
    console.log(`server is listening on port ${port}`);
})
