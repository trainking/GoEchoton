#!/usr/bin/python3

import subprocess

print("start!")

env_goarch = subprocess.check_output("go env get GOARCH")
env_goarch = env_goarch[1:-1]
env_goos = subprocess.check_output("go env get GOOS")
env_goos = env_goos[1:-1]

subprocess.run("go env -w GOARCH=amd64 GOOS=linux")
subprocess.run("go build -o ./GoEchoton ./")

subprocess.run("go env -w GOARCH=%s GOOS=%s" % (str(env_goarch, encoding="utf-8"), str(env_goos, encoding="utf-8")))

print("buid!")