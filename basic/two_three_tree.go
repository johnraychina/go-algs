package basic

// 2-3树重要特性：
// 1. 通过分裂、合并，可以使树完全平衡，可以解决BST退化为链表的问题。
// 2. 节点分裂时，子树之间仍然是相对有序的，只需要局部做转换，开销非常小。
// Splitting a 4-node is a local transformation: constant number of operations.
// Invariants. Maintains symmetric order and perfect balance.
// Guaranteed logarithmic performance for search and insert.

// 实现复杂 ---> 红黑树是基于它简化了
//・Maintaining multiple node types is cumbersome.
//・Need multiple compares to move down tree.
//・Need to move back up the tree to split 4-nodes.
//・Large number of cases for splitting.
