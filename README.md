## Environment

A VPS at Tencent Cloud.

- 1 core of AMD EPYC 7K62 48-Core Processor

- 1GB RAM

- Ubuntu 18.04

- Linux VM-0-125-ubuntu 4.15.0-88-generic #88-Ubuntu SMP Tue Feb 11 20:11:34 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux

- go version go1.14.7 linux/amd64

## Code

See `main.go` file.

## Result that I expected

The program runs forever, without an out-of-memory error.

## Result that actually happened

```
tick-tock 47854 0xc000010010
fatal error: runtime: out of memory

runtime stack:
runtime.throw(0x4e8b5e, 0x16)
        /usr/lib/go/src/runtime/panic.go:1116 +0x72
runtime.sysMap(0xc028000000, 0x14000000, 0x5cd258)
        /usr/lib/go/src/runtime/mem_linux.go:169 +0xc5
runtime.(*mheap).sysAlloc(0x5b99a0, 0x13800000, 0x5b99a8, 0x9a3c)
        /usr/lib/go/src/runtime/malloc.go:715 +0x1cd
runtime.(*mheap).grow(0x5b99a0, 0x9a3c, 0x0)
        /usr/lib/go/src/runtime/mheap.go:1286 +0x11c
runtime.(*mheap).allocSpan(0x5b99a0, 0x9a3c, 0x420100, 0x5cd268, 0x200000000)
        /usr/lib/go/src/runtime/mheap.go:1124 +0x6a0
runtime.(*mheap).alloc.func1()
        /usr/lib/go/src/runtime/mheap.go:871 +0x64
runtime.(*mheap).alloc(0x5b99a0, 0x9a3c, 0xc000000001, 0x5a27a0)
        /usr/lib/go/src/runtime/mheap.go:865 +0x81
runtime.largeAlloc(0x13478000, 0x5a0100, 0xc000031f01)
        /usr/lib/go/src/runtime/malloc.go:1152 +0x92
runtime.mallocgc.func1()
        /usr/lib/go/src/runtime/malloc.go:1047 +0x46
runtime.systemstack(0x45a8b4)
        /usr/lib/go/src/runtime/asm_amd64.s:370 +0x66
runtime.mstart()
        /usr/lib/go/src/runtime/proc.go:1041

goroutine 1 [running]:
runtime.systemstack_switch()
        /usr/lib/go/src/runtime/asm_amd64.s:330 fp=0xc000064d58 sp=0xc000064d50 pc=0x45a9b0
runtime.mallocgc(0x13478000, 0x0, 0xc000010200, 0xc0000102e0)
        /usr/lib/go/src/runtime/malloc.go:1046 +0x895 fp=0xc000064df8 sp=0xc000064d58 pc=0x40bad5
runtime.growslice(0x4c2140, 0xc000088000, 0xf6c6000, 0xf6c6000, 0xf6c6001, 0xc000088000, 0xc56a000, 0xf6c6000)
        /usr/lib/go/src/runtime/slice.go:175 +0x14e fp=0xc000064e60 sp=0xc000064df8 pc=0x4447ae
github.com/tendermint/tendermint/libs/rand.(*Rand).Str(0xc000010330, 0x0, 0x0, 0x10)
        /root/go/pkg/mod/github.com/tendermint/tendermint@v0.33.7/libs/rand/random.go:161 +0xf7 fp=0xc000064ee8 sp=0xc000064e60 pc=0x4b1c97
github.com/tendermint/tendermint/libs/rand.Str(...)
        /root/go/pkg/mod/github.com/tendermint/tendermint@v0.33.7/libs/rand/random.go:61
main.f1()
        /mnt/d/Users/Documents/GitHub/a-strange-thing-of-tendermint-rand/main.go:13 +0x12f fp=0xc000064f78 sp=0xc000064ee8 pc=0x4b20df
main.main()
        /mnt/d/Users/Documents/GitHub/a-strange-thing-of-tendermint-rand/main.go:48 +0x20 fp=0xc000064f88 sp=0xc000064f78 pc=0x4b2140
runtime.main()
        /usr/lib/go/src/runtime/proc.go:203 +0x1fa fp=0xc000064fe0 sp=0xc000064f88 pc=0x43200a
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1373 +0x1 fp=0xc000064fe8 sp=0xc000064fe0 pc=0x45c941
```

