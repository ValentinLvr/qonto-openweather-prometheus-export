# Alert definitions
1. At Qonto, we want to prevent incidents, but we also want to be able to commit to our reliability levels [...] Do you have other recommendations?

We could add another alerts when: 
- our prometheus instances/pods are in a bad state (high CPU, memory, crashloop)
- when prometheus query duration is too high. for example when the average of `prometheus_engine_query_duration_seconds` is above 1s in the last 10 minutes  (bad user experience but can be lowered).
Ideally we should add a SLO on it but since it is a time-series we can't evaluate each single event. we have to evaluate the average or p99 etc
- when we get too many errors while sending alert notification `prometheus_notifications_errors` (However we could face the egg or chicken problem)
- Alerts when we get target scrape errors thanks to the `prometheus_target_*` metrics 

---
2. Our contract with our customer SLA says we shouldnʼt fail more than 1% of the Prometheus requests we receive weekly.
    1. How can we know if we are breaching?
        - We've got an SLA with a target of 99% of good events over a 7 days time window.
        - for that we will set an SLI that define a "good event" as a prometheus http request that return a 200 code.
        - We can use the `prometheus_http_requests_total` to know the number of request with a 200 code we got over 7 days and divide it by the total number of requests
        - The result of the previous step should be >=0.99 according to our SLA. otherwise we breach our SLA

    2. Can we predict if weʼre going to breach it soon?
        - we can set an error budget alert notification. Let's say, if we should burn than 60% of our error budget on less than 4 hours than we get paged.
    3. Can we have a precise date?
        - theoritically yes. there is a known relation between the error budget burn rate and the time at which it should reach 0.
        ```
        (length of SLO target (7, 30 or 90 days)) / (burn rate) = time until error budget is fully consumed
        ```

        with burn rate:
        ```
        burn rate = (length of SLO target (in hours) * percentage of error budget consumed) / (long window (in hours) *100%)
        ```
        in our example long window = 4 hours, percentage of error budget consumed = 60%, length of SLO target = 7 days
