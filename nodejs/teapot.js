var app = require('express')()
  , server = require('http').createServer(app)
  , io = require('socket.io').listen(server);

server.listen(80);

app.get('/', function (req, res) {
  res.sendfile(__dirname + '/index.html');
});

io.sockets.on('connection', function (socket) {

  socket.on('annotateCollation', function (data) {
  	socket.emit('collationAnnotated', { name: data.name, value: data.value });
    console.log(data);
  });

  socket.on('annotateMedia', function (data) {
    socket.emit('mediaAnnotated', { name: data.name, value: data.value });
    console.log(data);
  });
});