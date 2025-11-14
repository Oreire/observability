# observability
Using Custom Built Exported for Monitoring and Observability

## 📊 FBref Premier League Prometheus Exporter

This project scrapes Premier League statistics from [FBref](https://fbref.com/en/comps/9/Premier-League-Stats), exposes them as Prometheus metrics, and visualizes them in Grafana. It is modular, reproducible, and secure-by-design — ideal for CPD simulation reports, recruiter-facing showcases, and sector-grade observability.

---

### 🧱 Project Structure

```
fbref_exporter/
├── .env
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── prometheus/
│   └── prometheus.yml
├── grafana/
│   ├── dashboards/
│   │   └── fbref_dashboard.json
│   └── provisioning/
│       ├── dashboards.yml
│       └── datasources.yml
```

---

### 🚀 How to Run Locally

```bash
# Navigate to project root
cd fbref_exporter

# Build and launch all services
docker-compose up --build
```

- Prometheus: [http://localhost:9090](http://localhost:9090)
- Grafana: [http://localhost:3000](http://localhost:3000)
  - Login: `admin / admin` (stored in `.env`)


### 📈 Metrics Exposed

| Metric Name                             | Description                                 |
|----------------------------------------|---------------------------------------------|
| `fbref_team_goals_total`               | Total goals scored by each team             |
| `fbref_team_possession_percent`        | Possession percentage by team               |
| `fbref_team_shots_total`               | Total shots taken                           |
| `fbref_team_shots_on_target`           | Shots on target                             |
| `fbref_team_pass_completion_percent`   | Pass completion percentage                  |
| `fbref_team_expected_goals`            | Expected goals (xG)                         |

Metrics are scraped every 15 minutes and exposed at `/metrics` on port `9101`.

### 📊 Grafana Dashboard

Auto-provisioned via:
- `grafana/provisioning/dashboards.yml`
- `grafana/provisioning/datasources.yml`

Includes:
- Bar charts, pie charts, gauges, and stat panels
- Real-time visualization of team performance
- Exporter health and scrape duration

### 🔐 Secrets Management

Grafana credentials are stored in `.env`:

Ensure `.env` is excluded from version control:
```
# .gitignore
.env
```

---

### 🎓 CPD Simulation & SFIA Mapping

| Competency | SFIA Code | Evidence |
|------------|-----------|----------|
| Software Development | PROG | Go-based exporter with Prometheus integration |
| Systems Integration | INAN | Docker Compose orchestration |
| Measurement | METL | Metric collection and dashboard visualization |
| Information Management | IRMG | Structured metric naming and scrape logic |
| Teaching & Simulation | TEAC | CPD-ready, reproducible deployment |


#  Local Deployment

From my `fbref_exporter` folder; I built and ran my Docker Compose setup:

```bash
docker-compose up --build
```

### 🧱 What This Does

- `--build` forces Docker to rebuild the `fbref_exporter` image from my local `Dockerfile`
- Starts all services: `fbref_exporter`, `prometheus`, and `grafana`
- Mounts volumes and loads environment variables from `.env`

---

### 🧪 Optional Commands

- **Run in background (detached mode):**
  ```bash
  docker-compose up --build -d
  ```

- **Stop all services:**
  ```bash
  docker-compose down
  ```

- **Rebuild only the exporter:**
  ```bash
  docker-compose build fbref_exporter
  ```

# CleanUP

**Make Executable:** chmod +x reset.sh
**run:** ./reset.sh

Notes
docker compose up -d --build
