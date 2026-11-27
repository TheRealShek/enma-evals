# v1 Eval Suite

Fixed 10-task benchmark (`v1/tasks/`) used to measure [enma-harness](https://github.com/TheRealShek/enma-harness) changes on a single model (`qwen35-4b-32k`). The task set stays constant across runs while the harness evolves.

## Results (`v1/results/`)

| Run | Pass rate | Notes |
| :--- | :--- | :--- |
| `2026-11-27_07-07-50-43a4d1` | `12/30` (`40%`) | Initial baseline on `qwen35-4b-32k` (`01`, `05`, `10` at `0/3`). |
| `2026-11-27_08-38-44-5e24c2` | `13/30` (`43%`) | Fixed Task 10 `go.mod` (`go 1.21` -> `go 1.22`, moving Task 10 `0/3` -> `2/3`) and switched `summary.json` to `summary.csv`. |

## Next (Run 3)

Add a `run_command` tool to [enma-harness](https://github.com/TheRealShek/enma-harness) so the model can run `go test` directly instead of burning turns guessing or re-reading files.
