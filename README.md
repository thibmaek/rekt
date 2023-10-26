# Rekt

```plaintext
=====================================================
                 __  __     __
 _ __     __    /\ \/  \   /\ \__
/\` __\ / ,.`\  \ \    <   \ \ ,_\
\ \ \/ /\  __/   \ \  ^  \  \ \ \/
 \ \_\ \ \____\   \ \_\ \_\  \ \ \_
  \/_/  \/____/    \/_/\/_/   \ \__\
                               \/__/
=====================================================
```

## Install

How to get rekt:

-

## Usage

### Docker

Pull the image and bindmount a volume `scan` containing the APK:

```console
$ docker pull @thibmaek/rekt
$ docker run -it --rm -v $(pwd)/scan:/scan @thibmaek/rekt <apk>
```

### CLI

A typical run of rekt using the cli involves running:

1. Decompile - Getting plain readable files
2. Probe - Gathering info about the decompiled app
3. Break - Finding secrets and credential files

Given an APK `com.my_app.apk` you'd get the results like this:

```console
$ rekt decompile -apk=./com.my_app.apk
$ rekt probe -outputDir=./scan/com_my_app
$ rekt break -outputDir-./scan/com_my_app
```

#### Decompile

```console
$ rekt decompile -apk=./com.my_app.apk
```

Optionally provide an output directory `-outputdir`. Defaults to `./scan/<bundle_id>`

```console
$ rekt decompile -apk=./com.my_app.apk -outputDir=./decompiled_app
```

#### Probe

```console
$ rekt probe -inputDir=./scan/com_my_app
```

#### Break

```console
$ rekt break -inputDir=./scan/com_my_app
```

## Building

```console
# Build Docker & Go
$ make build

# Build only the docker image
$ make build_docker

# Build only the CLI
$ make build_cli
```

## Todos

- Support for iOS IPA archives
- Gitlab CI support
- Github Actions support
- Azure Devops support
