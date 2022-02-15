LACNet ran different tests to understand the performance of the network with the GAS distribution mechanism when the network is being heavily utilized. The objective of the stress tests was to identify a utilization threshold where network performance begins to deteriorate. For the stress tests, LACNet set up smart contracts of various complexities and executed the stress test scenarios that were monitored and analyzed by considering the impact on network performance.

Network deterioration is quantified by examining block interval time. The genesis file for the network specifies the block interval time as 2 seconds. Blocks that take longer than 2 seconds to be produced constitute a situation in which network performance has deteriorated and has an impact on the transaction finality and throughput of the network. Thresholds for network deterioration are defined as situations where the likelihood of experiencing a block greater than 2 seconds is higher than normal. Deterioration is broken down into two types: 1) Moderate: block interval time between 2 and 4 seconds, and 2) Severe: block interval time greater than 4 seconds.

The hardware of the nodes is critical for the performance of the network. Also, all boots and validators must have similar hardware characteristics. Otherwise, the lower performing get delayed in transaction processing leading to delays in transaction propagation and block generation. The tests were run with the minimum hardware requirements indicated for LACNet Mainnet, which are:

* CPU: 4 virtual CPUs
* RAM Memory: 16 GB
* Hard Disk: 100 GB SSD (70,000 IOPS READ, 50,000 IOPS WRITE)
* Operating System: Ubuntu 16.04, Ubuntu 18.04, Ubuntu 20.04, Centos7, always 64 bits

The results are presented in Table 1:

|     Contract    | Tx/s Sent | Max_Tx/Block | Gas_Used/block | Gas/Transaction | Degradation_block_time | Finality |
|:---------------:|:---------:|:------------:|:--------------:|:---------------:|:--------:|:----------------------:|
| Simple Contract |    200    |    270 tx    |   33,301,601   |     123,339     |   ----   |        170 Tx/s        |
|      ERC20      |     70    |    134 tx    |   18,822,763   |     140,468     |   ----   |         67 Tx/s        |
|      ERC20      |    100    |    206 tx    |   28,930,167   |     140,468     |   ----   |         90 Tx/s        |
|      ERC20      |    150    |    364 tx    |   51,110,429   |     140,468     |  3-4 seg |        100 Tx/s        |
|     Identity    |     80    |    345 tx    |   93,139,220   |     269,968     |  3-4 seg |         77 Tx/s        |
|  Register Covid |     50    |    104 tx    |   35,862,281   |     344,830     |   ----   |         50 Tx/s        |
|  Register Covid |     80    |    306 tx    |   105,495,141  |     344,830     |  3-4 seg |         60 Tx/s        |
|   ARM Aduanas   |     7     |     21 tx    |   64,562,539   |    3'074,149    |  2-3 seg |         7 Tx/s         |
|   ARM Aduanas   |     10    |     41 tx    |   126,040,127  |    3'074,149    |  4-6 seg |         7 Tx/s         |
|   DIDRegistry   |     50    |    122 tx    |   12,046,599   |      98,669     |    ---   |         50 Tx/s        |
|   DIDRegistry   |    100    |    241 tx    |   23,826,973   |      98,669     |  3-4 seg |         70 Tx/s        |
|     ERC-721     |     50    |    102 Tx    |   26,491,918   |     310,468     |   ----   |         47 Tx/s        |
|     ERC-721     |    100    |    236 Tx    |   54,322,443   |     310,468     |   ----   |         57 Tx/s        |
|     ERC-721     |    150    |    280 Tx    |   61,007,722   |     310,468     |   3-seg  |         80 Tx/s        |

The results from the tests allow for several conclusions that need to be understood as conditioned to the hardware characteristics utilized:

* The network was able to process 170 tx/s simple transactions of 123,339 GAS/tx without degradation.
* The network was able to process 105M gas per block with a moderate degradation that delayed the generation of the block to 3-4 seconds instead of the expected 2 seconds.
* The two main parameters that produce block degradation are number of transactions and GAS used. The increase of these two parameters lead independently to block degradation.

