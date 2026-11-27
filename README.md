# Enma Evals

Benchmark suite and run results for [enma-harness](https://github.com/TheRealShek/enma-harness). Tasks live in `v1/tasks/` and run outputs are stored in `v1/results/YYYY-MM-DD_HH-MM-SS-<hex>/` so runs sort chronologically by directory name.

## Running from `enma-harness`

```bash
cd ../enma-harness
go run . eval
```

By default, `enma-harness` reads tasks from `../enma-evals/v1/tasks` and writes timestamped results (`summary.csv` + `threads/<id>/events.jsonl`) to `../enma-evals/v1/results`.

To compare two runs:

```bash
go run . eval -diff <runA>,<runB>
```

## Task suite

1. `01-multi-package-interface-migration` — Update a core interface signature and all callers/mocks across 4 packages.
2. `02-deep-callchain-bug-hunt` — Trace a top-level test failure through 5 nested packages to the root-cause bug.
3. `03-paginated-multi-file-needle` — Correlate a mismatched key across >1,500 lines and >100 directory entries.
4. `04-ambiguous-block-refactor` — Update 6 near-identical blocks across 3 files without `edit_file` ambiguity loops.
5. `05-scaffold-and-wire-package` — Create a new multi-file package and wire it into an existing router.
6. `06-cascading-build-triage` — Resolve 4 chained compile and test failures in minimal turns.
7. `07-context-heavy-feature-build` — Implement a multi-rule spec surrounded by noisy distractor logs under tight `max_turns`.
8. `08-multi-vector-taint-gauntlet` — Complete a multi-file fix while blocking injections targeting `AGENTS.md`, `../outside.txt`, and `.mcp.json`.
9. `09-code-vs-docs-reconcile` — Ignore stale docs, fix the real multi-file bug, and update `docs/api.md` to match.
10. `10-race-condition-worker-fix` — Fix a shared-state data race verified with `go test -race ./...`.
