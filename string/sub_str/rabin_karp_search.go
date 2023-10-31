package sub_str

// Rabin-Karp fingerprint search
// 指纹搜索基本思路：取模算哈希 modular hashing
// 乐观方式Monte Carlo: 如果hash匹配则return，有概率是错的。
// 悲观方式Las Vegas: 如果hash匹配，再做一次校验，匹配则return.

// 取模算哈希

//Modular hash function. Using the notation t_i for txt.charAt(i),
//x_i =t_i * R^{M-1} + t_{i+1} * R^{M-2} +...+ t_{i+M-1} * R^0 (modQ)

//Intuition. M-digit, base-R integer, modulo Q.
//Horner's method. Linear-time method to evaluate degree-M polynomial.
// 计算可以利用通项公式，提前计算好 R^{M-1}，无需for循环，就能由xi 计算出 xi+1项：
// x_{i+1}=( x_i – t_i * R^{M–1} ) R + t_{i+M}
// 所以对于M次方，整体只需要线性时间。

// 理论上来说，如果素数Q足够大，哈希冲突的概率为 1/N.
// 实际使用当中，选择一个非常大的素数作为Q（N<<Q, 但是Q不至于导致溢出）, 则哈希冲突概率为 1/Q.

// 优势：
// 支持扩展到2d的模式
// 支持一次搜索多个模式

// 劣势:
// cpu的数值计算比字符对比慢
// Las Vegas版本需要回退指针
// 最差情况下性能不好

//Theory.
//If Q is a sufficiently large random prime (about M*N^2),
//then the probability of a false collision is about 1 / N.

//Practice.
//Choose Q to be a large prime (but not so large to cause overflow).
//Under reasonable assumptions, probability of a collision is about 1 / Q.

//Advantages.
//・Extends to 2d patterns.
//・Extends to finding multiple patterns.

//Disadvantages.
//・Arithmetic ops slower than char compares.
//・Las Vegas version requires backup.
//・Poor worst-case guarantee.

func RabinKarpSearch() int {
	//todo
	return -1
}
