# Tokenizer baselines — C, Rust, Go (reference folder)

Same benchmark in three languages, for future reference: a faithful
port of SQLite's `sqlite3GetToken`, run on identical work.

## The work (identical for all three, and for Come)

- 4,000 SQL lines from `workload.sql` (copied into each folder),
  tokenized 20 times.
- Must produce exactly: 2,530,480 tokens,
  typesum 285119860, lensum 7962360.
- If a rebuild does not print those numbers, it is not the same
  benchmark — do not compare its time.

## Measured results (2026-10-24, this VM, 3 runs each, shell `date`)

- C: ~33 ms (cc -O2)
- Rust: ~80 ms (81 / 79 / 79; rustc -C opt-level=3, std HashMap)
- Go: ~131 ms (128 / 134 / 131; go build, default flags)
- Come at the time: default ~622 ms, --lto ~380 ms
  (Come has since improved a lot — see the Come logs, not this
  folder, for current Come numbers.)

VM timing is noisy, so treat these as approximate (~33 / ~80 /
~131), not exact to the millisecond. All three binaries were
re-run on 2026-10-26 and still print the exact checksums above.

## Folder layout

- `c/bench.c` — the C harness: reads the lines once, 20 passes
  calling `sqlite3GetToken`, accumulating the checksums so the
  loop cannot be optimised away. It is deliberately plain: no
  SIMD, no tricks.
- `c/tokenize.c.reference` — reference copy of SQLite's own
  `src/tokenize.c` (SQLite trunk f879fb1, from `~/workspace/
  sqlite-ref`). This is the production C code being timed,
  not a port written for the benchmark.
- `c/workload.sql` — the shared workload.
- `go/main.go`, `go/consts.go` — faithful Go port.
- `rust/main.rs`, `rust/consts.rs` — faithful Rust port.
- Each language folder also has its own `workload.sql` copy
  and its prebuilt binary (`bench_go`, `bench_rs`).

## What the C build actually links

`bench.c` only declares `sqlite3GetToken`. The implementation
comes from the SQLite amalgamation, which is NOT copied here
(8.9 MB, at `~/workspace/co_SQLite/tools/sqlite3.c`).
`tokenize.c.reference` is for reading, not for building.

## Build and run

Run each from inside its own folder, so it finds `workload.sql`:

C (needs the amalgamation):
```
gcc -O2 -DSQLITE_PRIVATE= bench.c ~/workspace/co_SQLite/tools/sqlite3.c -o bench_c -lm
./bench_c
```
Verified 2026-10-26: prints
`C: 2530480 tokens, typesum=285119860 lensum=7962360 (20 passes, 4000 lines)`.

Go:
```
~/sdk/go/bin/go build -o bench_go .
./bench_go
```

Rust (single-file build, no Cargo project):
```
~/.cargo/bin/rustc -C opt-level=3 main.rs -o bench_rs
./bench_rs
```

Toolchains (verified 2026-10-26): gcc 13.3.0, Go 1.23.0 at
`~/sdk/go`, rustc 1.99.1 at `~/.cargo/bin`.

## Provenance notes

- Rust/Go constants (186 TK_* codes, 147 keyword codes, CC_*
  classes) were generated from the Come sources by script, not
  hand-copied.
- Rust/Go started life as guesses (before the toolchains were
  installed here); the numbers above replaced those guesses
  once the ports were built and measured. Do not cite the old
  guess ranges.
- Come's own tokenizer source lives in the co_SQLite repo
  (`src/tokenizer/`), not here — this folder is only the
  C/Rust/Go reference side.
