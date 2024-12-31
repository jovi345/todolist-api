
#!/bin/bash

gcloud services enable artifactregistry.googleapis.com cloudbuild.googleapis.com run.googleapis.com

gcloud artifacts repositories create todolist-be --repository-format=docker --location=asia-southeast2 --async

gcloud builds submit --tag asia-southeast2-docker.pkg.dev/tugas-15-31122024/todolist-be/todolist-be

gcloud run deploy --image asia-southeast2-docker.pkg.dev/tugas-15-31122024/todolist-be/todolist-be