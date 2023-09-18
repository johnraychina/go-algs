package basic

// Minimum Spanning Tree 最小生成树，应用场景包括：
//・Dithering.
//・Cluster analysis. 聚簇分析
//・Max bottleneck paths. 最大瓶颈路径
//・Real-time face verification. 实时人脸验证
//・LDPC codes for error correction. LDPC 纠错码, see: https://zhuanlan.zhihu.com/p/514670102
//・Image registration with Renyi entropy. 基于Renyi熵的图像配准（registration）
//・Find road networks in satellite and aerial imagery.寻找卫星航空图像中的道路
//・Reducing data storage in sequencing amino acids in a protein.减少蛋白质序列的数据存储
//・Model locality of particle interactions in turbulent fluid flows.模拟湍流中粒子相互作用的局部性
//・Autoconfig protocol for Ethernet bridging to avoid cycles in a network. 自动配置以太网桥避免环
//・Approximation algorithms for NP-hard problems (e.g., TSP, Steiner tree). NP问题估计(Traveling Salesman Problem旅行商问题, Steiner树-组合优化）
//・Network design (communication, electrical, hydraulic, computer, road).网络设计（通信、电力、液压系统、计算机网络、道路）

// 定义“切分”: 一个图的切分指将定点分割成两个非空的集合
// 定义“跨连接边”:一个“跨连接边”连接两个集合中的两点
// 切分的性质：给定一个切分，跨连接边中的最小边属于MST（最小生成树）
//Def. A cut in a graph is a partition of its vertices into two (nonempty) sets.
//Def. A crossing edge connects a vertex in one set with a vertex in the other.
//Cut property. Given any cut, the crossing edge of min weight is in the MST.
