#!/usr/bin/python3

import os
import subprocess
import shutil

def build(project, goos="linux", goarch="amd64"):
    # save old go env
    env_goarch = subprocess.check_output("go env get GOARCH")
    env_goarch = env_goarch[1:-1]
    env_goos = subprocess.check_output("go env get GOOS")
    env_goos = env_goos[1:-1]

    # change to linux env
    subprocess.run("go env -w GOARCH=%s GOOS=%s" % (goarch, goos))
    # building...
    subprocess.run("go build -o ./bin/%s ./" % project)
    # copy config file: env.yaml
    shutil.copyfile('env.yaml', './bin/env.yaml')

    # change back old env
    subprocess.run("go env -w GOARCH=%s GOOS=%s" % (str(env_goarch, encoding="utf-8"), str(env_goos, encoding="utf-8")))

def main():
    print("start!")
    # check ./bin is exists
    if not os.path.exists("./bin"):
        os.mkdir("./bin")

    # find module name
    pwd = subprocess.check_output("go mod why")
    pwd = str(pwd, encoding="utf-8")
    i = pwd.find('\n')
    project = pwd[i+1:-1]

    # to build
    build(project)

    print("buid!")

if __name__ == "__main__":
    main()