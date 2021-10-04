# 算法训练营第二周
[toc]

## 本周总结 - 题解思路

### 跟两个端点、子段、窗口有关

#### 一端不变，并且满足区间减法性质，考虑使用前缀和/差分
`区间减法性质：[l, r]的信息可以由[1, r] 和[1, l - 1]推导出，[l, r] = s[r] - s[l - 1]`
`差分是前缀和的逆运算`
##### 注意点
+ 前缀中的min，max问题只是维护前缀信息，而不是前缀和（不满足区间减法性质）
##### 相关例题传送门
+ [统计「优美子数组」](https://leetcode-cn.com/problems/count-number-of-nice-subarrays/)
+ [二维区域和检索 - 矩阵不可变](https://leetcode-cn.com/problems/range-sum-query-2d-immutable/)
+ [航班预订统计](https://leetcode-cn.com/problems/corporate-flight-bookings/)
+ [最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)

#### 双端变化，维护双端信息，与双端之间的位置信息无关，考虑使用双指针扫描、滑动创港口
`双指针扫描方向：双指针一起移动/两端向中间移动`
##### 注意点
+ 避免重复值（相同的数不进行二次枚举）
##### 相关例题传送门
+ [两数之和](https://leetcode-cn.com/problems/two-sum/)
+ [两数之和 II - 输入有序数组](https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/)
+ [三数之和](https://leetcode-cn.com/problems/3sum/)
+ [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)

#### 一端固定，另一端需要对候选集合/一系列点/范围维护max，min等信息，考虑使用单调栈、单调队列(堆，有序集合，平衡树)
##### 注意点
+ 单调栈的单调性（递增 or 递减）
+ 固定端加入集合的时间点（构建单调栈前加入 or 构建单调栈后加入）
##### 相关例题传送门
+ [柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)
+ [滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)
+ [接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)

## 本周作业

+ [子域名访问计数](https://github.com/finger2011/algorithm/blob/master/week02/subdomain_visit_count.go)
+ [数组的度](https://github.com/finger2011/algorithm/blob/master/week02/degree_of_an_array.go)
+ [元素和为目标值的子矩阵数量](https://github.com/finger2011/algorithm/blob/master/week02/number_of_submatrices_that_sum_to_target.go)
+ [和为 K 的子数组](https://github.com/finger2011/algorithm/blob/master/week02/subarray_sum_equals_k.go)

## 实战例题

### 第三课

#### 无序集合、映射

+ [两数之和](https://github.com/finger2011/algorithm/blob/master/week02/onclass/two_sum.go)
+ [模拟行走机器人](https://github.com/finger2011/algorithm/blob/master/week02/onclass/walking_robot_simulation.go)
+ [字母异位词分组](https://github.com/finger2011/algorithm/blob/master/week02/onclass/group_anagrams.go)
+ [串联所有单词的子串](https://github.com/finger2011/algorithm/blob/master/week02/onclass/substring_with_concatenation_of_all_words.go)

#### LRU
+ [LRU 缓存机制](https://github.com/finger2011/algorithm/blob/master/week02/onclass/lru_cache.go)

### 第四课

#### 前缀和、差分

+ [统计「优美子数组」](https://github.com/finger2011/algorithm/blob/master/week02/onclass/count_number_of_nice_subarrays.go)
+ [二维区域和检索 - 矩阵不可变](https://github.com/finger2011/algorithm/blob/master/week02/onclass/range_sum_query_2d_immutable.go)
+ [航班预订统计](https://github.com/finger2011/algorithm/blob/master/week02/onclass/corporate_flight_bookings.go)
+ [最大子序和](https://github.com/finger2011/algorithm/blob/master/week02/onclass/maximum_subarray.go)

#### 双指针扫描、滑动窗口
+ [两数之和](https://github.com/finger2011/algorithm/blob/master/week02/two_sum.go)
+ [两数之和 II - 输入有序数组](https://github.com/finger2011/algorithm/blob/master/week02/onclass/two_sum_input_array_is_sorted.go)
+ [三数之和](https://github.com/finger2011/algorithm/blob/master/week02/onclass/3sum.go)
+ [盛最多水的容器](https://github.com/finger2011/algorithm/blob/master/week02/onclass/container_with_most_water.go)



