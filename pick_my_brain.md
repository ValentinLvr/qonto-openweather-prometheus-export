
1. Among the tools that youʼre regularly using, which ones are significantly improving the Mean Time To Recover MTTR? How so?
- All the GitOps tools (argoCD, etc). Because it makes a rollback way easier and faster. (for example with kubernetes and argocd, you just have to change the container image on the manifest to rollback a bugged version)
- All the observability tools because it can give you insights of your overall distributed system. (especially thanks to the traces).
In my case I use Honeycomb because we can use their BubbleUp feature than can tell you which attributes is responsible of a specific pattern (e.g: slow http responses )
- Slack because communication is way faster than email with it.
- Terraform because you can re-build a cloud resource really fast with the right configuration. (also because it can be versioned thanks to git)

2. How would you explain what are Service Level Indicators SLI, Objectives SLO, and Agreements SLA? What steps should we follow to define them?
- SLI should tell you if a specific event is considered as bad or good. For example, a single span (which is an event) can be considered as good if its attribute hhtp.status_code == 200.

- SLO is a target that defines what proportion of "good" events you would like to reach on a specific time window. SLO are based on SLIs to know what is considred as a "good" events.
(e.g: we should have 99,90% of good events on a 30 days time window)

- SLA, for me, is really simillar to SLO but it's rather for "external" contract. (e.g: contract between Google Cloud Platform Database availibility and external customers)

- we should, in order, define our SLIs and put SLOs based on them. Then, if you have a contract with external customers, you should define SLAs.

3. What are the benefits/risks of using metrics/logs/traces? When would you use each technique?

- LOGS: logs can be seperated in 2 main kinds: system & application logs. 
    - System logs are useful to debug a service that has low-level errors. Can be harsh sometimes to debug if we lack underlying system knowledge.
    - Application logs was the classic way to instument and "observe" your services. However, in a micro-services architecture, it tends to be inefficient to debug with them because

- TRACES: In a distributed system, traces is more advanced than classic request logs because it embed all your services informations in one full event.
    - we should use traces rather than request logs to debug and search for bottlenecks in a micro-services architecture.

- METRICS: classic way to monitor the underlying system & architecture of your services.
    - it is really handy and you can build a lot of dashboard with alerts to have insights of your overall system
    - However, it is used to only debug the "knowns-unknowns". We set alert threshold on specific metrics base on our experiences. In case of a unknown-unknown (for example something that we didn't monitor yet)
    - Can produces a lot of headaches when you have hundred,thousands of alerts with a tone of dashboards.
    - You should still monitor important metrics even with an observability approach (CPU, MEMORY, DISK)