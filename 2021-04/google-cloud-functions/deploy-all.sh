go build || exit 1
for i in \
    HelloWorld \
    LazyInit \
    HTML \
    ExperimentBlockRequests \
    ExperimentParallelRequests \
    ExperimentGoroutines
do
    gcloud functions deploy $i \
        --runtime go113 \
        --allow-unauthenticated \
        --trigger-http
done
