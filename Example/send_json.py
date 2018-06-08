#!/usr/bin/env python3
# coding:utf-8
from flask import Flask, jsonify
app = Flask(__name__)


@app.route("/")
def hello():

    result = {
        "Result": "Hello World"
    }

    return jsonify(result)


if __name__ == "__main__":
    app.run(debug=True)