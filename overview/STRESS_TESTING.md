LACNet ran different tests to understand the performance of the network with the GAS distribution mechanism when the network is being heavily utilized. The objective of the stress tests was to identify a utilization threshold where network performance begins to deteriorate. For the stress tests, LACNet set up smart contracts of various complexities and executed the stress test scenarios that were monitored and analyzed by considering the impact on network performance.

Network deterioration is quantified by examining block interval time. The genesis file for the network specifies the block interval time as 2 seconds. Blocks that take longer than 2 seconds to be produced constitute a situation in which network performance has deteriorated and has an impact on the transaction finality and throughput of the network. Thresholds for network deterioration are defined as situations where the likelihood of experiencing a block greater than 2 seconds is higher than normal. Deterioration is broken down into two types: 1) Moderate: block interval time between 2 and 4 seconds, and 2) Severe: block interval time greater than 4 seconds.

The hardware of the nodes is critical for the performance of the network. Also, all boots and validators must have similar hardware characteristics. Otherwise, the lower performing get delayed in transaction processing leading to delays in transaction propagation and block generation. The tests were run with the minimum hardware requirements indicated for LACNet Mainnet, which are:

* CPU: 4 virtual CPUs
* RAM Memory: 16 GB
* Hard Disk: 100 GB SSD (70,000 IOPS READ, 50,000 IOPS WRITE)
* Operating System: Ubuntu 16.04, Ubuntu 18.04, Ubuntu 20.04, Centos7, always 64 bits

The results are presented in Table 1:

[TABLE 1]

The results from the tests allow for several conclusions that need to be understood as conditioned to the hardware characteristics utilized:

* The network was able to process 170 tx/s simple transactions of 123,339 GAS/tx without degradation.
* The network was able to process 105M gas per block with a moderate degradation that delayed the generation of the block to 3-4 seconds instead of the expected 2 seconds.
* The two main parameters that produce block degradation are number of transactions and GAS used. The increase of these two parameters lead independently to block degradation.

