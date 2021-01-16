#!/usr/bin/python3

import os
import subprocess
import shutil
import zipfile

# zipfilename是压缩包名字，dirname是要打包的目录
def compress_file(zipfilename, dirname):
    if os.path.isfile(dirname):
        with zipfile.ZipFile(zipfilename, 'w') as z:
            z.write(dirname)
    else:
        with zipfile.ZipFile(zipfilename, 'w') as z:
            for root, dirs, files in os.walk(dirname):
                for single_file in files:
                    if single_file != zipfilename:
                        filepath = os.path.join(root, single_file)
                        z.write(filepath)

def build(project, build_path, goos="linux", goarch="amd64"):
    # save old go env
    env_goarch = subprocess.check_output("go env get GOARCH")
    env_goarch = env_goarch[1:-1]
    env_goos = subprocess.check_output("go env get GOOS")
    env_goos = env_goos[1:-1]

    # change to linux env
    subprocess.run("go env -w GOARCH=%s GOOS=%s" % (goarch, goos))
    # building...
    subprocess.run("go build -o %s/%s ./" % (build_path, project))
    # copy config file: env.yaml
    shutil.copyfile('env.yaml', "%s/env.yaml" % build_path)

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

    pkg_name = project+"_amd64_linux"
    build_path = "./bin" + os.path.sep + pkg_name 
    if not os.path.exists(build_path):
        os.makedirs(build_path)
    # to build
    build(project, build_path)
    print("buid!")
    compress_file(pkg_name+".zip", build_path)
    print("zip!")
    print("end!")

if __name__ == "__main__":
    main()