const express = require("express");
const app = express();
const result = require("./hello.json")

app.get("/", (req, res) => res.send(result));

app.listen(5000,'127.0.0.1');
