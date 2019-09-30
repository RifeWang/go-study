# NODE.JS MONITORING, ALERTING & RELIABILITY 101
# node.js 监控 & 报警 & 可靠性


## 1、理解、测量可靠性

### 明确可靠性指标
要确定系统的可靠性，首先要做的事情是定义我们的服务质量要求，以及用户可接受的程度。确定需求后，我们必须衡量系统满足这些要求的程度。

基于这些测量结果，你可以确定一个最低的质量标准。当然这些与业务保持一致。在确定之后，它们便可以推动你的产品开发并为用户提供合同的基础。

可靠性指标有三种。让我们看看它们的含义以及它们的作用：
##### Service Level Indicator ( SLI )
服务水平指标（SLI）是对所提供服务水平的某些方面的精心定义的量化指标。
系统中的服务水平指标可以是应用正常运行时间的百分比或 99th latency（ 99% 的请求应该比指定延时快 ）。

##### Service Level Objective ( SLO )
Service Level Objective 是 SLI 的目标。它定义了什么是可接受的，什么是不可接受的。
例如，Service Level Objective 可以是成功探测系统的频率，也可以是 99th latency 低于 250ms 的请求百分比。使用 SLO ，你可以计算服务的可以率或者衡量服务是否存在重大问题。
基于 SLO ，利益相关者可以决定在可靠性改进方面投入更多精力，甚至 hold back releases 阻止发布。
更严格的 SLO 要求总是需要更高的开发，维护和资源分配成本。例如，要达到更高的标准，你需要实现高级架构可靠性模式并提供更多资源以改善系统冗余。

##### Service Level Agreement ( SLA )
服务水平协议（SLA）是服务提供商（在这种情况下可能是您）与客户之间的合同。 它通常包括满足或缺少SLA的后果。 简而言之，服务水平协议是SLO +处罚。
为避免处罚，您应始终拥有比您在SLA中与客户达成协议的更严格的内部SLO。 例如，对于99.8％的SLO，您不应该签订99.9％的SLA合同。

#### CALCULATING YOUR ERROR BUDGET
Error Budget 告诉你指定时间内未达到 SLO 目标的频率，在支付罚款之前（由于违背了 SLA 协议），它还说明了在一段时间内对用户体验造成了多大影响。
例如，对于 99.9% 的正常运行时间目标，允许的一个月内故障的时间是 30 天 * 1440 分/天 * 0.1% = 43.2 分钟。

Error Budget 也应该在你的开发流程和业务决策中考虑在内。

#### The Cost of Nines
当可用性是主题时，您经常可以听到人们讨论9。 当服务提供三个九（99.9％）或五个九（99.999％）可用性时，这意味着什么？
你可能会认为三到九个服务之间的可用性差异可以忽略不计，但在实践中，它是巨大的！ 特别是如果我们谈论运行这些高弹性系统所需的工作量和成本。

#### A three nines service
A three nines service that has 99.9% uptime leads to 0.1% error budget. It means that before you violate your SLO, you can have:
- 43.2 minutes of downtime (in a 30 day period)
- 2.16 hours of downtime (in a 90 days period)
- 8.76 hours of downtime (in a 365 days period)

#### A five nines service
With a five nines service, (99.999% uptime and 0.001% error budget) the difference is shocking.Before you violate your SLO you can have:
- 25.9 seconds of downtime (in a 30 day period)
- 77.8 seconds of downtime (in a 90 days period)
- 5.26 minutes of downtime (in a 365 days period)
Only 25.9 seconds, let that sink in.
Can you or your monitoring product respond that fast to an ongoing incident?

#### Error Budget in practice
Error Budget 可以让你知道故障容忍的程度。但是如何使用它去做产品和商业决定？
例如，当你用完每月故障预算时，你应该考虑暂停发布，因为发布新的代码或配置始终会带来失败的风险。

Error Budget 还有助于优化资源使用情况。例如，当 uptime 运行时间远远好于预期时，你可以重新审视下冗余资源。Error Budget 还可以帮助利益相关者决定在产品中包含更多功能或者更多以可靠性为重点的开发。

总而言之，Error Budget 有助于回答以下问题：
- Should I release? Did I use all of my error budget for this month?
- Should I overprovision my resources or can I go hotter?
- Should I write a new feature or improve reliability?
- Should we do more testing?

使用 Error Budget 解决了开发与 SRE（Site Reliability Engineering）之间的根本冲突，因为目标不再是 "zero outages" 零中断。正如我们前面提到的，提高可靠性通常会逐渐增加成本。金额来自：
- 冗余资源的成本。
- 工程师构建故障转移解决方案而非面向用户的功能的机会成本。

#### Target level of availability
服务的可接受可用性水平在很大程度上取决于您的业务和客户需求，但您应该记住，智能手机，WiFi网络，电源和ISP等消费类设备的可用性低于99.999％。 这就是为什么用户通常无法区分99.99％和99.999％的可用性。 它们无法在WiFi连接和服务中断之间产生影响。

要计算目标可用性，请考虑一下标准：
- What level of service will the users expect?
- Does this service tie directly to revenue (either your income or your customers’ revenue)?
- Is this a paid service, or is it free?
- If there are competitors in the marketplace, what level of service do those competitors provide?
- Is this service targeted at consumers (B2C), or at enterprises (B2B)?
- How much would it cost to increase availability and how much would it increase the revenue?



## 2、监控基础
“To observe and check the progress or quality of (something) over a period of time; keep under systematic review.” 监控，牛津字典。
对于 "service monitoring"，我们指的是收集、处理、汇总、显示系统相关的实时数据。要分析数据，首先需要从系统中提取指标，例如特定应用程序实例的 CPU 使用情况，我们称之为 extraction instrumentation 提起仪器。
你可以手动检测系统，但大多数可用的监控解决方案都提供开箱即用的监控。在许多情况下，检测意味着添加带有性能开销的额外逻辑和代码片段。但是，通过监控和检测，目标是实现低开销，但并不一定意味着更大的性能影响对于更好的系统可见性是不合理的。
值得一提的是，日志记录与监控不相同。 日志是具有时间戳的不可变条目，您可以从中提取度量并随时监控其数量 - 例如给定时间段内的错误日志数。
监控基础架构的最常见用例是发现趋势，比较一段时间内的指标，构建实时洞察的仪表板，以及根据您的指标进行调试和警报。

### THE FOUR SIGNALS
每个服务都是不同的，您可以监控它们的许多方面。度量标准可以从CPU使用率等低级资源到注册次数等高级业务指标。我们建议您观看所有服务的这些信号：
- Error rate: Because errors are user facing and immediately affect your customers.
- Latency: Because the latency directly affects your customers.
- Throughput: The traffic helps to understand the context of increased error rates and the latency too.
- Saturation: It tells how “full” your service is. If the CPU usage is 90%, can your system handle more traffic?

### BLACK BOX AND WHITE BOX MONITORING
我们可以区分两种监控：黑盒和白盒监控。 它们都是彻底观察系统所必需的。

#### Black Box Monitoring ( probing )
黑盒监控（或探测）指的是在系统外部观察。由于监控者和被监控者彼此独立，因此共享故障的可能性较小。
正常运行时间检查是黑盒监控的一个很好的例子。在这种情况下，外部服务会在公共接口上探测您的系统。 检查可以是无状态的（例如，周期性地调用HTTP端点），但对于有状态服务，您可能需要具有实际数据写入和读取的有状态探测。
还可以从全球不同位置检查系统状态，并收集有关应用程序可用性和延迟的详细信息。

#### White Box Monitoring
当您通过仪器监控的系统提供指标时，我们使用术语白盒监控。 这种监视对于调试和收集有关内存使用情况，堆大小等数据至关重要。
由于白盒监控指标由运行系统本身提供，因此在事件发生期间可能会失去对数据的访问权限。


### SAMPLING RATE
由于您的大多数指标都是连续数据，因此您需要对其进行抽样。
通过抽样，您将失去准确性，这意味着数据可能会误导甚至阻止您发现问题的根本原因。 更频繁的采样提供了准确的结果，但开销可能很大，因此您必须选择合理的采样率。
更频繁的采样会给观察到的应用程序带来更大的性能开销，并且由于您必须存储和处理这些样本，因此成本也更高。
选择合适的采样率很困难，因为您还必须考虑受监控的资源和监控工具的功能。 我认为可以安全地说，为了监控应用程序中运行实例的数量，您可以选择较低的采样率（如一分钟），同时监控CPU使用率需要更高的频率。

### Metrics aggregations
另一个重要主题是选择聚合方法来对您的指标数据进行采样，分析和可视化。
您可以使用许多聚合方法来采样响应时间。 例如，您可以查找平均值，中位数，第95百分位数和第99百分位数（这意味着95％或99％的请求是在该时间之后提供的），但请记住，这些请求包含有关延迟的非常不同的信息。
例如，第95和第99响应时间是衡量一般用户体验的有用指标，但不提供有关最慢响应时间的任何信息。
处理和可视化已经采样的指标时要小心！ 某些聚合彼此不兼容，例如，如果使用较大的采样周期或采用第85个聚合的样本，则获取最后一小时的中位数将产生错误结果。


### ALERTING PHILOSOPHY
警报是一个将您设置的规则与指标相匹配的流程，并在您的条件匹配时通过一个或多个渠道发送通知。 警报对于运行可靠系统至关重要。 它们可以帮助您在系统中快速发现问题，并在灾难发生之前解决或最小化问题。
设置警报可能是一个复杂的过程，正确配置它们需要时间。
关于警报，您必须了解一件重要的事情：当一切正常时，您永远无法达到这一点，您永远不必再次触摸警报配置。 您的服务，业务和SLO会随着时间的推移而发生变化，因此您需要不时重新访问警报以匹配这些新情况和要求。
警报是一个不断发展的过程：如果您的监控系统显示的问题是真实的，您需要处理它。 如果是误报，您需要改进警报配置，甚至删除不足的警报。

#### Alerting Symptoms
您的系统停止执行有用工作的事实是一种症状。 最有效的警报通常基于症状。 这些警报取决于服务的外部可见行为。 了解对您的用户重要的内容，并使用此知识来设置您的标准。 不要在无意义的警报上浪费你的时间。 对于基于症状的警报，我们通常建议使用寻呼系统 - 因为这些问题可能会对您的业务造成最大的伤害。
Examples or symptoms:
- APIs: Downtime of your system, slow response times for minutes
- Jobs: Slow batch time completion time
- Pipelines: Big Buffer sizes, old data waiting for too long

#### JUAN criterion
 Judgement, Urgent, Actionable, Necessary
如果您从未听说过JUAN标准系统，请不要感到惊讶。 这是Google最新的Cloud Next大会上推出的新概念。 这里重要的是JUAN标准提供了验证警报的帮助，因此您可以停止在无意义的警报上浪费您的时间。

#### Alerting risks
设置警报会带来一些风险：如果目标太高（SLO），您最终会收到太多警报，目标太低，甚至可能没有注意到糟糕的服务质量。 更糟糕的情况是，当大部分警报不必要时，这会导致人们完全忽略它们（包括重要的警报）。
当团队无法跟上警报时，快速增长的用户群可以接受相对较低的SLO。 但是，如果这是一个有意识的商业决策，那么它是唯一可以接受的，并且您确信您的用户不会有糟糕的体验，从而使他们远离您的服务。 您的警报必须受到控制并与您的业务保持一致。

#### Cliff alerts
SLO警报仅在服务中断时触发，但悬崖警报甚至可以预测SLO违规，例如当磁盘使用率已达到90％且不断增长时。 您可以为许多类型的指标设置悬崖警报。 例如，如果您知道您的计算机只能处理具有特定内存设置的有限数量的Docker容器，那么当您应该向群集添加更多VM以防止停机时，您希望收到警报。

#### Notification delivery







## 构建可靠系统的技术



## node.js 使用 PROMETHEUS 监控


