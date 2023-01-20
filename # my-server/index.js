const express = require('express');
const app = express();
const PORT = 8080;

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

app.get('/get', (req, res) => {
    res.send("Hello from server!")
});

app.post('/post', (req, res) => {
    res.status(200).json(req.body);
});

app.post('/postform', (req, res) => {
    res.status(200).json(JSON.stringify(req.body));
})

app.listen(PORT, () => console.log("http://localhost:" + PORT));