# Demo project for stackdriver logging issue when running on Google Cloud Run

To build and deploy, run:
```
(export PROJECT_ID=[project ID] SERVICE_ID=[service ID] ;
gcloud builds submit --tag gcr.io/$PROJECT_ID/$SERVICE_ID &&
gcloud beta run deploy $SERVICE_ID --image gcr.io/$PROJECT_ID/$SERVICE_ID --platform managed)
```
Logs area available in google cloud console immediately after service start
