# NormalizeScore Specification

`NormalizeScore(raw, max int) int`:
1. If `max <= 0`, return `0`.
2. If `raw <= 0`, return `0`.
3. If `raw >= max`, return `100`.
4. Otherwise return `(raw * 100) / max`.
