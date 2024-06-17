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

## Prereqs

If running outside of Docker (e.g locally) you will need the following installed:
- jadx (Java)
- hermes-desc (Python)

## Install

How to get rekt:

- Use the binary in the `./rekt-cli/bin` folder
- Download from Github releases
- Build it yourself

## Usage

Rekt decompiles, analyzes and breaks application archives. A valid app archive file is needed for either of these platforms:

- iOS: iOS Package App Store (IPA, `.ipa`)
- Android: Android Package Kit (APK, `.apk`)

Both of these archives are essentialy glorified ZIP archives that you can also unpack manually.

### Docker

Pull the image and bindmount a volume `scan` containing your archive:

```console
$ docker pull @thibmaek/rekt
$ docker run -it --rm -v $(pwd)/scan:/scan @thibmaek/rekt <archive_file>
```

### CLI

A typical run of rekt using the cli involves running:

1. Decompile - Getting plain readable files
2. Probe - Gathering info about the decompiled app
3. Break - Finding secrets and credential files

Given an APK `com.my_app.apk` you'd get the results like this:

```console
$ rekt decompile -archive=./com.my_app.apk
$ rekt probe -outputDir=./scan/com_my_app
$ rekt break -outputDir-./scan/com_my_app
```

#### Decompile

```shell
# Decompiling an APK
$ rekt decompile -archive=./com.my_app.apk

# Decompiling an IPA
$ rekt decompile -archive=./com.my_app.ipa
```

Optionally provide an output directory `-outputdir`. Defaults to `./scan/<bundle_id>`

```console
$ rekt decompile -archive=./com.my_app.apk -outputDir=./decompiled_app
```

#### Probe

> [!WARNING]
> Probing is currently not supported on iOS

```console
$ rekt probe -inputDir=./scan/com_my_app
```

#### Break

> [!WARNING]
> Breaking is currently not supported on iOS

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
