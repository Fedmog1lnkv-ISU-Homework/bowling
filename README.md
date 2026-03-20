[![CI](https://github.com/Fedmog1lnkv-ISU-Homework/bowling/actions/workflows/ci.yml/badge.svg)](https://github.com/Fedmog1lnkv-ISU-Homework/bowling/actions/workflows/ci.yml)
[![Coverage](https://codecov.io/github/Fedmog1lnkv-ISU-Homework/bowling/graph/badge.svg?token=SS6VHS744Y)](https://codecov.io/github/Fedmog1lnkv-ISU-Homework/bowling)
## Bowling

Симулятор боулинга на Go.

## CI/CD

- `test` job запускает тесты и собирает покрытие (`coverage.txt`).
- `publish` job публикует Docker-образ в GHCR только после успешных тестов и только при `push` в `main`.

## Docker image

Образ публикуется в:

`ghcr.io/fedmog1lnkv-isu-homework/bowling`
