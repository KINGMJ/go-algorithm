# go-algorithm

golang algorithm practice

## 测试生成

```
cd xxx
# 初始化
ginkgo bootstrap
# 添加单元测试
ginkgo generate xxx
# 运行测试
ginkgo
```

## 50_algorithm（50 个必会的数据结构与算法）

- 01_dynamic_array：实现一个支持动态扩容的数组
- 02_ordered_array：实现一个大小固定的有序数组，支持动态增删改操作
- 03_merge_sorted_array：合并两个有序数组
- 04_circular_linked_list：循环链表实现
- 04_doubly_linked_list：双向链表实现
- 04_singly_linked_list：单链表实现
- 05_reverse_list：反转链表
- 06_merge_sorted_linked_list：合并两个有序链表
- 07_middle_node：链表的中间结点
- 08_stack_on_array：数组实现一个顺序栈
- 09_stack_on_list：链表实现一个顺序栈
- 10_simply_browser：编程模拟实现一个浏览器的前进、后退功能
- 11_queue_on_array：数组实现一个顺序队列
- 12_queue_on_list：链表实现一个链式队列
- 13_circular_queue：实现一个循环队列
- 14_permutations：全排列
- 15_sort：排序算法
- 22_binary_tree：二叉树的前、中、后、层序遍历

## structure （实现的各种数据结构）

- binary_tree：二叉树
- binary_search_tree：二叉查找树
- linked_list：链表
- sq_queue：顺序队列
- sq_stack：顺序栈

## leetcode

- 01_tow_sum: 1. 两数之和
- 02_contains_duplicate：217. 存在重复元素
- 03_sell_stock：121. 买卖股票的最佳时机
- 04_n_queens：51. N 皇后
- 05_permutations：46. 全排列
- 06_assign_cookies：455. 分发饼干
- 07_product_of_array_except_self：238. 除自身以外数组的乘积
- 08_valid_parentheses：20. 有效的括号
- 09_valid_palindrome：125. 验证回文串
- 10_invert_binary_tree：226. 翻转二叉树
- 11_valid_anagram：242. 有效的字母异位词
- 12_binary_search：704. 二分查找
- 13_flood_fill：733. 图像渲染
- 14_lowest_common_ancestor_of_bst：235. 二叉搜索树的最近公共祖先
- 15_balanced_binary_tree：110. 平衡二叉树
- 16_linked_list_cycle：141. 环形链表
- 17_implement_queue_using_stacks：232. 用栈实现队列
- 18_first_bad_version：278. 第一个错误的版本
- 19_ransom_note：383. 赎金信
- 20_climbing_stairs：70. 爬楼梯
- 21_longest_palindrome：409. 最长回文串
- 22_majority_element：169. 多数元素
- 23_remove_duplicates：26. 删除有序数组中的重复项
- 24_move_zeroes：283. 移动零
- 25_remove_element：27. 移除元素
- 26_remove_duplicates_ii：80. 删除有序数组中的重复项 II
- 27_two_sum_ii：167. 两数之和 II - 输入有序数组
- 28_reverse_string：344. 反转字符串
- 29_squares_of_sorted_array：977. 有序数组的平方
- 30_container_with_most_water：11. 盛最多水的容器
- 31_jump_game：55. 跳跃游戏
- 32_jump_game_ii：45. 跳跃游戏 II
- 33_candy：135. 分发糖果
- 34_reverse_linked_list：206. 反转链表
- 35_reverse_linked_list_ii：92. 反转链表 II



## prompt

@Web @https://leetcode.cn/problems/move-zeroes/description/
请在 leetcode 目录生成 10_xxx 目录，里面包含两个文件：xxx.go 和 xxx_test.go。分别是算法实现和对应的测试用例。
算法实现由我来实现，你只需要将函数定义出来就可以了；测试用例你需要提供尽可能完整的。请直接生成可以Apply的文件。













中等题：
142 环形链表 II（找到环的入口）
19 删除链表的倒数第 N 个结点
2 两数相加
86 分隔链表
143 重排链表
148 排序链表
82 删除排序链表中的重复元素 II
24 两两交换链表中的节点
61 旋转链表
147 对链表进行插入排序
138 复制带随机指针的链表
困难题：
23 合并K个升序链表
2. 25 K 个一组翻转链表

建议刷题顺序：
接着是 19（快慢指针的应用）
后面按照个人兴趣选择

每类题目的关键点：
反转类：206, 92, 25
快慢指针：141, 142, 19
合并类：21, 23
排序类：148, 147
特殊指针：138