# graph-update-backend

A small **Bounty platform API** written in Go (chi router). It is one half of a
stakgraph cross-repo sync test fixture; its companion is
[`graph-update-frontend`](https://github.com/fayekelmith/graph-update-frontend),
whose HTTP requests should link to the endpoints defined here.

## Branches

- `main` — **initial state** (ingest target).
- `after` — **evolved state** (sync target): adds `assign` + workspace-bounties
  endpoints, renames `/people` → `/users`, and removes `DELETE /bounties/{id}`.

## Endpoints (`main`)

| Verb   | Path                     | Handler        |
| ------ | ------------------------ | -------------- |
| GET    | `/bounties`              | ListBounties   |
| POST   | `/bounties`              | CreateBounty   |
| GET    | `/bounties/{id}`         | GetBounty      |
| PUT    | `/bounties/{id}`         | UpdateBounty   |
| DELETE | `/bounties/{id}`         | DeleteBounty   |
| GET    | `/people`                | ListPeople     |
| GET    | `/people/{id}`           | GetPerson      |
| POST   | `/auth/login`            | Login          |
| GET    | `/workspaces/{id}`       | GetWorkspace   |

`GET /workspaces/{id}` intentionally has no frontend caller — a dangling
endpoint used to assert the linker does not invent edges.

## Run

```bash
go mod tidy
go run .
go test ./...
```
