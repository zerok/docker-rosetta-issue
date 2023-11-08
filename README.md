# Rosetta issue in Docker For Mac

This is intended as minimal example of something that causes a segmentation fault on macOS with an Apple Silicon CPU when cross-compiling to linux/amd64:

```
go run .

...
  exec go build ERROR: process "/dev/.buildkit_qemu_emulator go build" did not complete successfully: exit code: 1
  input:1: container.from.withMountedDirectory.withWorkdir.withExec.sync process "/dev/.buildkit_qemu_emulator go build" did not complete successfully: exit code: 1
...
```

The issue is tracked on <https://github.com/docker/for-mac/issues/6773>.
